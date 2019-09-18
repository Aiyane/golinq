package interpreters

import (
	sql_parser "github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
)

var (
	SimpleDeleteInterp types.SimpleDeleteInterp
	OrderDeleteInterp  types.OrderDeleteInterp
)

func DeleteInterp(deleteExpr *sql_parser.DeleteExpr, dataSources types.DataSources) int {
	env = NewEnv()
	table := deleteExpr.Table
	singleTableName = table // 单一表名
	records := dataSources[table]
	var n int
	if deleteExpr.Order != nil {
		records, n = OrderDeleteInterp.Interp(deleteExpr, records)
	} else {
		records, n = SimpleDeleteInterp.Interp(deleteExpr, records)
	}
	dataSources[table] = records
	return n
}

type SimpleDelete struct{}

func (*SimpleDelete) Interp(deleteExpr *sql_parser.DeleteExpr, records types.Records) (types.Records, int) {
	total := len(records)
	needDeleteIndex := make([]int, 0, total)
	whereExpr := deleteExpr.Where
	for i, record := range records {
		if whereExpr == nil || visitFunc(whereExpr, types.InstanceDict{deleteExpr.Table: record}, -1) == true {
			needDeleteIndex = append(needDeleteIndex, i)
		}
	}
	if deleteExpr := deleteExpr.Limit; deleteExpr != nil {
		needDeleteIndex = deleteOrUpdateLimit(deleteExpr, needDeleteIndex)
	}
	deleteLen := len(needDeleteIndex)
	return deleteRecords(total-deleteLen, needDeleteIndex, records), deleteLen
}

func deleteRecords(retLen int, needDeleteIndex []int, records types.Records) types.Records {
	ret := make(types.Records, 0, retLen)
	subNum := 0
	for _, i := range needDeleteIndex {
		index := i - subNum
		ret = append(ret, records[:index]...)
		records = records[index+1:]
		subNum += index + 1
	}
	ret = append(ret, records...)
	return ret
}

func deleteOrUpdateLimit(sqlExpr *sql_parser.Limit, needDeleteIndex []int) []int {
	offset := visitConst(sqlExpr.Offset).(int)
	limit := visitConst(sqlExpr.Limit).(int)
	length := len(needDeleteIndex)
	if offset >= length {
		return []int{}
	}
	end := offset + limit
	if end > length {
		end = length
	}
	return needDeleteIndex[offset:end]
}
