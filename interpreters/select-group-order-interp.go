package interpreters

import "github.com/Aiyane/golinq/types"

type GroupOrder struct{}

func (*GroupOrder) Interp(dataSources types.DataSources) {
	// 排序规则
	orders := getSelectOrders()

	// group by 聚合
	groupDataSources := getGroupDataSources(dataSources)

	// having 筛选
	for pk, instanceDicts := range groupDataSources {
		orderHavingFilter(types.GroupDataSources{pk: instanceDicts}, orders)
	}

	// select 语句
	newRes := make([]interface{}, 0, 10)
	for _, subRes := range res {
		newRes = append(newRes, groupVisitSelect(expr.SelectExpr, subRes.(types.GroupDataSources)))
	}
	res = newRes
}

func orderHavingFilter(groupDataSources types.GroupDataSources, orders []*types.OrderFields) {
	if expr.HavingExpr == nil || groupVisit(expr.HavingExpr, groupDataSources, -1) == true {
		index := groupBisectLeft(res, groupDataSources, orders)
		// 插入 index 位置
		if index == len(res) {
			res = append(res, groupDataSources)
		} else {
			res = append(res[:index+1], res[index:]...)
			res[index] = groupDataSources
		}
	}
}

func groupBisectLeft(a []interface{}, x types.GroupDataSources, orders []*types.OrderFields) int {
	hi := len(a)
	lo := 0
	for lo < hi {
		mid := (lo + hi) / 2
		if groupCompareRecord(a[mid].(types.GroupDataSources), x, orders) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func groupCompareRecord(item1 types.GroupDataSources, item2 types.GroupDataSources, orders []*types.OrderFields) bool {
	for _, condition := range orders {
		order := condition.Order
		for _, instanceDicts := range item1 {
			for _, instanceDicts2 := range item2 {
				var v1, v2 interface{}
				for i, key := range condition.Fields {
					if i == 0 {
						v1 = instanceDicts[0][key]
						v2 = instanceDicts2[0][key]
						// 单一表省略了表名
						if v1 == nil && v2 == nil {
							v1 = instanceDicts[0][singleTableName].(types.Record)
							v2 = instanceDicts2[0][singleTableName].(types.Record)
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
				break
			}
			break
		}
	}
	return true
}
