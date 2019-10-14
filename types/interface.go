package types

import parser "github.com/Aiyane/golinq/sql-parser"

type Entitier interface {
	TableName() string
}

type Interp interface {
	Interp(sqlString string, sqlTree *parser.Tree, dataSources DataSources, recordStructName string) interface{}
}

type SelectInterp interface {
	Interp
}

type InsertInterp interface {
	Interp
}

type UpdateInterp interface {
	Interp
}

type DeleteInterp interface {
	Interp
}

type SelectSubInterp interface {
	Interp(dataSources DataSources)
}

type SimpleInterp interface {
	SelectSubInterp
}

type OrderInterp interface {
	SelectSubInterp
}

type GroupInterp interface {
	SelectSubInterp
}

type GroupOrderInterp interface {
	SelectSubInterp
}

type DeleteSubInterp interface {
	Interp(deleteExpr *parser.DeleteExpr, records Records) (Records, int)
}

type SimpleDeleteInterp interface {
	DeleteSubInterp
}

type OrderDeleteInterp interface {
	DeleteSubInterp
}

type UpdateSubInterp interface {
	Interp(updateExpr *parser.UpdateExpr, table string, records Records) (Records, int)
}

type SimpleUpdateInterp interface {
	UpdateSubInterp
}

type OrderUpdateInterp interface {
	UpdateSubInterp
}
