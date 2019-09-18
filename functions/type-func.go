package functions

import (
	"fmt"
)

func intFunc() SqlFunc {
	return func(x ...interface{}) interface{} {
		return ToInt(x[0])
	}
}

func int8Func() SqlFunc {
	return func(x ...interface{}) interface{} {
		return toInt8(x[0])
	}
}

func int16Func() SqlFunc {
	return func(x ...interface{}) interface{} {
		return toInt16(x[0])
	}
}

func int32Func() SqlFunc {
	return func(x ...interface{}) interface{} {
		return toInt32(x[0])
	}
}

func int64Func() SqlFunc {
	return func(x ...interface{}) interface{} {
		return toInt64(x[0])
	}
}

func float32Func() SqlFunc {
	return func(x ...interface{}) interface{} {
		return toFloat32(x[0])
	}
}

func float64Func() SqlFunc {
	return func(x ...interface{}) interface{} {
		return toFloat64(x[0])
	}
}

func stringFunc() SqlFunc {
	return func(x ...interface{}) interface{} {
		return fmt.Sprintf("%v", x[0])
	}
}
