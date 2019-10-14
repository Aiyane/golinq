package interpreters

import (
	"github.com/Aiyane/golinq/functions"
	"github.com/Aiyane/golinq/types"
	"strings"
)

// 判断字典是否为空
func isEmptyMap(map_ types.DataSources) bool {
	for range map_ {
		return false
	}
	return true
}

/*
值比较函数, 没有泛型！！！！！
item1 > item2 返回 1
item2 < item2 返回 -1
item1 == item2 返回 0
*/
func compareValue(item interface{}, item2 interface{}) int {
	switch item1 := item.(type) {
	case string:
		return strings.Compare(item1, item2.(string))
	case []interface{}:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]interface{})[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []string:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]string)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []int:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]int)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []int8:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]int8)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []int16:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]int16)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []int32:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]int32)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []int64:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]int64)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []float32:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]float32)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []float64:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]float64)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []uint8:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]uint8)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []uint16:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]uint16)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []uint32:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]uint32)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []uint64:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]uint64)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	case []uint:
		for i, v1 := range item1 {
			temRes := compareValue(v1, item2.([]uint)[i])
			if temRes == 1 {
				return 1
			} else if temRes == -1 {
				return -1
			}
		}
		return 0
	default:
		if functions.Greater(item1, item2) {
			return 1
		} else if functions.Less(item1, item2) {
			return -1
		}
		return 0
	}
}

/*
弹出 indexDict 中的一个元素
例如:
indexDict = {
"b": {
  1: [
	{"b": {"a_id": 1, "other": "ha"}},
	{"b": {"a_id": 1, "other": "hh"}},
  ],
 },
"c": {
   2: [
 	{"c": {"a_id": 2, "field": "yo"}},
		{"c": {"a_id": 2, "field": "ya"}},
   ]
  },
}
返回:
"b": {
 1: [
   {"b": {"a_id": 1, "other": "ha"}},
   {"b": {"a_id": 1, "other": "hh"}},
 ],
},
并且
indexDict = {
"c": {
   2: [
 	{"c": {"a_id": 2, "field": "yo"}},
		{"c": {"a_id": 2, "field": "ya"}},
   ]
  },
}
*/
func popItem(indexDict types.IndexDict) (string, types.IndexRecords) {
	for key, value := range indexDict {
		delete(indexDict, key)
		return key, value
	}
	return "", nil
}

// 深拷贝一条记录
// 例如:
// instanceDict = {
// 	"a": {"a_id": 1, "name": "hi"},
// 	"b": {"b_id": 1, "other": "ha"}
// }
// tableName = "c"
// frontInstance = {
// 	"c_id": 1, "recordValue": "yo"
// }
// 返回 {
//  "a": {"a_id": 1, "name": "hi"},
// 	"b": {"b_id": 1, "other": "ha"},
// 	"c": {"c_id": 1, "recordValue": "yo"}
// }
func deepCopyInstanceDict(instanceDict types.InstanceDict, tableName string, record types.Record) types.InstanceDict {
	copyInstanceDict := make(types.InstanceDict, 10)
	for key, record := range instanceDict {
		copyInstanceDict[key] = record
	}
	if tableName != "" {
		copyInstanceDict[tableName] = record
	}
	return copyInstanceDict
}

func useStruct() bool {
	return types.RecordStructName != ""
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
