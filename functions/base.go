package functions

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type NumType int

const (
	NoNum   NumType = -1
	Uint    NumType = 0
	Uint8   NumType = 1
	Uint16  NumType = 2
	Uint32  NumType = 3
	Uint64  NumType = 4
	Int     NumType = 5
	Int8    NumType = 6
	Int16   NumType = 7
	Int32   NumType = 8
	Int64   NumType = 9
	Float32 NumType = 10
	Float64 NumType = 11
	String  NumType = 12
	Time    NumType = 13
)

var (
	Type2Func = map[NumType]func(interface{}) interface{}{
		Uint:    toUint,
		Uint8:   toUint8,
		Uint16:  toUint16,
		Uint32:  toUint32,
		Uint64:  toUint64,
		Int:     ToInt,
		Int8:    toInt8,
		Int16:   toInt16,
		Int32:   toInt32,
		Int64:   toInt64,
		Float32: toFloat32,
		Float64: toFloat64,
		String:  toString,
		Time:    toTime,
		NoNum:   toSelf,
	}
	Name2Type = map[string]NumType{
		"uint":       Uint,
		"uint8":      Uint8,
		"uint16":     Uint16,
		"uint32":     Uint32,
		"uint64":     Uint64,
		"int":        Int,
		"int8":       Int8,
		"int16":      Int16,
		"int32":      Int32,
		"int64":      Int64,
		"float32":    Float32,
		"float64":    Float64,
		"string":     String,
		"time.Time":  Time,
	}
)

func Add(i interface{}, i2 interface{}) interface{} {
	t := GNumType(i, i2)
	switch t {
	case Uint:
		return Type2Func[t](i).(uint) + Type2Func[t](i2).(uint)
	case Uint8:
		return Type2Func[t](i).(uint8) + Type2Func[t](i2).(uint8)
	case Uint16:
		return Type2Func[t](i).(uint16) + Type2Func[t](i2).(uint16)
	case Uint32:
		return Type2Func[t](i).(uint32) + Type2Func[t](i2).(uint32)
	case Uint64:
		return Type2Func[t](i).(uint64) + Type2Func[t](i2).(uint64)
	case Int:
		return Type2Func[t](i).(int) + Type2Func[t](i2).(int)
	case Int8:
		return Type2Func[t](i).(int8) + Type2Func[t](i2).(int8)
	case Int16:
		return Type2Func[t](i).(int16) + Type2Func[t](i2).(int16)
	case Int32:
		return Type2Func[t](i).(int32) + Type2Func[t](i2).(int32)
	case Int64:
		return Type2Func[t](i).(int64) + Type2Func[t](i2).(int64)
	case Float32:
		return Type2Func[t](i).(float32) + Type2Func[t](i2).(float32)
	case Float64:
		return Type2Func[t](i).(float64) + Type2Func[t](i2).(float64)
	default:
		logrus.Errorf("[GreaterEqual] type error: v1[%T], v2[%T]", i, i2)
		panic("type error")
	}
}

func Mul(i interface{}, i2 interface{}) interface{} {
	t := GNumType(i, i2)
	switch t {
	case Uint:
		return Type2Func[t](i).(uint) * Type2Func[t](i2).(uint)
	case Uint8:
		return Type2Func[t](i).(uint8) * Type2Func[t](i2).(uint8)
	case Uint16:
		return Type2Func[t](i).(uint16) * Type2Func[t](i2).(uint16)
	case Uint32:
		return Type2Func[t](i).(uint32) * Type2Func[t](i2).(uint32)
	case Uint64:
		return Type2Func[t](i).(uint64) * Type2Func[t](i2).(uint64)
	case Int:
		return Type2Func[t](i).(int) * Type2Func[t](i2).(int)
	case Int8:
		return Type2Func[t](i).(int8) * Type2Func[t](i2).(int8)
	case Int16:
		return Type2Func[t](i).(int16) * Type2Func[t](i2).(int16)
	case Int32:
		return Type2Func[t](i).(int32) * Type2Func[t](i2).(int32)
	case Int64:
		return Type2Func[t](i).(int64) * Type2Func[t](i2).(int64)
	case Float32:
		return Type2Func[t](i).(float32) * Type2Func[t](i2).(float32)
	case Float64:
		return Type2Func[t](i).(float64) * Type2Func[t](i2).(float64)
	default:
		logrus.Errorf("[GreaterEqual] type error: v1[%T], v2[%T]", i, i2)
		panic("type error")
	}
}

func Res(i interface{}, i2 interface{}) interface{} {
	t := GNumType(i, i2)
	switch t {
	case Uint:
		return Type2Func[t](i).(uint) % Type2Func[t](i2).(uint)
	case Uint8:
		return Type2Func[t](i).(uint8) % Type2Func[t](i2).(uint8)
	case Uint16:
		return Type2Func[t](i).(uint16) % Type2Func[t](i2).(uint16)
	case Uint32:
		return Type2Func[t](i).(uint32) % Type2Func[t](i2).(uint32)
	case Uint64:
		return Type2Func[t](i).(uint64) % Type2Func[t](i2).(uint64)
	case Int:
		return Type2Func[t](i).(int) % Type2Func[t](i2).(int)
	case Int8:
		return Type2Func[t](i).(int8) % Type2Func[t](i2).(int8)
	case Int16:
		return Type2Func[t](i).(int16) % Type2Func[t](i2).(int16)
	case Int32:
		return Type2Func[t](i).(int32) % Type2Func[t](i2).(int32)
	case Int64:
		return Type2Func[t](i).(int64) % Type2Func[t](i2).(int64)
	default:
		logrus.Errorf("[Res] type error: v1[%T], v2[%T]", i, i2)
		panic("type error")
	}
}

func Div(i interface{}, i2 interface{}) interface{} {
	return Mul(i, Countdown(i2))
}

func Sub(i interface{}, i2 interface{}) interface{} {
	return Add(i, Negate(i2))
}

func Countdown(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return 1 / v
	case float64:
		return 1 / v
	case int:
		return 1 / float32(v)
	case int8:
		return 1 / float32(v)
	case int16:
		return 1 / float32(v)
	case int32:
		return 1 / float32(v)
	case int64:
		return 1 / float32(v)
	case uint:
		return 1 / float32(v)
	case uint8:
		return 1 / float32(v)
	case uint16:
		return 1 / float32(v)
	case uint32:
		return 1 / float32(v)
	case uint64:
		return 1 / float32(v)
	default:
		logrus.Errorf("[Negate] type error: v1[%T]", v)
		panic("type error")
	}
}

func Negate(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return v * -1
	case float64:
		return v * -1
	case int:
		return v * -1
	case int8:
		return v * -1
	case int16:
		return v * -1
	case int32:
		return v * -1
	case int64:
		return v * -1
	case uint:
		return int(v) * -1
	case uint8:
		return int8(v) * -1
	case uint16:
		return int16(v) * -1
	case uint32:
		return int32(v) * -1
	case uint64:
		return int64(v) * -1
	default:
		logrus.Errorf("[Negate] type error: v1[%T]", v)
		panic("type error")
	}
}

func Equal(i interface{}, i2 interface{}) bool {
	t := GNumType(i, i2)
	if t == Time {
		return Type2Func[t](i).(time.Time).Equal(Type2Func[t](i2).(time.Time))
	}
	return Type2Func[t](i) == Type2Func[t](i2)
}

func GreaterEqual(i interface{}, i2 interface{}) bool {
	return Greater(i, i2) || Equal(i, i2)
}

func Greater(i interface{}, i2 interface{}) bool {
	t := GNumType(i, i2)
	switch t {
	case Uint:
		return Type2Func[t](i).(uint) > Type2Func[t](i2).(uint)
	case Uint8:
		return Type2Func[t](i).(uint8) > Type2Func[t](i2).(uint8)
	case Uint16:
		return Type2Func[t](i).(uint16) > Type2Func[t](i2).(uint16)
	case Uint32:
		return Type2Func[t](i).(uint32) > Type2Func[t](i2).(uint32)
	case Uint64:
		return Type2Func[t](i).(uint64) > Type2Func[t](i2).(uint64)
	case Int:
		return Type2Func[t](i).(int) > Type2Func[t](i2).(int)
	case Int8:
		return Type2Func[t](i).(int8) > Type2Func[t](i2).(int8)
	case Int16:
		return Type2Func[t](i).(int16) > Type2Func[t](i2).(int16)
	case Int32:
		return Type2Func[t](i).(int32) > Type2Func[t](i2).(int32)
	case Int64:
		return Type2Func[t](i).(int64) > Type2Func[t](i2).(int64)
	case Float32:
		return Type2Func[t](i).(float32) > Type2Func[t](i2).(float32)
	case Float64:
		return Type2Func[t](i).(float64) > Type2Func[t](i2).(float64)
	case Time, String:
		t1 := toTime(i).(time.Time)
		t2 := toTime(i2).(time.Time)
		return t1.After(t2)
	default:
		logrus.Errorf("[GreaterEqual] type error: v1[%T], v2[%T]", i, i2)
		panic("type error")
	}
}

func LessEqual(i interface{}, i2 interface{}) bool {
	return !Greater(i, i2)
}

func Less(i interface{}, i2 interface{}) bool {
	return !Greater(i, i2) && !Equal(i, i2)
}

func GNumType(i interface{}, i2 interface{}) NumType {
	t1 := numType(i)
	t2 := numType(i2)
	if t1 > t2 {
		return t1
	}
	return t2
}

func numType(i interface{}) NumType {
	switch i.(type) {
	case float32:
		return Float32
	case float64:
		return Float64
	case int:
		return Int
	case int8:
		return Int8
	case int16:
		return Int16
	case int32:
		return Int32
	case int64:
		return Int64
	case uint:
		return Uint
	case uint8:
		return Uint8
	case uint16:
		return Uint16
	case uint32:
		return Uint32
	case uint64:
		return Uint64
	case string:
		return String
	case time.Time:
		return Time
	default:
		return NoNum
	}
}

func toSelf(i interface{}) interface{} {
	return i
}

func toString(i interface{}) interface{} {
	return fmt.Sprintf("%v", i)
}

func ToInt(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return int(v)
	case float64:
		return int(v)
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[ToInt] type error: v1[%T]", v)
			panic(err)
		}
		return ret
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case nil:
		return 0
	default:
		logrus.Errorf("[ToInt] type error: v1[%T]", v)
		panic("type err")
	}
}

func toInt8(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return int8(v)
	case float64:
		return int8(v)
	case int:
		return int8(v)
	case int8:
		return v
	case int16:
		return int8(v)
	case int32:
		return int8(v)
	case int64:
		return int8(v)
	case uint:
		return int8(v)
	case uint8:
		return int8(v)
	case uint16:
		return int8(v)
	case uint32:
		return int8(v)
	case uint64:
		return int8(v)
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[toInt8] type error: v1[%T]", v)
			panic(err)
		}
		return int8(ret)
	case bool:
		if v {
			return int8(1)
		} else {
			return int8(0)
		}
	case nil:
		return int8(0)
	default:
		logrus.Errorf("[toInt8] type error: v1[%T]", v)
		panic("type err")
	}
}

func toInt16(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return int16(v)
	case float64:
		return int16(v)
	case int:
		return int16(v)
	case int8:
		return int16(v)
	case int16:
		return v
	case int32:
		return int16(v)
	case int64:
		return int16(v)
	case uint:
		return int16(v)
	case uint8:
		return int16(v)
	case uint16:
		return int16(v)
	case uint32:
		return int16(v)
	case uint64:
		return int16(v)
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[toInt16] type error: v1[%T]", v)
			panic(err)
		}
		return int16(ret)
	case bool:
		if v {
			return int16(1)
		} else {
			return int16(0)
		}
	case nil:
		return int16(0)
	default:
		logrus.Errorf("[toInt16] type error: v1[%T]", v)
		panic("type err")
	}
}

func toInt32(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return int32(v)
	case float64:
		return int32(v)
	case int:
		return int32(v)
	case int8:
		return int32(v)
	case int16:
		return int32(v)
	case int32:
		return v
	case int64:
		return int32(v)
	case uint:
		return int32(v)
	case uint8:
		return int32(v)
	case uint16:
		return int32(v)
	case uint32:
		return int32(v)
	case uint64:
		return int32(v)
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[toInt32] type error: v1[%T]", v)
			panic(err)
		}
		return int32(ret)
	case bool:
		if v {
			return int32(1)
		} else {
			return int32(0)
		}
	case nil:
		return int32(0)
	default:
		logrus.Errorf("[toInt32] type error: v1[%T]", v)
		panic("type err")
	}
}

func toInt64(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[toInt64] type error: v1[%T]", v)
			panic(err)
		}
		return int64(ret)
	case bool:
		if v {
			return int64(1)
		} else {
			return int64(0)
		}
	case nil:
		return int64(0)
	default:
		logrus.Errorf("[toInt64] type error: v1[%T]", v)
		panic("type err")
	}
}

func toFloat32(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return v
	case float64:
		return float32(v)
	case int:
		return float32(v)
	case int8:
		return float32(v)
	case int16:
		return float32(v)
	case int32:
		return float32(v)
	case int64:
		return float32(v)
	case uint:
		return float32(v)
	case uint8:
		return float32(v)
	case uint16:
		return float32(v)
	case uint32:
		return float32(v)
	case uint64:
		return float32(v)
	case string:
		ret, err := strconv.ParseFloat(v, 32)
		if err != nil {
			logrus.Errorf("[toFloat32] type error: v1[%T]", v)
			panic(err)
		}
		return ret
	case bool:
		if v {
			return float32(1)
		} else {
			return float32(0)
		}
	case nil:
		return float32(0)
	default:
		logrus.Errorf("[toFloat32] type error: v1[%T]", v)
		panic("type err")
	}
}

func toFloat64(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case string:
		ret, err := strconv.ParseFloat(v, 64)
		if err != nil {
			logrus.Errorf("[toFloat64] type error: v1[%T]", v)
			panic(err)
		}
		return ret
	case bool:
		if v {
			return float64(1)
		} else {
			return float64(0)
		}
	case nil:
		return float64(0)
	default:
		logrus.Errorf("[toFloat64] type error: v1[%T]", v)
		panic("type err")
	}
}

func toUint(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return uint(v)
	case float64:
		return uint(v)
	case int:
		return uint(v)
	case int8:
		return uint(v)
	case int16:
		return uint(v)
	case int32:
		return uint(v)
	case int64:
		return uint(v)
	case uint:
		return v
	case uint8:
		return uint(v)
	case uint16:
		return uint(v)
	case uint32:
		return uint(v)
	case uint64:
		return uint(v)
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[toUint] type error: v1[%T]", v)
			panic(err)
		}
		return uint(ret)
	case bool:
		if v {
			return uint(1)
		} else {
			return uint(0)
		}
	case nil:
		return uint(0)
	default:
		logrus.Errorf("[toUint] type error: v1[%T]", v)
		panic("type err")
	}
}

func toUint8(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return uint8(v)
	case float64:
		return uint8(v)
	case int:
		return uint8(v)
	case int8:
		return uint8(v)
	case int16:
		return uint8(v)
	case int32:
		return uint8(v)
	case int64:
		return uint8(v)
	case uint:
		return uint8(v)
	case uint8:
		return v
	case uint16:
		return uint8(v)
	case uint32:
		return uint8(v)
	case uint64:
		return uint8(v)
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[toUint8] type error: v1[%T]", v)
			panic(err)
		}
		return uint8(ret)
	case bool:
		if v {
			return uint8(1)
		} else {
			return uint8(0)
		}
	case nil:
		return uint8(0)
	default:
		logrus.Errorf("[toUint8] type error: v1[%T]", v)
		panic("type err")
	}
}

func toUint16(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return uint16(v)
	case float64:
		return uint16(v)
	case int:
		return uint16(v)
	case int8:
		return uint16(v)
	case int16:
		return uint16(v)
	case int32:
		return uint16(v)
	case int64:
		return uint16(v)
	case uint:
		return uint16(v)
	case uint8:
		return uint16(v)
	case uint16:
		return v
	case uint32:
		return uint16(v)
	case uint64:
		return uint16(v)
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[toUint16] type error: v1[%T]", v)
			panic(err)
		}
		return uint16(ret)
	case bool:
		if v {
			return uint16(1)
		} else {
			return uint16(0)
		}
	case nil:
		return uint16(0)
	default:
		logrus.Errorf("[toUint16] type error: v1[%T]", v)
		panic("type err")
	}
}

func toUint32(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return uint32(v)
	case float64:
		return uint32(v)
	case int:
		return uint32(v)
	case int8:
		return uint32(v)
	case int16:
		return uint32(v)
	case int32:
		return uint32(v)
	case int64:
		return uint32(v)
	case uint:
		return uint32(v)
	case uint8:
		return uint32(v)
	case uint16:
		return uint32(v)
	case uint32:
		return v
	case uint64:
		return uint32(v)
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[toUint32] type error: v1[%T]", v)
			panic(err)
		}
		return uint32(ret)
	case bool:
		if v {
			return uint32(1)
		} else {
			return uint32(0)
		}
	case nil:
		return uint32(0)
	default:
		logrus.Errorf("[toUint32] type error: v1[%T]", v)
		panic("type err")
	}
}

func toUint64(i interface{}) interface{} {
	switch v := i.(type) {
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case string:
		ret, err := strconv.Atoi(v)
		if err != nil {
			logrus.Errorf("[toUint64] type error: v1[%T]", v)
			panic(err)
		}
		return uint64(ret)
	case bool:
		if v {
			return uint64(1)
		} else {
			return uint64(0)
		}
	case nil:
		return uint64(0)
	default:
		logrus.Errorf("[toUint64] type error: v1[%T]", v)
		panic("type err")
	}
}

func toTime(stringTime interface{}) interface{} {
	switch v := stringTime.(type) {
	case time.Time:
		return v
	case string:
		l := len([]rune(v))
		if l == 10 {
			theTime, err := time.ParseInLocation(dateLayout, v, time.Local)
			if err != nil {
				logrus.Errorf("[toTime] has error the stringTime len: %v", l)
				panic(err)
			}
			return theTime
		}
		if l >= 19 {
			theTime, err := time.ParseInLocation(timeLayout, v[:19], time.Local)
			if err != nil {
				logrus.Errorf("[toTime] has error the stringTime len: %v", l)
				panic(err)
			}
			return theTime
		}
	case nil:
		return time.Time{}
	}
	logrus.Errorf("[toTime] has error the stringTime: %T", stringTime)
	panic("type error")
}
