package types

import (
	"github.com/Aiyane/golinq/functions"
	parser "github.com/Aiyane/golinq/sql-parser"
)

var RecordStructName string // 结果集类型名

// {"age": 30, "name": "jay"} | {}
type Record interface{}

// [
//     {"age": 30, "name": "jay"},
//     {"age": 30, "name": "jay"}
// ]
type Records []Record

// {
//   "teacher": [
//       {"age": 30, "name": "jay"},
//       {"age": 30, "name": "jay"}
//   ]
// }
type InnerTable struct {
	Name  string
	Table Records
}

// {
//   "Teacher": {"age": 30, "name": "jay"},
//   "Student": {"class": "group1", "number": 20}
// }
type InstanceDict map[string]Record

// [instanceDict1, instanceDict2]
type InstanceDicts []InstanceDict

// {
//   "Teacher": [
//      {"age": 30, "name": "jay"},
//      {"age": 30, "name": "jay"}
//  ],
// }
type DataSources map[string]Records

// {
//   [1, "key1", "key2"]: [instanceDict1, instanceDict2],
//   [2, "key3", "key4"]: [instanceDict3, instanceDict4]
// } 或
// {[1, "key1", "key2"]: [instanceDict1, instanceDict2]}
type GroupDataSources map[string]InstanceDicts

// 30|"xx"|true
type IndexKey interface{}

// {
//    30: [{"age": 30, "name": "jay"}, {"age": 30, "name": "jay"}],
// }
type IndexRecords map[IndexKey]Records

// {
//   "Teacher": {
//    30: [{"age": 30, "name": "jay"}, {"age": 30, "name": "jay"}],
//    20: [{"age": 20, "name": "wang"}, {"age": 20, "name": "li"}],
//   },
// }
type IndexDict map[string]IndexRecords
type UsedKey map[string]bool

type Order int

const (
	// 正序
	PositiveOrder Order = 0
	// 倒序
	ReverseOrder Order = 1
)

// 有序字段
type OrderFields struct {
	Order  Order    // 顺序
	Fields []string // 字段
}

type Env struct {
	Funcs       *functions.FuncProxy
	Index       []*parser.Func
	AggrFuncs   map[string]functions.SqlFunc
	AggPosition map[int]bool
	HasAgg      bool
	HasIndex    bool
	Link        map[int][]interface{}
	LinkCache   map[int]interface{}
	GroupList   []interface{}
	GroupMap    map[string]bool
	FieldName   []string
	HasNames    map[string]bool
	NameIds     []int
	ReTableName map[string]string
}
