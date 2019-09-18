package functions

import (
	"github.com/sirupsen/logrus"
	"time"
)

func deflateFunc() SqlFunc {
	return func(x ...interface{}) interface{} {
		switch x0 := x[0].(type) {
		case GroupRecord:
			return deflate(x0, x[1], x[2].(GroupRecord))
		default:
			return deflate(x[0].([]interface{}), x[1], x[2].([]interface{}))
		}
	}
}

func deflate(itemTypes []interface{}, valType interface{}, itemValues []interface{}) interface{} {
	// 多个记录中 itemTypes 字段的值等于 valType 的记录的值求和
	var res interface{} = 0
	for i, subType := range itemTypes {
		if Equal(subType, valType) {
			res = Add(res, itemValues[i])
		}
	}
	return res
}

const dateLayout = "2006-01-02"
const timeLayout = "2006-01-02 15:04:05"

func dateToStr() SqlFunc {
	return func(x ...interface{}) interface{} {
		switch v := x[0].(type) {
		case time.Time:
			return v.Format(dateLayout)
		case string:
			return []rune(v)[:10]
		default:
			logrus.Errorf("[dateToStr] type error: %T", v)
			panic("type error")
		}
	}
}

func timeToStr() SqlFunc {
	return func(x ...interface{}) interface{} {
		switch v := x[0].(type) {
		case time.Time:
			return v.Format(timeLayout)
		case string:
			return []rune(v)[:19]
		default:
			logrus.Errorf("[timeToStr] type error: %T", v)
			panic("type error")
		}
	}
}
