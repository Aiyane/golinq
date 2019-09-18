package interpreters

import (
	"github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
)

var (
	SimpleInterp     types.SimpleInterp
	OrderInterp      types.OrderInterp
	GroupInterp      types.GroupInterp
	GroupOrderInterp types.GroupOrderInterp
)

func SqlInterpreter(tree *sql_parser.State, dataSources types.DataSources, newEnv *types.Env) []interface{} {
	// 初始化环境
	env = newEnv
	// 初始化结果集
	res = make([]interface{}, 0, 100)
	// 初始化表达式
	expr = sql_parser.NewExpr(tree)
	selectSqlInterpreterExec(dataSources)
	// 清理环境
	clearEnv()
	return res
}

func selectSqlInterpreterExec(dataSources types.DataSources) {
	// 如果有 from 执行 from 筛选
	if expr.FromExpr != nil {
		dataSources = visitTables(expr.FromExpr, dataSources)
		if len(dataSources) > 1 {
			singleTableName = "" // 非单一表删除表名
		}
	}
	// 通过 select 语句先确定有哪些聚合函数，每一列的都需要新生成一个函数
	pretreatment(expr.SelectExpr)
	// 选择解释器
	if expr.GroupExpr != nil {
		if expr.OrderExpr != nil {
			GroupOrderInterp.Interp(dataSources)
		} else {
			GroupInterp.Interp(dataSources)
		}
	} else if expr.OrderExpr != nil {
		OrderInterp.Interp(dataSources)
	} else {
		SimpleInterp.Interp(dataSources)
	}
	// 如果有聚合函数
	if env.HasAgg {
		if length := len(res); length > 0 {
			res = []interface{}{res[length-1]}
		} else {
			res = []interface{}{}
		}
	}
	// 如果有 limit 语句
	if expr.LimitExpr != nil {
		visitLimit(expr.LimitExpr)
	}
}
