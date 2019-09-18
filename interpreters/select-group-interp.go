package interpreters

import "github.com/Aiyane/golinq/types"

type Group struct{}

func (*Group) Interp(dataSources types.DataSources) {
	/*
		无 order 语句
	*/
	// group by 聚合
	groupDataSources := getGroupDataSources(dataSources)

	// having 筛选
	for pk, instanceDicts := range groupDataSources {
		havingFilter(types.GroupDataSources{pk: instanceDicts})
	}
}

func havingFilter(groupDataSources types.GroupDataSources) {
	if expr.HavingExpr == nil || groupVisit(expr.HavingExpr, groupDataSources, -1) == true {
		res = append(res, groupVisitSelect(expr.SelectExpr, groupDataSources))
	}
}
