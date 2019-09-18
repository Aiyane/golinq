package functions

import "sync"

var (
	funcProxy       *FuncProxy
	once            sync.Once
	externalMethods = make(map[string]SqlClosureFunc, 10)
)

type SqlFunc func(...interface{}) interface{}
type SqlClosureFunc func() SqlFunc
type GroupRecord []interface{}

type FuncProxy struct {
	AggregateFuncNameSet map[string]bool
	SqlFunc              map[string]SqlClosureFunc
}

func NewFuncProxy() *FuncProxy {
	once.Do(func() {
		funcProxy = &FuncProxy{
			AggregateFuncNameSet: map[string]bool{
				Keys.Sum:   true,
				Keys.Count: true,
			},
			SqlFunc: map[string]SqlClosureFunc{
				Keys.First:             first,
				Keys.Not:               not,
				Keys.SignalNot:         signalNot,
				Keys.SignalExclamation: signalExclamation,
				Keys.SignalAdd:         signalAdd,
				Keys.SignalSub:         signalSub,
				Keys.DoubleSub:         doubleSub,
				Keys.Add:               add,
				Keys.Sub:               sub,
				Keys.Mul:               mul,
				Keys.Div:               div,
				Keys.Res:               res,
				Keys.Exist:             exist,
				Keys.Is:                is,
				Keys.IsNot:             isNot,
				Keys.In:                in,
				Keys.NotIn:             notIn,
				Keys.Or:                or,
				Keys.And:               and,
				Keys.Equal:             equal,
				Keys.GreaterEqual:      greaterEqual,
				Keys.LessEqual:         lessEqual,
				Keys.Greater:           greater,
				Keys.Less:              less,
				Keys.NotEqual1:         notEqual,
				Keys.NotEqual2:         notEqual,
				Keys.NotEqual3:         notEqual,
				Keys.AllEqual:          allEqual,
				Keys.AllGreaterEqual:   allGreaterEqual,
				Keys.AllLessEqual:      allLessEqual,
				Keys.AllGreater:        allGreater,
				Keys.AllLess:           allLess,
				Keys.AllNotEqual1:      allNotEqual,
				Keys.AllNotEqual2:      allNotEqual,
				Keys.AllNotEqual3:      allNotEqual,
				Keys.AnyEqual:          anyEqual,
				Keys.AnyGreaterEqual:   anyGreaterEqual,
				Keys.AnyLessEqual:      anyLessEqual,
				Keys.AnyGreater:        anyGreater,
				Keys.AnyLess:           anyLess,
				Keys.AnyNotEqual1:      anyNotEqual,
				Keys.AnyNotEqual2:      anyNotEqual,
				Keys.AnyNotEqual3:      anyNotEqual,
				Keys.SomeEqual:         anyEqual,
				Keys.SomeGreaterEqual:  anyGreaterEqual,
				Keys.SomeLessEqual:     anyLessEqual,
				Keys.SomeGreater:       anyGreater,
				Keys.SomeLess:          anyLess,
				Keys.SomeNotEqual1:     anyNotEqual,
				Keys.SomeNotEqual2:     anyNotEqual,
				Keys.SomeNotEqual3:     anyNotEqual,
				Keys.Int:               intFunc,
				Keys.Int8:              int8Func,
				Keys.Int16:             int16Func,
				Keys.Int32:             int32Func,
				Keys.Int64:             int64Func,
				Keys.Float32:           float32Func,
				Keys.Float64:           float64Func,
				Keys.NotBetween:        notBetween,
				Keys.Between:           between,
				Keys.NotRegexp:         notRegexp,
				Keys.Regexp:            regexpFunc,
				Keys.Like:              like,
				Keys.NotLike:           notLike,
				Keys.Sum:               sum,
				Keys.Count:             count,
				Keys.Gcount:            gcount,
				Keys.Gsum:              gsum,
				Keys.Len:               lenFunc,
				Keys.Deflate:           deflateFunc,
				Keys.DateToStr:         dateToStr,
				Keys.TimeToStr:         timeToStr,
				Keys.String:            stringFunc,
			},
		}
		for key, externalMethod := range externalMethods {
			funcProxy.SqlFunc[key] = externalMethod
		}
	})
	return funcProxy
}

func AddFunc(key string, f SqlClosureFunc) {
	externalMethods[key] = f
}

func (self *FuncProxy) Get(name string) SqlFunc {
	return self.SqlFunc[name]()
}
