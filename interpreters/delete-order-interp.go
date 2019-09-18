package interpreters

import (
	sql_parser "github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
)

type OrderDelete struct{}

func (*OrderDelete) Interp(deleteExpr *sql_parser.DeleteExpr, records types.Records) (types.Records, int) {
	// 初始化结果集
	res = make([]interface{}, 0, 100)
	orders := getOrders(deleteExpr.Order)
	total := len(records)
	needDeleteIndex := make([]int, 0, total)
	whereExpr := deleteExpr.Where
	for i, record := range records {
		instanceDict := types.InstanceDict{deleteExpr.Table: record}
		if whereExpr == nil || visitFunc(whereExpr, instanceDict, -1) == true {
			index := bisectLeft(res, instanceDict, orders)
			// 插入 index 位置
			if index == len(res) {
				res = append(res, instanceDict)
				needDeleteIndex = append(needDeleteIndex, i)
			} else {
				res = append(res[:index+1], res[index:]...)
				res[index] = instanceDict
				needDeleteIndex = append(needDeleteIndex[:index+1], needDeleteIndex[index:]...)
				needDeleteIndex[index] = i
			}
		}
	}
	if deleteExpr := deleteExpr.Limit; deleteExpr != nil {
		needDeleteIndex = deleteOrUpdateLimit(deleteExpr, needDeleteIndex)
	}
	deleteLen := len(needDeleteIndex)
	return deleteRecords(total-deleteLen, needDeleteIndex, records), deleteLen
}

func getOrders(expr []sql_parser.SqlToken) []*types.OrderFields {
	orders := make([]*types.OrderFields, 0, 10)
	for _, subExpr := range expr {
		orders = append(orders, getOrderList(subExpr))
	}
	return orders
}
