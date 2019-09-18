package functions

import (
	"regexp"
	"strings"
)

func notRegexp() SqlFunc {
	return func(x ...interface{}) interface{} {
		ok, _ := regexp.MatchString(x[0].(string), x[1].(string))
		return ok == false
	}
}

func regexpFunc() SqlFunc {
	return func(x ...interface{}) interface{} {
		ok, _ := regexp.MatchString(x[0].(string), x[1].(string))
		return ok != false
	}
}

func like() SqlFunc {
	return func(x ...interface{}) interface{} {
		replace := strings.Replace
		ok, _ := regexp.MatchString(
			replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(
				x[1].(string), "\\", "\\\\", -1),
				".", "\\.", -1),
				"^", "\\^", -1),
				"$", "\\$", -1),
				"*", "\\*", -1),
				"+", "\\+", -1),
				"?", "\\?", -1),
				"{", "\\{", -1),
				"}", "\\}", -1),
				"|", "\\|", -1),
				"(", "\\(", -1),
				")", "\\)", -1),
				"%", ".*", -1), x[0].(string))
		return ok != false
	}
}

func notLike() SqlFunc {
	return func(x ...interface{}) interface{} {
		replace := strings.Replace
		ok, _ := regexp.MatchString(
			replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(replace(
				x[1].(string), "\\", "\\\\", -1),
				".", "\\.", -1),
				"^", "\\^", -1),
				"$", "\\$", -1),
				"*", "\\*", -1),
				"+", "\\+", -1),
				"?", "\\?", -1),
				"{", "\\{", -1),
				"}", "\\}", -1),
				"|", "\\|", -1),
				"(", "\\(", -1),
				")", "\\)", -1),
				"%", ".*", -1), x[0].(string))
		return ok == false
	}
}

func first() SqlFunc {
	return func(x ...interface{}) interface{} {
		return string([]rune(x[0].(string))[0])
	}
}
