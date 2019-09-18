package interpreters

import (
	sql_parser "github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
)

var (
	SimpleUpdateInterp types.SimpleUpdateInterp
	OrderUpdateInterp  types.OrderUpdateInterp
)

func UpdateInterp(updateExpr *sql_parser.UpdateExpr, dataSources types.DataSources) int {
	env = NewEnv()
	var n int
	table, records := updateTable(updateExpr.Table, dataSources)
	singleTableName = table // 单一表名
	if orderExpr := updateExpr.Order; orderExpr != nil {
		records, n = OrderUpdateInterp.Interp(updateExpr, table, records)
	} else {
		records, n = SimpleUpdateInterp.Interp(updateExpr, table, records)
	}
	dataSources[table] = records
	return n
}

type SimpleUpdate struct{}

func (*SimpleUpdate) Interp(updateExpr *sql_parser.UpdateExpr, table string, records types.Records) (types.Records, int) {
	whereExpr := updateExpr.Where
	total := len(records)
	ret := make(types.Records, 0, total)
	var limit, offset = -1, -1
	if limitExpr := updateExpr.Limit; limitExpr != nil {
		offset = visitConst(limitExpr.Offset).(int)
		limit = visitConst(limitExpr.Limit).(int)
	}
	noLimit := limit == -1
	updateCount := 0
	var ok bool
	values := updateExpr.Values
	for _, record := range records {
		instanceDict := types.InstanceDict{table: record}
		if whereExpr == nil || visitFunc(whereExpr, instanceDict, -1) == true {
			if ok, limit, offset = canUpdate(noLimit, limit, offset); ok {
				for _, columnValue := range values {
					names := getVar(columnValue.Column)
					name := names[len(names)-1]
					value := visit(columnValue.Value, instanceDict, -1)
					setValue(record, name, value, false)
				}
				updateCount += 1
			}
		}
		ret = append(ret, record)
	}
	return ret, updateCount
}

func canUpdate(noLimit bool, hadLimit int, hadOffset int) (bool, int, int) {
	if noLimit {
		return true, 0, 0
	}
	if hadLimit == 0 && hadOffset > 0 {
		return true, hadLimit, hadOffset - 1
	}
	return false, hadLimit, hadOffset
}

func updateTable(tableExpr sql_parser.SqlToken, dataSources types.DataSources) (string, types.Records) {
	var key string
	records := visit(tableExpr, dataSources, -1)
	switch subExpr := tableExpr.(type) {
	case *sql_parser.Var:
		key = subExpr.Value.Value.(string)
	case *sql_parser.As:
		records = visit(subExpr.Value, dataSources, -1)
		rename := subExpr.NewName.Value.(string)
		env.ReTableName[subExpr.Value.(*sql_parser.Var).Value.Value.(string)] = rename
	}
	return key, records.(types.Records)
}
