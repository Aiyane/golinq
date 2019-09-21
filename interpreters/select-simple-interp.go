package interpreters

import "github.com/Aiyane/golinq/types"

type Simple struct{}

// 解释顶层接口
func (*Simple) Interp(dataSources types.DataSources) {
	/*
		无 group 与 order 语句
	*/
	// 如果使用 join on 索引优化
	if env.HasIndex {
		simpleIndexLoop(dataSources)
	} else {
		simpleMainLoop(dataSources, make(types.InstanceDict, 10))
	}
}

// 主循环
func simpleMainLoop(dataSources types.DataSources, instanceDict types.InstanceDict) {
	simpleMainIterator(make(types.UsedKey, 10), 0, len(dataSources), instanceDict, dataSources)
}

// 主迭代器
func simpleMainIterator(usedKey types.UsedKey, num int, dataLen int, instanceDict types.InstanceDict,
	dataSources types.DataSources) {
	if num == dataLen {
		copyInstanceDict := deepCopyInstanceDict(instanceDict, "", nil)
		simpleWhereFilter(copyInstanceDict)
	} else {
		// 取一个表
		for name, records := range dataSources {
			if _, ok := usedKey[name]; !ok {
				usedKey[name] = struct{}{}
				for _, record := range records {
					instanceDict[name] = record
					simpleMainIterator(usedKey, num+1, dataLen, instanceDict, dataSources)
					delete(instanceDict, name)
				}
				delete(usedKey, name)
				return
			}
		}
	}
}

// 索引主循环
func simpleIndexLoop(dataSources types.DataSources) {
	var indexDict = make(types.IndexDict)
	makeIndexDict(env.Index, dataSources, indexDict)
	name, indexRecords := popItem(indexDict)
	indexLen := len(indexDict)
	hasDataSources := len(dataSources) > 0
	for indexKey, records := range indexRecords {
		simpleIndexIterator(hasDataSources, indexKey, make(types.UsedKey, 10), indexDict,
			dataSources, make(types.InstanceDict, 10), 0, indexLen, name, records)
	}
}

// 索引迭代器
func simpleIndexIterator(hasDataSources bool, indexKey interface{}, usedKeys types.UsedKey, indexDict types.IndexDict,
	dataSources types.DataSources, instanceDict types.InstanceDict, num int, indexLen int, tableName string, records types.Records) {
	if num == indexLen {
		for _, record := range records {
			copyInstanceDict := deepCopyInstanceDict(instanceDict, tableName, record)
			if hasDataSources {
				simpleMainLoop(dataSources, copyInstanceDict)
			} else {
				simpleWhereFilter(copyInstanceDict)
			}
		}
	} else {
		for name, indexRecords := range indexDict {
			if _, ok := usedKeys[name]; !ok {
				usedKeys[name] = struct{}{}
				if records2, exist2 := indexRecords[indexKey]; exist2 {
					for _, record := range records2 {
						instanceDict[name] = record
						simpleIndexIterator(hasDataSources, indexKey, usedKeys, indexDict, dataSources,
							instanceDict, num+1, indexLen, tableName, records)
						delete(instanceDict, name)
					}
				}
				delete(usedKeys, name)
				return
			}
		}
	}
}

// 筛选器
func simpleWhereFilter(instanceDict types.InstanceDict) {
	if expr.WhereExpr == nil || visitFunc(expr.WhereExpr, instanceDict, -1) == true {
		res = append(res, visitSelect(expr.SelectExpr, instanceDict))
	}
}
