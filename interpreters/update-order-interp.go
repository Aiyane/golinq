package interpreters

import (
	sql_parser "github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
)

type OrderUpdate struct{}

func (*OrderUpdate) Interp(updateExpr *sql_parser.UpdateExpr, table string, records types.Records) (types.Records, int) {
	whereExpr := updateExpr.Where
	res = make([]interface{}, 0, 100)
	total := len(records)
	orders := getOrders(updateExpr.Order)
	needUpdateIndex := make([]int, 0, total)
	updateCount := 0
	for i, record := range records {
		instanceDict := types.InstanceDict{table: record}
		if whereExpr == nil || visitFunc(whereExpr, instanceDict, -1) == true {
			index := bisectLeft(res, instanceDict, orders)
			// 插入 index 位置
			if index == len(res) {
				res = append(res, instanceDict)
				needUpdateIndex = append(needUpdateIndex, i)
			} else {
				res = append(res[:index+1], res[index:]...)
				res[index] = instanceDict
				needUpdateIndex = append(needUpdateIndex[:index+1], needUpdateIndex[index:]...)
				needUpdateIndex[index] = i
			}
		}
	}
	if limitExpr := updateExpr.Limit; limitExpr != nil {
		needUpdateIndex = deleteOrUpdateLimit(limitExpr, needUpdateIndex)
	}
	values := updateExpr.Values
	for _, index := range needUpdateIndex {
		record := records[index]
		instanceDict := types.InstanceDict{table: record}
		for _, columnValue := range values {
			names := getVar(columnValue.Column)
			name := names[len(names)-1]
			value := visit(columnValue.Value, instanceDict, -1)
			setValue(record, name, value, false)
		}
		records[index] = record
		updateCount += 1
	}
	return records, updateCount
}
