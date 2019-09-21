package interpreters

import (
	"encoding/json"
	"github.com/Aiyane/golinq/functions"
	"github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
	"unsafe"
)

func groupVisit(sqlExpr interface{}, groupDataSources types.GroupDataSources, index int) interface{} {
	switch subExpr := sqlExpr.(type) {
	case *sql_parser.Var:
		return groupVisitVar(subExpr, groupDataSources, index)
	case *sql_parser.Const:
		return groupVisitConst(subExpr)
	case *sql_parser.As:
		return groupVisit(subExpr.Value, groupDataSources, index)
	case *sql_parser.Func:
		return groupVisitFunc(subExpr, groupDataSources, index)
	case *sql_parser.Link:
		return groupVisitLink(subExpr)
	case *sql_parser.Case:
		return groupVisitCase(subExpr, groupDataSources, index)
	case []sql_parser.SqlToken:
		var ret []interface{}
		for _, subExpr := range subExpr {
			ret = append(ret, groupVisit(subExpr, groupDataSources, index))
		}
		return ret
	default:
		return nil
	}
}

func groupVisitCase(sqlExpr *sql_parser.Case, groupDataSources types.GroupDataSources, index int) interface{} {
	for _, item := range sqlExpr.Values {
		condition := item.Condition
		if condition == nil || groupVisit(condition, groupDataSources, index).(bool) {
			return groupVisit(item.Value, groupDataSources, index)
		}
	}
	return nil
}

func groupVisitLink(sqlExpr *sql_parser.Link) interface{} {
	if ret, exist := env.LinkCache[sqlExpr.Id]; exist {
		return ret
	}
	var ret []interface{}
	for _, resRecord := range env.Link[sqlExpr.Id] {
		var line []interface{}
		for key := range preFieldName {
			v, _ := recordValue(resRecord, key)
			line = append(line, v)
		}
		if len(line) == 1 {
			ret = append(ret, line[0])
		} else {
			ret = append(ret, line)
		}
	}
	env.LinkCache[sqlExpr.Id] = ret
	return ret
}

func groupVisitFunc(sqlExpr *sql_parser.Func, groupDataSources types.GroupDataSources, index int) interface{} {
	function := getFunc(index, sqlExpr.Name)

	var funcArgs []interface{}
	for _, _funcArg := range sqlExpr.Args {
		funcArg := groupVisit(_funcArg, groupDataSources, index)
		if funcArg == "*" {
			funcArgs = append(funcArgs, groupFuncStar(groupDataSources)...)
		} else {
			funcArgs = append(funcArgs, funcArg)
		}
	}
	return function(funcArgs...)
}

func groupVisitConst(sqlExpr *sql_parser.Const) interface{} {
	return sqlExpr.Value
}

// 支持 A.b.c.d 这类操作
func groupVisitVar(sqlExpr *sql_parser.Var, groupDataSources types.GroupDataSources, index int) interface{} {
	name := groupVisit(sqlExpr.Value, groupDataSources, index).(string)
	if rename, exist := env.ReTableName[name]; exist {
		name = rename
	}
	next := sqlExpr.Next
	// 单表省略表名
	if next == nil {
		next = sqlExpr
		name = singleTableName
		sqlExpr = &sql_parser.Var{
			Value: &sql_parser.Const{Value: singleTableName},
			Next:  next,
		}
	}
	_subName := groupVisit(next.Value, groupDataSources, index).(string)
	_subNext := next.Next
	var value functions.GroupRecord
	for _, instanceDicts := range groupDataSources {
		for _, instanceDict := range instanceDicts {
			subName, subNext := _subName, _subNext
			subValue, _ := recordValue(instanceDict[name], subName)
			for subNext != nil {
				subName = groupVisit(subNext.Value, groupDataSources, index).(string)
				subValue, _ = recordValue(subValue, subName)
				subNext = subNext.Next
			}
			// 如果是group by的字段查询,查出第一个就可以了,因为group后的值是相等的
			// 如果是函数之类的, 很可能计算数量或者其他, 那么需要取出全部
			if varInGroupMap(sqlExpr) {
				return subValue
			}
			value = append(value, subValue)
		}
		return value
	}
	return nil
}

func groupVisitSelect(sqlExpr []sql_parser.SqlToken, groupDataSources types.GroupDataSources) interface{} {
	var resRecord interface{}
	if useStruct() {
		resRecord = newStruct(types.RecordStructName)
		hadSetStructFieldValue = make(types.UsedKey, 10)
		tmpFieldValue = make(map[string]interface{}, 10)
	} else {
		resRecord = make(map[string]interface{})
	}

	if visit(sqlExpr[1], nil, -1) == true {
		return groupSelectStar(resRecord, groupDataSources)
	}

	for i, subExpr := range sqlExpr[2:] {
		name := env.FieldName[i]
		val := groupVisit(subExpr, groupDataSources, i)
		switch value := val.(type) {
		case functions.GroupRecord:
			setValue(resRecord, name, value[len(value)-1], false)
		default:
			setValue(resRecord, name, value, false)
		}
	}
	return resRecord
}

// group 语句顶层接口, 做 group 聚合与 where 筛选, 给后续的 having 筛选创造新的数据源
func getGroupDataSources(dataSources types.DataSources) types.GroupDataSources {
	getGroupsList(expr.GroupExpr)
	groupDataSources := make(types.GroupDataSources, 10)
	if env.HasIndex {
		groupIndexLoop(dataSources, groupDataSources)
	} else {
		groupMainLoop(dataSources, make(types.InstanceDict, 10), groupDataSources)
	}
	return groupDataSources
}

// 主循环
func groupMainLoop(dataSources types.DataSources, instanceDict types.InstanceDict,
	groupDataSources types.GroupDataSources) {
	groupMainIterator(groupDataSources, make(types.UsedKey, 10), 0, len(dataSources), instanceDict, dataSources)
}

// 主迭代器
func groupMainIterator(groupDataSources types.GroupDataSources, usedKey types.UsedKey, num int, dataLen int,
	instanceDict types.InstanceDict, dataSources types.DataSources) {
	if num == dataLen {
		copyInstanceDict := deepCopyInstanceDict(instanceDict, "", nil)
		groupWhereFilter(copyInstanceDict, groupDataSources)
	} else {
		// 取一张表
		for name, records := range dataSources {
			if _, ok := usedKey[name]; !ok {
				usedKey[name] = struct{}{}
				for _, record := range records {
					instanceDict[name] = record
					groupMainIterator(groupDataSources, usedKey, num+1, dataLen, instanceDict, dataSources)
					delete(instanceDict, name)
				}
				delete(usedKey, name)
				return
			}
		}
	}
}

// 索引循环
func groupIndexLoop(dataSources types.DataSources, groupDataSources types.GroupDataSources) {
	var indexDict = make(types.IndexDict, 10)
	makeIndexDict(env.Index, dataSources, indexDict)
	name, indexRecords := popItem(indexDict)
	indexLen := len(indexDict)
	hasDataSources := !isEmptyMap(dataSources)
	for indexKey, records := range indexRecords {
		groupIndexIterator(hasDataSources, groupDataSources, indexKey, make(types.UsedKey, 10), indexDict,
			dataSources, make(types.InstanceDict, 10), 0, indexLen, name, records)
	}
}

// 索引迭代器
func groupIndexIterator(hasDataSources bool, groupDataSources types.GroupDataSources, indexKey types.IndexKey,
	usedKey types.UsedKey, indexDict types.IndexDict, dataSources types.DataSources, instanceDict types.InstanceDict,
	num int, indexNum int, tableName string, records types.Records) {
	if num == indexNum {
		for _, record := range records {
			copyInstanceDict := deepCopyInstanceDict(instanceDict, tableName, record)
			if hasDataSources {
				groupMainLoop(dataSources, instanceDict, groupDataSources)
			} else {
				groupWhereFilter(copyInstanceDict, groupDataSources)
			}
		}
	} else {
		for name, indexRecords := range indexDict {
			if _, ok := usedKey[name]; !ok {
				usedKey[name] = struct{}{}
				if records2, exist2 := indexRecords[indexKey]; exist2 {
					for _, record := range records2 {
						instanceDict[name] = record
						groupIndexIterator(hasDataSources, groupDataSources, indexKey, usedKey, indexDict,
							dataSources, instanceDict, num+1, indexNum, tableName, records)
						delete(instanceDict, name)
					}
				}
				delete(usedKey, name)
				return
			}
		}
	}
}

// 筛选器
func groupWhereFilter(instanceDict types.InstanceDict, groupDataSources types.GroupDataSources) {
	if expr.WhereExpr == nil || visitFunc(expr.WhereExpr, instanceDict, -1) == true {
		pk := getRecordPk(instanceDict)
		if instanceDicts, exist := groupDataSources[pk]; exist {
			groupDataSources[pk] = append(instanceDicts, instanceDict)
		} else {
			groupDataSources[pk] = []types.InstanceDict{instanceDict}
		}
	}
}

// 获取记录主键值
// 先序列化, 后转为字符串
func getRecordPk(instanceDict types.InstanceDict) string {
	pkSlice := make([]interface{}, 0, 10)
	for _, _groupKey := range env.GroupList {
		groupKey := _groupKey.([]string)
		name := groupKey[0]
		if rename, exist := env.ReTableName[name]; exist {
			name = rename
		}
		record, _ := instanceDict[name]
		for _, key := range groupKey[1:] {
			record, _ = recordValue(record, key)
		}
		pkSlice = append(pkSlice, record)
	}
	pk, _ := json.Marshal(pkSlice)
	// 据说这样转字符串更快??
	return *(*string)(unsafe.Pointer(&pk))
}

func getGroupsList(varExpr []*sql_parser.Var) {
	groupList := make([]interface{}, 0, 10)
	for _, subExpr := range varExpr {
		groupKey := getVar(subExpr)
		// 单表可以省略表名
		if singleTableName != "" && singleTableName != groupKey[0] {
			_groupKey := make([]string, 0, len(groupKey)+1)
			_groupKey = append(_groupKey, singleTableName)
			groupKey = append(_groupKey, groupKey...)
		}
		groupList = append(groupList, groupKey)
		setGroupMap(groupKey)
	}
	env.GroupList = groupList
}

func groupFuncStar(groupDataSources types.GroupDataSources) []interface{} {
	var argRes []interface{}
	for _, instanceDicts := range groupDataSources {
		length := len(instanceDicts)
		instanceDict := instanceDicts[0]
		for tableName, record := range instanceDict {
			recordKeyValue(record, func(columnName string, v interface{}) {
				row := []interface{}{v}
				i := 0
				for ; i+1 < length; i++ {
					instanceDict2 := instanceDicts[i+1]
					v, _ := recordValue(instanceDict2[tableName], columnName)
					row = append(row, v)
				}
				argRes = append(argRes, row)
			})
		}
		return argRes
	}
	return nil
}

// 调整重名
func groupAdjustName(newName string, resRecord interface{}, hadAdjust types.UsedKey,
	keyName map[string]string, value interface{}, tableName string, oldName string) {
	if value2, ok := recordValue(resRecord, newName); ok {
		newName2 := keyName[newName] + "_" + newName
		setValue(resRecord, newName, nil, true)
		groupAdjustName(newName2, resRecord, hadAdjust, keyName, value2, keyName[newName], oldName)
	}
	if lstInGroupMap([]string{tableName, oldName}) {
		setValue(resRecord, newName, value, false)
		hadAdjust[newName] = struct{}{}
	} else {
		setValue(resRecord, newName, nil, false)
		hadAdjust[newName] = struct{}{}
	}
}

func groupSelectStar(resRecord interface{}, groupDataSources types.GroupDataSources) interface{} {
	hadAdjust := make(types.UsedKey, 10)
	keyName := make(map[string]string)
	hadName := make(types.UsedKey, 10)
	for _, instanceDicts := range groupDataSources {
		for _, instanceDict := range instanceDicts {
			for tableName, record := range instanceDict {
				recordKeyValue(record, func(key string, value interface{}) {
					if value2, ok := recordValue(resRecord, key); ok {
						newName := tableName + "_" + key
						groupAdjustName(newName, resRecord, hadAdjust, keyName, value, tableName, key)
						// 存在的key未被调整过的状态
						if _, ok := hadAdjust[key]; !ok {
							newName2 := keyName[key] + "_" + key
							// 先将存在的删除
							setValue(resRecord, key, nil, true)
							groupAdjustName(newName2, resRecord, hadAdjust, keyName, value2, keyName[key], key)
						}
					} else {
						// 存在key已经是已有字段, 需要调整
						if _, ok := hadName[key]; ok {
							newName := tableName + "_" + key
							groupAdjustName(newName, resRecord, hadAdjust, keyName, value, tableName, key)
							// key不是已有字段, 并且key也没有存在的
						} else {
							if lstInGroupMap([]string{tableName, key}) {
								setValue(resRecord, key, value, false)
								hadName[key] = struct{}{}
								keyName[key] = tableName
							} else {
								setValue(resRecord, key, nil, false)
								hadName[key] = struct{}{}
								keyName[key] = tableName
							}
						}
					}
				})
			}
			return resRecord
		}
	}
	return resRecord
}

func setGroupMap(groupKey []string) {
	key, _ := json.Marshal(groupKey)
	env.GroupMap[*(*string)(unsafe.Pointer(&key))] = true
}

func varInGroupMap(expr *sql_parser.Var) bool {
	key := make([]string, 0, 10)
	for expr != nil {
		key = append(key, visitConst(expr.Value).(string))
		expr = expr.Next
	}
	pk, _ := json.Marshal(key)
	if env.GroupMap[*(*string)(unsafe.Pointer(&pk))] {
		return true
	}
	return false
}

func lstInGroupMap(key []string) bool {
	pk, _ := json.Marshal(key)
	if env.GroupMap[*(*string)(unsafe.Pointer(&pk))] {
		return true
	}
	return false
}
