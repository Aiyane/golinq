package test

import (
	"strings"
)

func isEqual(res1 []interface{}, res2 []interface{}) bool {
	if len(res1) != len(res2) {
		return false
	}
	for index, dict := range res1 {
		if len(dict.(map[string]interface{})) != len(res2[index].(map[string]interface{})) {
			return false
		} else {
			for key, value := range dict.(map[string]interface{}) {
				switch expr := value.(type) {
				case []interface{}:
					if lstEqual(expr, res2[index].(map[string]interface{})[key].([]interface{})) == false {
						return false
					}
				case []string:
					if lstStringEqual(expr, res2[index].(map[string]interface{})[key].([]string)) == false {
						return false
					}
				case map[string]interface{}:
					if dictEqual(expr, res2[index].(map[string]interface{})[key].(map[string]interface{})) == false {
						return false
					}
				default:
					if value != res2[index].(map[string]interface{})[key] {
						return false
					}
				}

			}
		}
	}
	return true
}

func dictEqual(items1 map[string]interface{}, items2 map[string]interface{}) bool {
	for key, value := range items1 {
		switch expr := value.(type) {
		case []interface{}:
			if lstEqual(expr, items2[key].([]interface{})) == false {
				return false
			}
		case map[string]interface{}:
			if dictEqual(expr, items2[key].(map[string]interface{})) == false {
				return false
			}
		default:
			if value != items2[key] {
				return false
			}
		}
	}
	return true
}

func lstEqual(items1 []interface{}, items2 []interface{}) bool {
	for i, v1 := range items1 {
		switch expr := v1.(type) {
		case []interface{}:
			if lstEqual(expr, items2[i].([]interface{})) == false {
				return false
			}
		case map[string]interface{}:
			if dictEqual(expr, items2[i].(map[string]interface{})) == false {
				return false
			}
		default:
			if v1 != items2[i] {
				return false
			}
		}
	}
	return true
}

func lstStringEqual(items1 []string, items2 []string) bool {
	for i, v1 := range items1 {
		if v1 != items2[i] {
			return false
		}
	}
	return true
}

func quickSort(arr *[]interface{}, start, end int, keys []string) {
	if start < end {
		i, j := start, end
		key := (*arr)[(start+end)/2]
		for i <= j {
			for compareRecord((*arr)[i].(map[string]interface{}), key.(map[string]interface{}), keys) < 0 {
				i++
			}
			for compareRecord((*arr)[j].(map[string]interface{}), key.(map[string]interface{}), keys) > 0 {
				j--
			}
			if i <= j {
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j, keys)
		}
		if end > i {
			quickSort(arr, i, end, keys)
		}
	}
}

func compare(item interface{}, item2 interface{}) int {
	switch item1 := item.(type) {
	case string:
		return strings.Compare(item1, item2.(string))
	case int:
		if item1 > item2.(int) {
			return 1
		} else if item1 < item2.(int) {
			return -1
		}
		return 0
	case float64:
		if item1 > item2.(float64) {
			return 1
		} else if item1 < item2.(float64) {
			return -1
		}
		return 0
	default:
		panic("类型错误")
	}
}

func compareRecord(record map[string]interface{}, record2 map[string]interface{}, keys []string) int {
	for _, key := range keys {
		res := compare(record[key], record2[key])
		if res > 0 {
			return 1
		} else if res < 0 {
			return -1
		}
	}
	return 0
}

func sorted(res *[]interface{}, keys []string) {
	end := len(*res) - 1
	quickSort(res, 0, end, keys)
}
