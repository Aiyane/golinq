package functions

func not() SqlFunc {
	return func(x ...interface{}) interface{} {
		return x[0] == nil || x[0] == false || len(x[0].([]interface{})) == 0
	}
}

func signalNot() SqlFunc {
	return func(x ...interface{}) interface{} {
		return x[0] == nil || x[0] == false || len(x[0].([]interface{})) == 0
	}
}

func signalExclamation() SqlFunc {
	return func(x ...interface{}) interface{} {
		return x[0] == nil || x[0] == false || len(x[0].([]interface{})) == 0
	}
}

func exist() SqlFunc {
	return func(x ...interface{}) interface{} {
		return x != nil
	}
}

func is() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Equal(x[0], x[1])
	}
}

func isNot() SqlFunc {
	return func(x ...interface{}) interface{} {
		return !Equal(x[0], x[1])
	}
}

func in() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _x := range x[1].([]interface{}) {
			if Equal(_x, x[0]) {
				return true
			}
		}
		return false
	}
}

func notIn() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _x := range x[1].([]interface{}) {
			if Equal(_x, x[0]) {
				return false
			}
		}
		return true
	}
}

func or() SqlFunc {
	return func(x ...interface{}) interface{} {
		return x[0].(bool) || x[1].(bool)
	}
}

func and() SqlFunc {
	return func(x ...interface{}) interface{} {
		return x[0].(bool) && x[1].(bool)
	}
}

func equal() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Equal(x[0], x[1])
	}
}

func greaterEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		return GreaterEqual(x[0], x[1])
	}
}

func lessEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		return LessEqual(x[0], x[1])
	}
}

func greater() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Greater(x[0], x[1])
	}
}

func less() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Less(x[0], x[1])
	}
}

func notEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		return !Equal(x[0], x[1])
	}
}

func allEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if !Equal(x[0], _y) {
				return false
			}
		}
		return true
	}
}

func allGreaterEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if Less(x[0], _y) {
				return false
			}
		}
		return true
	}
}

func allLessEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if Greater(x[0], _y) {
				return false
			}
		}
		return true
	}
}

func allGreater() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if LessEqual(x[0], _y) {
				return false
			}
		}
		return true
	}
}

func allLess() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if GreaterEqual(x[0], _y) {
				return false
			}
		}
		return true
	}
}

func allNotEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if !Equal(x[0], _y) {
				return false
			}
		}
		return true
	}
}

func anyEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if Equal(x[0], _y) {
				return true
			}
		}
		return false
	}
}

func anyGreaterEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if GreaterEqual(x[0], _y) {
				return true
			}
		}
		return false
	}
}

func anyLessEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if LessEqual(x[0], _y) {
				return true
			}
		}
		return false
	}
}

func anyGreater() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if Greater(x[0], _y) {
				return true
			}
		}
		return false
	}
}

func anyLess() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if Less(x[0], _y) {
				return true
			}
		}
		return false
	}
}

func anyNotEqual() SqlFunc {
	return func(x ...interface{}) interface{} {
		for _, _y := range x[1].([]interface{}) {
			if !Equal(x[0], _y) {
				return true
			}
		}
		return false
	}
}

func notBetween() SqlFunc {
	return func(x ...interface{}) interface{} {
		return Less(x[0], x[1]) || Greater(x[0], x[2])
	}
}

func between() SqlFunc {
	return func(x ...interface{}) interface{} {
		return GreaterEqual(x[0], x[1]) || LessEqual(x[0], x[2])
	}
}
