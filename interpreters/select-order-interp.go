package interpreters

import (
	"github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
	"github.com/sirupsen/logrus"
)

type Order struct{}

func (*Order) Interp(dataSources types.DataSources) {
	/*
		无 group 有 Order 语句
	*/
	// 排序规则
	orders := getSelectOrders()

	// 如果使用 join on 索引优化
	if env.HasIndex {
		orderIndexLoop(dataSources, orders)
	} else {
		orderMainLoop(dataSources, make(types.InstanceDict, 10), orders)
	}

	// select 语句, 因为是有序插入, 需要做比较, 所以不能提前 select, 如果提前 select 会导致之前的记录没有做比较的字段
	newRes := make([]interface{}, 0, 100)
	for _, instanceDict := range res {
		newRes = append(newRes, visitSelect(expr.SelectExpr, instanceDict.(types.InstanceDict)))
	}
	res = newRes
}

// 主循环
func orderMainLoop(dataSources types.DataSources, instanceDict types.InstanceDict, orders []*types.OrderFields) {
	orderMainIterator(make(types.UsedKey, 10), 0, len(dataSources), instanceDict, dataSources, orders)
}

// 主迭代器
func orderMainIterator(usedKey types.UsedKey, num int, dataNum int, instanceDict types.InstanceDict, dataSources types.DataSources,
	orders []*types.OrderFields) {
	if num == dataNum {
		copyInstanceDict := deepCopyInstanceDict(instanceDict, "", nil)
		orderWhereFilter(copyInstanceDict, orders)
	} else {
		// 取一个表
		for name, records := range dataSources {
			if _, exist := usedKey[name]; !exist {
				usedKey[name] = true
				for _, record := range records {
					instanceDict[name] = record
					orderMainIterator(usedKey, num+1, dataNum, instanceDict, dataSources, orders)
					delete(instanceDict, name)
				}
				delete(usedKey, name)
				return
			}
		}
	}
}

// 索引主循环
func orderIndexLoop(dataSources types.DataSources, orders []*types.OrderFields) {
	var indexDict = make(types.IndexDict, 10)
	makeIndexDict(env.Index, dataSources, indexDict)
	name, indexRecords := popItem(indexDict)
	indexLen := len(indexDict)
	hasDataSources := len(dataSources) > 0
	for indexKey, records := range indexRecords {
		orderIndexIterator(hasDataSources, indexKey, make(types.UsedKey, 10), indexDict,
			dataSources, make(types.InstanceDict, 10), 0, indexLen, name, records, orders)
	}
}

// 索引迭代器
func orderIndexIterator(hasDataSources bool, indexKey interface{}, usedKey types.UsedKey, indexDict types.IndexDict,
	dataSources types.DataSources, instanceDict types.InstanceDict, num int, indexLen int, tableName string, records types.Records,
	orders []*types.OrderFields) {
	if num == indexLen {
		for _, record := range records {
			copyInstanceDict := deepCopyInstanceDict(instanceDict, tableName, record)
			if hasDataSources {
				orderMainLoop(dataSources, copyInstanceDict, orders)
			} else {
				orderWhereFilter(copyInstanceDict, orders)
			}
		}
	} else {
		for name, indexRecords := range indexDict {
			if _, exist := usedKey[name]; !exist {
				usedKey[name] = true
				if records2, exist2 := indexRecords[indexKey]; exist2 {
					for _, record := range records2 {
						instanceDict[name] = record
						orderIndexIterator(hasDataSources, indexKey, usedKey, indexDict, dataSources,
							instanceDict, num+1, indexLen, tableName, records, orders)
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
func orderWhereFilter(instanceDict types.InstanceDict, orders []*types.OrderFields) {
	if expr.WhereExpr == nil || visitFunc(expr.WhereExpr, instanceDict, -1) == true {
		index := bisectLeft(res, instanceDict, orders)
		// 插入 index 位置
		if index == len(res) {
			res = append(res, instanceDict)
		} else {
			res = append(res[:index+1], res[index:]...)
			res[index] = instanceDict
		}
	}
}

func getSelectOrders() []*types.OrderFields {
	orders := make([]*types.OrderFields, 0, 10)
	for _, subExpr := range expr.OrderExpr {
		orders = append(orders, getOrderList(subExpr))
	}
	return orders
}

func getOrderList(sqlExpr interface{}) *types.OrderFields {
	var orderFields *types.OrderFields
	switch subExpr := sqlExpr.(type) {
	case *sql_parser.Var:
		orderFields = &types.OrderFields{
			Order:  types.PositiveOrder,
			Fields: getVar(subExpr),
		}
	case *sql_parser.Desc:
		orderFields = &types.OrderFields{
			Order:  types.ReverseOrder,
			Fields: getVar(subExpr.Value),
		}
	default:
		logrus.Errorf("[getOrderList] type error: sqlExpr[%T]:[%v]", sqlExpr, sqlExpr)
		panic("")
	}
	return orderFields
}

func getVar(sqlExpr *sql_parser.Var) []string {
	res := make([]string, 0, 10)
	for sqlExpr != nil {
		res = append(res, visitConst(sqlExpr.Value).(string))
		sqlExpr = sqlExpr.Next
	}
	return res
}

// 获取插入位置
func bisectLeft(a []interface{}, x types.InstanceDict, orders []*types.OrderFields) int {
	hi := len(a)
	lo := 0
	for lo < hi {
		mid := (lo + hi) / 2
		if compareRecord(a[mid].(types.InstanceDict), x, orders) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// 记录比较函数
func compareRecord(item1 types.InstanceDict, item2 types.InstanceDict, orders []*types.OrderFields) bool {
	for _, condition := range orders {
		order := condition.Order
		var v1, v2 interface{}
		for i, key := range condition.Fields {
			if i == 0 {
				v1 = item1[key]
				v2 = item2[key]
				// 单一表省略了表名
				if v1 == nil && v2 == nil {
					v1 = item1[singleTableName].(types.Record)
					v2 = item2[singleTableName].(types.Record)
					v1, _ = recordValue(v1, key)
					v2, _ = recordValue(v2, key)
				}
			} else {
				v1, _ = recordValue(v1, key)
				v2, _ = recordValue(v2, key)
			}
		}
		compareRes := compareValue(v1, v2)
		if compareRes < 0 {
			if order == types.PositiveOrder {
				return true
			}
			return false
		} else if compareRes > 0 {
			if order == types.PositiveOrder {
				return false
			}
			return true
		}
	}
	return true
}
