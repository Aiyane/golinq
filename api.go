package golinq

import (
	"github.com/Aiyane/golinq/functions"
	"github.com/Aiyane/golinq/gen"
	"github.com/Aiyane/golinq/interpreters"
	"github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang-collections/collections/queue"

	"github.com/sirupsen/logrus"
	"reflect"
	"runtime"
	"sync"
)

var (
	sqlCache     map[string]*sql_parser.Tree // 缓存
	once         sync.Once
	lock         sync.Mutex
	SelectInterp types.SelectInterp
	InsertInterp types.InsertInterp
	UpdateInterp types.UpdateInterp
	DeleteInterp types.DeleteInterp
	DB2DataBase  map[string]types.DataSources
)

func init() {
	sqlCache = make(map[string]*sql_parser.Tree)

	SelectInterp = &selectInterp{}
	InsertInterp = &insertInterp{}
	UpdateInterp = &updateInterp{}
	DeleteInterp = &deleteInterp{}

	interpreters.SimpleInterp = &interpreters.Simple{}
	interpreters.OrderInterp = &interpreters.Order{}
	interpreters.GroupInterp = &interpreters.Group{}
	interpreters.GroupOrderInterp = &interpreters.GroupOrder{}

	interpreters.SimpleDeleteInterp = &interpreters.SimpleDelete{}
	interpreters.OrderDeleteInterp = &interpreters.OrderDelete{}

	interpreters.SimpleUpdateInterp = &interpreters.SimpleUpdate{}
	interpreters.OrderUpdateInterp = &interpreters.OrderUpdate{}

	interpreters.TagString = "sql"
}

func getSqlTree(sqlExpr string) *sql_parser.Tree {
	inputStream := antlr.NewInputStream(sqlExpr)
	lexer := parser.NewMySqlLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parserRes := parser.NewMySqlParser(stream)
	visitor := sql_parser.NewVisitor()
	tree := parserRes.DmlStatement()
	visitor.Visit(tree.(*parser.DmlStatementContext))
	return visitor.GetTree()
}

func getStack() []byte {
	const size = 4096
	buf := make([]byte, size)
	stackSize := runtime.Stack(buf, false)
	buf = buf[:stackSize]
	return buf
}

// sqlString: select 语句
// dataSources: DataSources 类型对象
// recordStructName: 结果集元素类型名称
// 返回 interface{} 可通过断言获取相应切片类型, 若 recordStructName 传入 "", 则返回值可断言为[]interface{}
func SqlRun(sqlString string, dataSources types.DataSources, recordStructName string) interface{} {
	defer func() {
		lock.Unlock()
		if e := recover(); e != nil {
			panic(e)
			//logrus.Errorf("error=%+v, stack=%s", e, string(getStack()))
		}
	}()
	once.Do(func() {
		lock = sync.Mutex{}
	})
	lock.Lock()
	var sqlTree *sql_parser.Tree
	if tree, exist := sqlCache[sqlString]; exist {
		sqlTree = tree
	} else {
		sqlTree = getSqlTree(sqlString)
		sqlCache[sqlString] = sqlTree
	}
	if sqlTree.InsertStatement != nil {
		return InsertInterp.Interp(sqlString, sqlTree, dataSources, recordStructName)
	} else if sqlTree.UpdateStatement != nil {
		return UpdateInterp.Interp(sqlString, sqlTree, dataSources, recordStructName)
	} else if sqlTree.DeleteStatement != nil {
		return DeleteInterp.Interp(sqlString, sqlTree, dataSources, recordStructName)
	} else if sqlTree.SelectStatementsQueue.Len() > 0 {
		return SelectInterp.Interp(sqlString, sqlTree, dataSources, recordStructName)
	}
	logrus.Infof("warning: no valid sql statement")
	return nil
}

type selectInterp struct{}

func (*selectInterp) Interp(sqlString string, sqlTree *sql_parser.Tree, dataSources types.DataSources,
	recordStructName string) interface{} {
	sqlQueue := sqlTree.SelectStatementsQueue
	sqlCacheQueue := queue.New()
	env := interpreters.NewEnv()
	types.RecordStructName = recordStructName
	for {
		sqlExpr := sqlQueue.Dequeue().(*sql_parser.SelectStatement)
		sqlCacheQueue.Enqueue(sqlExpr)
		tree := sqlExpr.Tree
		result := interpreters.SqlInterpreter(tree, dataSources, env)
		if sqlQueue.Len() == 0 {
			sqlCache[sqlString].SelectStatementsQueue = sqlCacheQueue
			if recordStructName != "" {
				// 获取指针类型
				ptrType := interpreters.TypeRegistry[recordStructName]
				// 结果切片
				resSlice := reflect.MakeSlice(reflect.SliceOf(ptrType), 0, len(result))
				for _, v := range result {
					resSlice = reflect.Append(resSlice, reflect.ValueOf(v))
				}
				return resSlice.Interface()
			}
			return result
		}
		env.Link[sqlExpr.Id] = result
	}
}

type insertInterp struct{}

func (*insertInterp) Interp(sqlString string, sqlTree *sql_parser.Tree, dataSources types.DataSources,
	recordStructName string) interface{} {
	tree := sqlTree.InsertStatement
	if sqlTree.InsertStatement.Link != nil {
		res := SelectInterp.Interp(sqlString, sqlTree, dataSources, "").([]interface{})
		return interpreters.InsertInterp(tree, dataSources, res)
	}
	return interpreters.InsertInterp(tree, dataSources, nil)
}

type updateInterp struct{}

func (*updateInterp) Interp(sqlString string, sqlTree *sql_parser.Tree, dataSources types.DataSources,
	recordStructName string) interface{} {
	return interpreters.UpdateInterp(sqlTree.UpdateStatement, dataSources)
}

type deleteInterp struct{}

func (*deleteInterp) Interp(sqlString string, sqlTree *sql_parser.Tree, dataSources types.DataSources,
	recordStructName string) interface{} {
	return interpreters.DeleteInterp(sqlTree.DeleteStatement, dataSources)
}

// 通过结构体列表构建 DataSources 类型
func NewDataSources(_dataSources ...interface{}) types.DataSources {
	dataSources := make(types.DataSources, len(_dataSources))
	for _, records := range _dataSources {
		v := reflect.ValueOf(records)
		l := v.Len()
		data := make(types.Records, 0, l)
		for i := 0; i < l; i++ {
			data = append(data, v.Index(i).Interface())
		}
		dataSources[v.Index(0).Interface().(types.Entitier).TableName()] = data
	}
	return dataSources
}

// 注册结果集/原数据集元素类型
// elem = (*T)(nil)
func RegisterType(elem interface{}) {
	t := reflect.TypeOf(elem)
	table := reflect.New(t.Elem()).Interface()
	switch v := table.(type) {
	case types.Entitier:
		interpreters.TypeRegistry[v.TableName()] = t
	default:
		e := t.Elem()
		interpreters.TypeRegistry[e.Name()] = t
	}
}

// 清空注册的类型
func ClearType() {
	interpreters.TypeRegistry = make(map[string]reflect.Type, 0)
}

// 注册内建函数
func RegisterFunc(key string, f functions.SqlClosureFunc) {
	functions.AddFunc(key, f)
}

// 注册 db 对象与数据库映射
func RegisterDB2DataBase(dB2DataBase map[string]types.DataSources) {
	DB2DataBase = dB2DataBase
}
