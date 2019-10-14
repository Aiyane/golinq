package interpreters

import (
	"github.com/Aiyane/golinq/functions"
	sql_parser "github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
)

// insertExpr 插入语句
// dataSources 源数据
// records select 语句的结果集
func InsertInterp(insertExpr *sql_parser.InsertExpr, dataSources types.DataSources, records []interface{}) int {
	env = NewEnv()
	table := insertExpr.Table
	maxId, _ := getIdSet(dataSources[table])
	_columns := insertExpr.Columns
	columns := make([]string, 0, len(_columns))
	hasId := false
	for _, column := range _columns {
		columnName := visit(column, nil, -1).(string)
		if columnName == "id" {
			hasId = true
		}
		columns = append(columns, columnName)
	}
	var values interface{} = insertExpr.Values
	if records != nil {
		values = records
	}
	_, isStruct := TypeRegistry[table]
	tmpFieldValue = make(map[string]interface{}, 10)
	hadSetStructFieldValue = make(types.UsedKey, 10)
	insertRecordsIterator(values, func(record interface{}) {
		switch r := record.(type) {
		case map[string]interface{}:
			if isStruct {
				resRecord := newStruct(table)
				for _, column := range columns {
					if column == "id" {
						maxId = r[column]
					}
					setValue(resRecord, column, r[column], false)
				}
				if !hasId {
					newId := functions.Add(maxId, 1)
					setValue(resRecord, "id", newId, false)
					maxId = newId
				}
				dataSources[table] = append(dataSources[table], resRecord)
			} else {
				resRecord := make(map[string]interface{})
				for _, column := range columns {
					if column == "id" {
						maxId = r[column]
					}
					resRecord[column] = r[column]
				}
				if !hasId {
					newId := functions.Add(maxId, 1)
					resRecord["id"] = newId
					maxId = newId
				}
				dataSources[table] = append(dataSources[table], resRecord)
			}
		case []sql_parser.SqlToken:
			if isStruct {
				resRecord := newStruct(table)
				for i, column := range columns {
					if column == "id" {
						maxId = visit(r[i], nil, -1)
					}
					setValue(resRecord, column, visit(r[i], nil, -1), false)
				}
				if !hasId {
					newId := functions.Add(maxId, 1)
					setValue(resRecord, "id", newId, false)
					maxId = newId
				}
				dataSources[table] = append(dataSources[table], resRecord)
			} else {
				resRecord := make(map[string]interface{})
				for i, column := range columns {
					if column == "id" {
						maxId = visit(r[i], nil, -1)
					}
					resRecord[column] = visit(r[i], nil, -1)
				}
				if !hasId {
					newId := functions.Add(maxId, 1)
					resRecord["id"] = newId
					maxId = newId
				}
				dataSources[table] = append(dataSources[table], resRecord)
			}
		}
	})
	return functions.ToInt(maxId).(int)
}

func insertRecordsIterator(records interface{}, fn func(interface{})) {
	switch r := records.(type) {
	case [][]sql_parser.SqlToken:
		for _, record := range r {
			fn(record)
		}
	case []interface{}:
		for _, record := range r {
			fn(record)
		}
	}
}

func getIdSet(records types.Records) (maxId interface{}, idSet map[interface{}]struct{}){
	idSet = make(map[interface{}]struct{}, len(records))
	maxId = 0
	for _, record := range records {
		id, _ := recordValue(record, "id")
		if functions.Greater(id, maxId) {
			maxId = id
		}
		idSet[id] = struct{}{}
	}
	return maxId, idSet
}
