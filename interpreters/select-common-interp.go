package interpreters

import (
	"github.com/Aiyane/golinq/functions"
	"github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
	"strconv"
)

// 该死的 Golang 没有泛型！！！
// 给不同 tag 的 SqlToken 对象分配不同的解释函数, 返回执行结果
// dataSources: 数据源
// env: 环境
// index: select 语句中的列号
// 例如:
// "SELECT COUNT(a.name), COUNT(a.age) FROM a" 中的
// 	 子语法树 SqlToken{COUNT(a.name)}
// 	 其 tag = FUNC, 其 index = 0 (备注说明: "COUNT(a.age)" 语句的 index = 1, 其他语句的 index = -1)
func visit(sqlExpr interface{}, dataSources interface{}, index int) interface{} {
	/*
		index: -1 时没有意义
	*/
	switch dataSourcesOrInstanceDict := dataSources.(type) {
	case types.DataSources:
		return visitDataSources(sqlExpr, dataSourcesOrInstanceDict)
	case types.InstanceDict:
		return visitInstanceDict(sqlExpr, dataSourcesOrInstanceDict, index)
	}
	switch subExpr := sqlExpr.(type) {
	case *sql_parser.Const:
		return visitConst(subExpr)
	case *sql_parser.Var:
		return subExpr.Value.Value
	}
	return nil
}

func visitDataSources(sqlExpr interface{}, dataSources types.DataSources) interface{} {
	switch subExpr := sqlExpr.(type) {
	case *sql_parser.Var:
		return visitFromVar(subExpr, dataSources)
	case *sql_parser.Inner:
		return visitInner(subExpr, dataSources)
	}
	return nil
}

func visitInstanceDict(sqlExpr interface{}, instanceDict types.InstanceDict, index int) interface{} {
	switch subExpr := sqlExpr.(type) {
	case *sql_parser.As:
		return visit(subExpr.Value, instanceDict, index)
	case *sql_parser.Link:
		return visitLink(subExpr)
	case *sql_parser.Case:
		return visitCase(subExpr, instanceDict, index)
	case *sql_parser.Func:
		return visitFunc(subExpr, instanceDict, index)
	case *sql_parser.Var:
		return visitVar(subExpr, instanceDict)
	case *sql_parser.Const:
		return visitConst(subExpr)
	case []sql_parser.SqlToken:
		var ret []interface{}
		for _, subExpr := range subExpr {
			ret = append(ret, visitInstanceDict(subExpr, instanceDict, index))
		}
		return ret
	}
	return nil
}

// 解释 LIMIT 语句
// 例如: res = [
// 	{"name": "li"},
//  {"age": 18},
// ]
// sqlExpr = SqlToken{LIMIT 1, 1}
// 返回: [
// 	{"age": 18},
// ]
func visitLimit(sqlExpr *sql_parser.Limit) {
	offset := visitConst(sqlExpr.Offset).(int)
	limit := visitConst(sqlExpr.Limit).(int)
	length := len(res)
	if offset >= length {
		res = []interface{}{}
	}
	end := offset + limit
	if end > length {
		end = length
	}
	res = res[offset:end]
}

// 链接子 SELECT 语句结果
// 例如: sqlExpr = SqlToken{Tag=LINK, Args=1}
// 表示链接 id=1 的 SELECT 语句结果, 结果保存在 env.Link 中
// env.Link[1] = [
// 	{"name": "li", "time": 20},
//  {"name": "jay", "time": 30},
// ]
// 返回:
// [
// 	["li", 20],
//  ["jay", 30],
// ]
func visitLink(sqlExpr *sql_parser.Link) interface{} {
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

// 解释 CASE 语句 [[条件, 结果], [条件, 结果], ..., [nil, 结果]]
// 	 满足条件, 返回对应结果, 如果最后一个条件是 ELSE, 其值保存的是 nil
func visitCase(sqlExpr *sql_parser.Case, instanceDict types.InstanceDict, index int) interface{} {
	for _, caseItem := range sqlExpr.Values {
		condition, val := caseItem.Condition, caseItem.Value
		if condition == nil || visitFunc(condition.(*sql_parser.Func), instanceDict, index).(bool) {
			return visit(val, instanceDict, index)
		}
	}
	return nil
}

// 解释函数, 一般函数保存在 env.Funcs 中, 聚合函数
// 	 以 `函数名+列号` 保存在 env.AggrFuncs 中
// 	 函数参数保存在 Args 属性中, 第一个参数是 DISTINCT, 忽略
func visitFunc(sqlExpr *sql_parser.Func, instanceDict types.InstanceDict, index int) interface{} {
	function := getFunc(index, sqlExpr.Name)

	var funcArgs []interface{}
	for _, _funcArg := range sqlExpr.Args {
		funcArg := visit(_funcArg, instanceDict, index)
		if funcArg == "*" {
			funcArgs = append(funcArgs, funcStar(instanceDict)...)
		} else {
			funcArgs = append(funcArgs, funcArg)
		}
	}
	return function(funcArgs...)
}

// 解释常数, 例如 sqlExpr = SqlToken{100}
func visitConst(sqlExpr *sql_parser.Const) interface{} {
	return sqlExpr.Value
}

// 解释 JOIN ON 语句, 将 ON 部分语句保存在环境中, 返回 JOIN 表结果
// 例如: "JOIN b ON a.id=b.a_id"
// dataSources = {
// 	"a": [
// 	  {"id": 1, "name": "hi"},
//    {"id": 1, "name": "ok"},
//    {"id": 2, "name": "hi"},
//    {"id": 2, "name": "ok"},
// 	],
//  "b": [
//    {"a_id": 1, "other": "ha"},
//    {"a_id": 1, "other": "hh"},
//  ],
//  "c": [
//    {"a_id": 2, "field": "yo"},
//    {"a_id": 2, "field": "ya"},
//  ]
// }
// sqlExpr = SqlToken{JOIN b ON a.id=b.a_id}
// env.HasIndex <= true
// env.Index <= SqlToken{a.id=b.a_id}
// 返回: [
//  "b", [
//    {"a_id": 1, "other": "ha"},
//    {"a_id": 1, "other": "hh"},
//  ],
// ]
func visitInner(sqlExpr *sql_parser.Inner, dataSources types.DataSources) interface{} {
	if sqlExpr.Condition != nil {
		env.HasIndex = true
		env.Index = append(env.Index, sqlExpr.Condition)
	}
	return &types.InnerTable{
		Name:  sqlExpr.Table.Value.Value.(string),
		Table: visit(sqlExpr.Table, dataSources, -1).(types.Records),
	}
}

// 解释变量, 返回变量的真实值
// 例如: sqlExpr = SqlToken{"a.name.first"}
// dataSourcesOrInstanceDict = {
// 	"a": {"id": 1, "name": "hi"},
//  "b": {"a_id": 1, "other": "ha"},
//  "c": {"a_id": 2, "field": "yo"},
// }
// "a" => {"id": 1, "name": "hi"}
// "a.name" => "hi",
// "a.name.first" => "h"
// 返回: "h"
// 再例如: sqlExpr = SqlToken{"a"}
// dataSourcesOrInstanceDict = {
// 	"a": [
// 	  {"id": 1, "name": "hi"},
//    {"id": 1, "name": "ok"},
//    {"id": 2, "name": "hi"},
//    {"id": 2, "name": "ok"},
// 	],
//  "b": [
//    {"a_id": 1, "other": "ha"},
//    {"a_id": 1, "other": "hh"},
//  ],
//  "c": [
//    {"a_id": 2, "field": "yo"},
//    {"a_id": 2, "field": "ya"},
//  ]
// }
// 返回:
// "a": [
// 	  {"id": 1, "name": "hi"},
//    {"id": 1, "name": "ok"},
//    {"id": 2, "name": "hi"},
//    {"id": 2, "name": "ok"},
// 	],
func visitVar(sqlExpr *sql_parser.Var, instanceDict types.InstanceDict) interface{} {
	name := visitConst(sqlExpr.Value).(string)
	if rename, exist := env.ReTableName[name]; exist {
		name = rename
	}
	record := instanceDict[name]
	var ok bool
	// 单一表可以省略表名
	if record == nil && singleTableName != "" {
		record = instanceDict[singleTableName]
		record, ok = recordValue(record, name)
		// 解决没有引号引起的var错误
		if record == nil && !ok {
			return name
		}
	}
	next := sqlExpr.Next
	for next != nil {
		name = visitConst(next.Value).(string)
		record, _ = recordValue(record, name)
		next = next.Next
	}
	return record
}

// 获取表内容
func visitFromVar(sqlExpr *sql_parser.Var, dataSources types.DataSources) interface{} {
	name := visitConst(sqlExpr.Value)
	records := dataSources[name.(string)]
	return records
}

// 解释 FROM 语句, 返回数据源
// 例如: "FROM b, c"
// dataSources = {
// 	"a": [
// 	  {"id": 1, "name": "hi"},
//    {"id": 1, "name": "ok"},
//    {"id": 2, "name": "hi"},
//    {"id": 2, "name": "ok"},
// 	],
//  "b": [
//    {"a_id": 1, "other": "ha"},
//    {"a_id": 1, "other": "hh"},
//  ],
//  "c": [
//    {"a_id": 2, "field": "yo"},
//    {"a_id": 2, "field": "ya"},
//  ]
// }
// sqlExpr = [SqlToken{b}, SqlToken{c}]
// 返回:
// {
//  "b": [
//    {"a_id": 1, "other": "ha"},
//    {"a_id": 1, "other": "hh"},
//  ],
//  "c": [
//    {"a_id": 2, "field": "yo"},
//    {"a_id": 2, "field": "ya"},
//  ]
// }
func visitTables(sqlExpr *sql_parser.Tables, dataSources types.DataSources) types.DataSources {
	var key string
	var ret = make(types.DataSources, 10)
	for _, tableExpr := range sqlExpr.Values {
		records := visit(tableExpr, dataSources, -1)
		switch subExpr := tableExpr.(type) {
		case *sql_parser.Var:
			key = subExpr.Value.Value.(string)
			singleTableName = key // 单一表可以确定表名
		case *sql_parser.As:
			records = visit(subExpr.Value, dataSources, -1)
			key = subExpr.NewName.Value.(string)
			env.ReTableName[subExpr.Value.(*sql_parser.Var).Value.Value.(string)] = key
		case *sql_parser.Inner:
			key = records.(*types.InnerTable).Name
			records = records.(*types.InnerTable).Table
		}
		ret[key] = records.(types.Records)
	}
	return ret
}

func visitSelect(sqlExpr []sql_parser.SqlToken, instanceDict types.InstanceDict) interface{} {
	// * 表达式
	if visit(sqlExpr[1], nil, -1) == true {
		return selectStar(instanceDict)
	}
	var resRecord interface{}
	if useStruct() {
		resRecord = newStruct(types.RecordStructName)
		hadSetStructFieldValue = make(types.UsedKey, 10)
		tmpFieldValue = make(map[string]interface{}, 10)
		for i, subExpr := range sqlExpr[2:] {
			setValue(resRecord, env.FieldName[i], visit(subExpr, instanceDict, i), false)
		}
	} else {
		resRecord = make(map[string]interface{}, len(sqlExpr)-2)
		for i, subExpr := range sqlExpr[2:] {
			resRecord.(map[string]interface{})[env.FieldName[i]] = visit(subExpr, instanceDict, i)
		}
	}
	return resRecord
}

// 解释函数中 * 表达式
func funcStar(instanceDict types.InstanceDict) []interface{} {
	var ret []interface{}
	for _, record := range instanceDict {
		recordKeyValue(record, func(name string, value interface{}) {
			ret = append(ret, value)
		})
	}
	return ret
}

// 获取函数
func getFunc(index int, name string) functions.SqlFunc {
	var function functions.SqlFunc
	if index >= 0 && env.AggPosition[index] {
		if _function, exist := env.AggrFuncs[name+strconv.Itoa(index)]; exist {
			function = _function
		} else {
			function = env.Funcs.Get(name)
		}
	} else {
		function = env.Funcs.Get(name)
	}
	return function
}

// 创建索引, 由 indexExpr、dataSources 条件构造 indexDict
// 例如:
// indexExpr = "FROM a JOIN b ON a.id = b.a_id join c ON a.id = c.a_id" 的
//   条件语法树 [SqlToken{a.id = b.a_id}, SqlToken{a.id = c.a_id}]
// dataSources = {
// 	"a": [
// 	  {"id": 1, "name": "hi"},
//    {"id": 1, "name": "ok"},
//    {"id": 2, "name": "hi"},
//    {"id": 2, "name": "ok"},
// 	],
//  "b": [
//    {"a_id": 1, "other": "ha"},
//    {"a_id": 1, "other": "hh"},
//  ],
//  "c": [
//    {"a_id": 2, "field": "yo"},
//    {"a_id": 2, "field": "ya"},
//  ]
// }
// indexDict <=
// {
// "a": {
// 	  1: [
// 		{"a": {"id": 1, "name": "hi"}},
//		{"a": {"id": 1, "name": "ok"}},
// 	  ],
//    2: [
//  	{"a": {"id": 2, "name": "hi"}},
//		{"a": {"id": 2, "name": "ok"}},
//    ]
//   },
// "b": {
// 	  1: [
// 		{"b": {"a_id": 1, "other": "ha"}},
//		{"b": {"a_id": 1, "other": "hh"}},
// 	  ],
//   },
// "c": {
//    2: [
//  	{"c": {"a_id": 2, "field": "yo"}},
//		{"c": {"a_id": 2, "field": "ya"}},
//    ]
//   },
// }
func makeIndexDict(indexExpr_ interface{}, dataSources types.DataSources, indexDict types.IndexDict) {
	switch indexExpr := indexExpr_.(type) {
	case []*sql_parser.Func:
		for _, subExpr := range indexExpr {
			makeIndexDict(subExpr, dataSources, indexDict)
		}
	case *sql_parser.Func:
		if indexExpr.Name == functions.Keys.Equal {
			makeIndexDict(indexExpr.Args[0], dataSources, indexDict)
			makeIndexDict(indexExpr.Args[1], dataSources, indexDict)
		} else {
			// 不是等号, 放入 where 语句中
			expr.WhereExpr = &sql_parser.Func{
				Name: functions.Keys.And,
				Args: []interface{}{expr.WhereExpr, indexExpr},
			}
		}
	case *sql_parser.Var:
		tableName := visitConst(indexExpr.Value).(string)
		if records, exist := dataSources[tableName]; exist {
			delete(dataSources, tableName)
			indexRecords := make(types.IndexRecords, 10)
			for _, record := range records {
				value, next := record, indexExpr.Next
				for next != nil {
					newName := visitConst(next.Value).(string)
					value, _ = recordValue(value, newName)
					next = next.Next
				}
				if _, ok := indexRecords[value]; !ok {
					indexRecords[value] = []types.Record{}
				}
				indexRecords[value] = append(indexRecords[value], record)
			}
			indexDict[tableName] = indexRecords
		}
	}
}

// 调整重名
func adjustName(newName string, resRecord interface{}, hadAdjust map[string]bool,
	key2table map[string]string, value interface{}) {
	if value2, ok := recordValue(resRecord, newName); ok {
		newName2 := key2table[newName] + "_" + newName
		setValue(resRecord, newName, nil, true)
		adjustName(newName2, resRecord, hadAdjust, key2table, value2)
	}
	setValue(resRecord, newName, value, false)
	hadAdjust[newName] = true
}

// 解释 SELECT 语句中的 * 表达式
func selectStar(instanceDict types.InstanceDict) interface{} {
	var resRecord interface{}
	if useStruct() {
		resRecord = newStruct(types.RecordStructName)
		hadSetStructFieldValue = make(types.UsedKey, 10)
		tmpFieldValue = make(map[string]interface{}, 10)
	} else {
		resRecord = make(map[string]interface{})
	}
	hadAdjust := make(map[string]bool)
	key2table := make(map[string]string)
	hadName := make(map[string]bool)
	for table, record := range instanceDict {
		recordKeyValue(record, func(key string, value interface{}) {
			if value2, ok := recordValue(resRecord, key); ok {
				newName := table + "_" + key
				adjustName(newName, resRecord, hadAdjust, key2table, value)
				// 存在的key未被调整过的状态
				if !hadAdjust[key] {
					newName2 := key2table[key] + "_" + key
					// 先将存在的删除
					setValue(resRecord, key, nil, true)
					adjustName(newName2, resRecord, hadAdjust, key2table, value2)
				}
			} else {
				// 存在key已经是已有字段, 需要调整
				if hadName[key] {
					newName := table + "_" + key
					adjustName(newName, resRecord, hadAdjust, key2table, value)
					// key不是已有字段, 并且key也没有存在的
				} else {
					setValue(resRecord, key, value, false)
					hadName[key] = true
					key2table[key] = table
				}
			}
		})
	}
	return resRecord
}

// 若 SELECT 语句中有聚合函数，则在环境中保存聚合函数，保存相应列位置，保存是否有聚合函数
// 例如:
// token = "SELECT COUNT(a.name), COUNT(a.age) FROM a" 中的
//   选择语法树 [false, false, SqlToken{COUNT(a.name)}, SqlToken{COUNT(a.age)}], 其中
//   列表前两个元素表示是否有 DISTINCT、*
// env.AggPosition <= {0: true, 1: true}
// env.initAggfuncsFieldName <= {"COUNT1": count函数对象, "COUNT2": count函数对象}}
// env.HasAgg <= true
func initAggfuncsFieldName(token interface{}, i int) string {
	switch subExpr := token.(type) {
	case *sql_parser.As:
		initAggfuncsFieldName(subExpr.Value, i)
		return visitConst(subExpr.NewName).(string)
	case *sql_parser.Func:
		// CASE 语句最后一个条件 ELSE 是 nil
		if subExpr == nil {
			return ""
		}
		if env.Funcs.AggregateFuncNameSet[subExpr.Name] {
			env.HasAgg = true
			env.AggPosition[i] = true
			env.AggrFuncs[subExpr.Name+strconv.Itoa(i)] = env.Funcs.Get(subExpr.Name)
		}
		// 防止嵌套函数
		for _, subExpr := range subExpr.Args {
			initAggfuncsFieldName(subExpr, i)
		}
		return subExpr.Name
	case *sql_parser.Var:
		name := visitConst(subExpr.Value).(string)
		next := subExpr.Next
		for next != nil {
			name = visitConst(next.Value).(string)
			next = next.Next
		}
		return name
	case *sql_parser.Case:
		for _, subExpr := range subExpr.Values {
			initAggfuncsFieldName(subExpr, i)
		}
		return "case"
	case *sql_parser.CaseItem:
		initAggfuncsFieldName(subExpr.Condition, i)
		initAggfuncsFieldName(subExpr.Value, i)
	}
	return ""
}

// 预处理
func pretreatment(token []sql_parser.SqlToken) {
	length := len(token) - 2
	env.NameIds = make([]int, length)
	env.FieldName = make([]string, length)
	for j, subSubExpr := range token[2:] {
		fieldName := initAggfuncsFieldName(subSubExpr, j)
		// 非 * 表达式
		if visit(token[1], nil, -1) != true {
			// 结果集字段名冲突
			if _, exist := env.HasNames[fieldName]; exist {
				env.NameIds[j]++
				newFieldName := fieldName + "_" + string(env.NameIds[j])
				env.FieldName[j] = newFieldName
				env.HasNames[newFieldName] = true
			} else {
				env.FieldName[j] = fieldName
				env.HasNames[fieldName] = true
			}
		}
	}
}
