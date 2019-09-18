package functions

func signalAdd() SqlFunc {
	return func(x ...interface{}) interface{} {
		return x[0]
	}
}

func signalSub() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Negate(x[0])
	}
}

func doubleSub() SqlFunc {
	return func(x ...interface{}) interface{} {
		return x[0]
	}
}

func add() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Add(x[0], x[1])
	}
}

func sub() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Sub(x[0], x[1])
	}
}

func mul() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Mul(x[0], x[1])
	}
}

func div() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Div(x[0], x[1])
	}
}

func res() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Res(x[0], x[1])
	}
}

func sum() SqlFunc {
	var s interface{} = 0
	return func(x ...interface{}) interface{} {
		switch v := x[0].(type) {
		case []float64:
			for _, item := range v {
				s = Add(s, item)
			}
		case []int:
			for _, item := range v {
				s = Add(s, item)
			}
		case []float32:
			for _, item := range v {
				s = Add(s, item)
			}
		case []int8:
			for _, item := range v {
				s = Add(s, item)
			}
		case []int16:
			for _, item := range v {
				s = Add(s, item)
			}
		case []int32:
			for _, item := range v {
				s = Add(s, item)
			}
		case []int64:
			for _, item := range v {
				s = Add(s, item)
			}
		case []uint:
			for _, item := range v {
				s = Add(s, item)
			}
		case []uint8:
			for _, item := range v {
				s = Add(s, item)
			}
		case []uint16:
			for _, item := range v {
				s = Add(s, item)
			}
		case []uint32:
			for _, item := range v {
				s = Add(s, item)
			}
		case []uint64:
			for _, item := range v {
				s = Add(s, item)
			}
		default:
			s = Add(s, v)
		}
		return s
	}
}

func count() SqlFunc {
	var s interface{} = 0
	return func(x ...interface{}) interface{} {
		switch v := x[0].(type) {
		case []interface{}:
			s = Add(s, len(v))
		default:
			s = Add(s, 1)
		}
		return s
	}
}

func gcount() SqlFunc {
	return func(x ...interface{}) interface{} {
		switch v := x[0].(type) {
		case []interface{}:
			return len(v)
		case GroupRecord:
			return len(v)
		default:
			return 1
		}
	}
}

func gsum() SqlFunc {
	return func(x ...interface{}) interface{} {
		switch v := x[0].(type) {
		case []interface{}:
			var s interface{} = 0
			for _, val := range v {
				s = Add(s, val)
			}
			return s
		case GroupRecord:
			var s interface{} = 0
			for _, val := range v {
				switch subV := val.(type) {
				case int, float32, float64, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
					s = Add(s, subV)
				default:
					s = Add(s, 1)
				}
			}
			return s
		case int, float32, float64, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			return v
		default:
			return 1.0
		}
	}
}

func lenFunc() SqlFunc {
	return func(x ...interface{}) interface{} {
		switch v := x[0].(type) {
		case []interface{}:
			return len(v)
		case GroupRecord:
			return len(v)
		case string:
			return len([]rune(v))
		default:
			return len(x)
		}
	}
}
