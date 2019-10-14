package interpreters

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/Aiyane/golinq/functions"
	"github.com/Aiyane/golinq/types"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"time"
)

var TagString string

// 赋值 设置 struct/map 的值
// isDelete 删除 key 的值
func setValue(resRecord interface{}, key string, value interface{}, isDelete bool) {
	if isStruct(resRecord) {
		setStructValue(resRecord, key, value, isDelete)
	} else {
		setMapVlaue(resRecord, key, value, isDelete)
	}
}

// 赋值 为 map 赋值
func setMapVlaue(resRecord interface{}, key string, value interface{}, isDelete bool) {
	if isDelete {
		delete(resRecord.(map[string]interface{}), key)
	} else {
		if val, ok := value.(driver.Valuer); ok {
			s, err := val.Value()
			if err != nil {
				panic(err)
			}
			resRecord.(map[string]interface{})[key] = []byte(fmt.Sprintf("%v", s))
		} else {
			resRecord.(map[string]interface{})[key] = value
		}
	}
}

// 赋值 为 struct 赋值
func setStructValue(resRecord interface{}, key string, value interface{}, isDelete bool) {
	t := reflect.TypeOf(resRecord)
	v := reflect.ValueOf(resRecord)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	var resSetName string
	switch table := resRecord.(type) {
	case types.Entitier:
		resSetName = table.TableName()
	default:
		resSetName = t.Name()
	}
	if resSetName == types.RecordStructName {
		if isDelete {
			delete(hadSetStructFieldValue, key)
			delete(tmpFieldValue, key)
		} else {
			hadSetStructFieldValue[key] = struct{}{}
		}
	}
	camelKey := camelString(key)
	if _, exist := t.FieldByName(camelKey); exist {
		setExistValue(camelKey, value, v, isDelete)
	} else {
		if _, exist := t.FieldByName("ID"); exist && camelKey == "Id" {
			setExistValue("ID", value, v, isDelete)
		} else {
			setValueFromTag(t, v, key, value, isDelete)
			tmpFieldValue[key] = value
		}
	}
}

// 赋值 存在的字段
func setExistValue(camelKey string, value interface{}, v reflect.Value, isDelete bool) {
	// 设置过属性值
	hadSetStructFieldValue[camelKey] = struct{}{}
	field := v.FieldByName(camelKey)
	if isDelete || value == nil {
		field.Set(reflect.Zero(field.Type()))
	} else {
		setFieldValue(field, value)
	}
}

// 赋值 从 tag 名中设置值
func setValueFromTag(resRecordOfType reflect.Type, resRecordOfValue reflect.Value, key string, value interface{}, isDelete bool) {
	fieldNum := resRecordOfType.NumField()
	var tags []string
	var name string
	for i := 0; i < fieldNum; i++ {
		if isNestStruct(resRecordOfValue.Field(i)) {
			structField := resRecordOfValue.Field(i).Type()
			for j := 0; j < structField.NumField(); j++ {
				name = structField.Field(j).Name
				if name == camelString(key) {
					setFieldValue(resRecordOfValue.Field(i).Field(j), value)
					return
				}
				tags = strings.Split(string(structField.Field(j).Tag), "\"")
			}
		} else {
			name = resRecordOfType.Field(i).Name
			tags = strings.Split(string(resRecordOfType.Field(i).Tag), "\"")
		}
		if len(tags) > 1 {
			for j, tag := range tags {
				if (tag == TagString+":" || tag == " "+TagString+":") && len(tags) >= j+2 {
					if tags[j+1] == key {
						field := resRecordOfValue.FieldByName(name)
						// 设置过属性值
						if isDelete || value == nil {
							field.Set(reflect.Zero(field.Type()))
						} else {
							setFieldValue(field, value)
						}
						return
					}
					break
				}
			}
		}
	}
	// 没有的字段不用设置值
}

// 赋值 为结构体字段赋值 field 结构体字段, key 字段名, v 值
// mysql 只有三类数据类型: 数值、字符串、日期
// golang 缺少泛型
func setFieldValue(field reflect.Value, v interface{}) {
	typeName := field.Type().String()
	realTypeName := strings.TrimLeft(typeName, "*")
	t, ok := functions.Name2Type[realTypeName]
	if !ok {
		logrus.Warnf("[setFieldValue] type %v not exist in Name2Type", typeName)
		if val, ok := field.Addr().Interface().(sql.Scanner); ok {
			err := val.Scan([]byte(fmt.Sprintf("%v", v)))
			if err != nil {
				panic(err)
			}
		} else {
			field.Set(reflect.ValueOf(v))
		}
	} else {
		v = functions.Type2Func[t](reflect.ValueOf(v).Interface())
		numOfAddress := len(typeName) - len(realTypeName)
		for i := 0; i < numOfAddress; i++ {
			switch val := v.(type) {
			case float32:
				v = &val
			case float64:
				v = &val
			case int:
				v = &val
			case int8:
				v = &val
			case int16:
				v = &val
			case int32:
				v = &val
			case int64:
				v = &val
			case uint:
				v = &val
			case uint8:
				v = &val
			case uint16:
				v = &val
			case uint32:
				v = &val
			case uint64:
				v = &val
			case string:
				v = &val
			case time.Time:
				v = &val
			}
		}
		field.Set(reflect.ValueOf(v))
	}
}

// 取值
func recordValue(record interface{}, key string) (interface{}, bool) {
	if isStruct(record) {
		return valueFromStruct(record, key)
	}
	return valueFromMap(record, key)
}

// 取值 返回 {key:value}
func recordKeyValue(record interface{}, fn func(name string, value interface{})) {
	if !isStruct(record) {
		for k, v := range record.(map[string]interface{}) {
			fn(k, v)
		}
	} else {
		t := reflect.TypeOf(record)
		v := reflect.ValueOf(record)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
			v = v.Elem()
		}
		var tagName string
		var tags []string
		fieldNum := t.NumField()
		for i := 0; i < fieldNum; i++ {
			name := t.Field(i).Name
			if isNestStruct(v.Field(i)) {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					newName := structField.Field(j).Name
					tagName = snakeString(newName)
					tags = strings.Split(string(structField.Field(j).Tag), "\"")
					fn(getTagName(tagName, tags), reflect.Indirect(reflect.ValueOf(record)).FieldByName(newName).Interface())
				}
			} else {
				tagName = snakeString(t.Field(i).Name)
				tags = strings.Split(string(t.Field(i).Tag), "\"")
				fn(getTagName(tagName, tags), reflect.Indirect(reflect.ValueOf(record)).FieldByName(name).Interface())
			}
		}
	}
}

func getTagName(tagName string, tags []string) string {
	if len(tags) > 1 {
		for j, tag := range tags {
			if (tag == TagString+":" || tag == " "+TagString+":") && len(tags) >= j+2 {
				return tags[j+1]
			}
		}
	}
	return tagName
}

// 判断是否是嵌套结构
func isNestStruct(v reflect.Value) bool {
	name := v.Type().Name()
	if v.Type().Kind() == reflect.Struct && name != "Time" {
		return true
	}
	return false
}

// 取值 从 map 中取值
func valueFromMap(record interface{}, key string) (interface{}, bool) {
	v, ok := record.(map[string]interface{})[key]
	return v, ok
}

// 取值 从 struct 中取值
func valueFromStruct(record interface{}, key string) (interface{}, bool) {
	t := reflect.TypeOf(record)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	var had bool
	if _, ok := hadSetStructFieldValue[key]; ok {
		had = true
	} else {
		had = false
	}
	camelKey := camelString(key)
	if _, exist := t.FieldByName(camelKey); exist {
		return valueFromField(reflect.Indirect(reflect.ValueOf(record)).FieldByName(camelKey)), had
	}
	if _, exist := t.FieldByName("ID"); exist && camelKey == "Id" {
		return valueFromField(reflect.Indirect(reflect.ValueOf(record)).FieldByName("ID")), had
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			for j, tag := range tags {
				if (tag == TagString+":" || tag == " "+TagString+":") && len(tags) >= j+2 {
					if tags[j+1] == key {
						return valueFromField(reflect.Indirect(reflect.ValueOf(record)).FieldByName(name)), had
					}
					break
				}
			}
		}
		if snakeString(name) == key {
			return valueFromField(reflect.Indirect(reflect.ValueOf(record)).FieldByName(name)), had
		}
	}
	if had {
		return tmpFieldValue[key], had
	}
	return nil, had
}

// 取值 通过字段名取值
func valueFromField(field reflect.Value) interface{} {
	value := field.Interface()
	if val, ok := value.(driver.Valuer); ok {
		s, err := val.Value()
		if err != nil {
			panic(err)
		}
		return []byte(fmt.Sprintf("%v", s))
	}
	return value
}

// 实例化结构体
func newStruct(name string) interface{} {
	elem, ok := TypeRegistry[name]
	if !ok {
		return nil
	}
	return reflect.New(elem.Elem()).Interface()
}

// 判断是 struct
func isStruct(record interface{}) bool {
	t := reflect.TypeOf(record)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return false
	}
	return true
}
