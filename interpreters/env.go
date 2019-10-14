package interpreters

import (
	"github.com/Aiyane/golinq/functions"
	"github.com/Aiyane/golinq/sql-parser"
	"github.com/Aiyane/golinq/types"
	"reflect"
)

var (
	singleTableName        string                          // 在能确定的单一表执行情况下保存的表名
	env                    *types.Env                      // 环境
	res                    []interface{}                   // 结果集
	expr                   *sql_parser.EXPR                // 语句
	TypeRegistry           = make(map[string]reflect.Type) // 类型字典
	hadSetStructFieldValue types.UsedKey                   // 是否设置过 struct 属性值
	tmpFieldValue          map[string]interface{}          // 调整重名时保存临时字段名的值
	preFieldName           map[string]struct{}             // 前一个select语句的 select 字段名
)

// 生成解释环境
func NewEnv() *types.Env {
	return &types.Env{
		// 将内置函数代理放入环境中
		Funcs:       functions.NewFuncProxy(),
		Index:       make([]*sql_parser.Func, 0, 10),
		AggrFuncs:   make(map[string]functions.SqlFunc, 10),
		AggPosition: make(map[int]bool, 10),
		Link:        make(map[int][]interface{}, 3),
		LinkCache:   make(map[int]interface{}, 3),
		HasAgg:      false,
		HasIndex:    false,
		GroupList:   make([]interface{}, 0, 10),
		GroupMap:    make(map[string]bool, 10),
		HasNames:    make(map[string]bool, 10),
		ReTableName: make(map[string]string, 10),
	}
}

func clearEnv() {
	env.HasAgg = false
	env.AggrFuncs = make(map[string]functions.SqlFunc, 10)
	env.HasIndex = false
	env.GroupMap = make(map[string]bool, 10)
	env.GroupList = make([]interface{}, 0, 10)
	env.Index = make([]*sql_parser.Func, 0, 10)
	env.AggPosition = make(map[int]bool, 10)
	env.HasNames = make(map[string]bool, 10)
	env.ReTableName = make(map[string]string, 10)
	// 设置前一个select语句的 select 字段名, 为Link取值用
	preFieldName = make(map[string]struct{}, len(env.FieldName))
	for _, name := range env.FieldName {
		preFieldName[name] = struct{}{}
	}
}
