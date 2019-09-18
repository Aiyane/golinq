package interpreters

import (
	"github.com/Aiyane/golinq/functions"
	"github.com/Aiyane/golinq/types"
	"reflect"
	"strings"
)

var TagString string

func RegisterType(elem interface{}) {
	t := reflect.TypeOf(elem)
	e := t.Elem()
	TypeRegistry[e.Name()] = t
}

func newStruct(name string) interface{} {
	elem, ok := TypeRegistry[name]
	if !ok {
		return nil
	}
	return reflect.New(elem.Elem()).Interface()
}

// 设置 struct/map 的值
// isDelete 删除 key 的值
func setValue(resRecord interface{}, key string, value interface{}, isDelete bool) {
	if isStruct(resRecord) {
		setStructValue(resRecord, key, value, isDelete)
	} else {
		setMapVlaue(resRecord, key, value, isDelete)
	}
}

// 为 map 赋值
func setMapVlaue(resRecord interface{}, key string, value interface{}, isDelete bool) {
	if isDelete {
		delete(resRecord.(map[string]interface{}), key)
	} else {
		resRecord.(map[string]interface{})[key] = value
	}
}

// 为 struct 赋值
func setStructValue(resRecord interface{}, key string, value interface{}, isDelete bool) {
	t := reflect.TypeOf(resRecord)
	v := reflect.ValueOf(resRecord)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if t.Name() == types.RecordStructName {
		if isDelete {
			delete(hadSetStructFieldValue, key)
			delete(tmpFieldValue, key)
		} else {
			hadSetStructFieldValue[key] = true
		}
	}
	if _, exist := t.FieldByName(key); exist {
		// 设置过属性值
		hadSetStructFieldValue[key] = true
		setFieldValue(v.FieldByName(key), value)
	} else {
		setValueFromTag(t, v, key, value, isDelete)
		tmpFieldValue[key] = value
	}
}

// 从 tag 名中设置值
func setValueFromTag(resRecordOfType reflect.Type, resRecordOfValue reflect.Value, key string, value interface{}, isDelete bool) {
	fieldNum := resRecordOfType.NumField()
	for i := 0; i < fieldNum; i++ {
		name := resRecordOfType.Field(i).Name
		tags := strings.Split(string(resRecordOfType.Field(i).Tag), "\"")
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
}

// 为结构体字段赋值 field 结构体字段, key 字段名, v 值
func setFieldValue(field reflect.Value, v interface{}) {
	typeName := field.Type().String()
	t, ok := functions.Name2Type[typeName]
	if !ok {
		field.Set(reflect.ValueOf(v))
	} else {
		realV := functions.Type2Func[t](reflect.ValueOf(v).Interface())
		field.Set(reflect.ValueOf(realV))
	}
}

// 从 map 中取值
func valueFromMap(record interface{}, key string) (interface{}, bool) {
	v, ok := record.(map[string]interface{})[key]
	return v, ok
}

// 从 struct 中取值
func valueFromStruct(record interface{}, key string) (interface{}, bool) {
	t := reflect.TypeOf(record)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	var had bool
	if hadSetStructFieldValue[key] {
		had = true
	} else {
		had = false
	}
	if _, exist := t.FieldByName(key); exist {
		return reflect.Indirect(reflect.ValueOf(record)).FieldByName(key).Interface(), had
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			for j, tag := range tags {
				if (tag == TagString+":" || tag == " "+TagString+":") && len(tags) >= j+2 {
					if tags[j+1] == key {
						return reflect.Indirect(reflect.ValueOf(record)).FieldByName(name).Interface(), had
					}
					break
				}
			}
		}
	}
	if had {
		return tmpFieldValue[key], had
	}
	return nil, had
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

// 取值
func recordValue(record interface{}, key string) (interface{}, bool) {
	if isStruct(record) {
		return valueFromStruct(record, key)
	}
	return valueFromMap(record, key)
}

// 返回 {key:value}
func recordKeyValue(record interface{}, fn func(name string, value interface{})) {
	if !isStruct(record) {
		for k, v := range record.(map[string]interface{}) {
			fn(k, v)
		}
	} else {
		t := reflect.TypeOf(record)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		fieldNum := t.NumField()
		for i := 0; i < fieldNum; i++ {
			name := t.Field(i).Name
			tagName := t.Field(i).Name
			tags := strings.Split(string(t.Field(i).Tag), "\"")
			if len(tags) > 1 {
				for j, tag := range tags {
					if (tag == TagString+":" || tag == " "+TagString+":") && len(tags) >= j+2 {
						tagName = tags[j+1]
						break
					}
				}
			}
			fn(tagName, reflect.Indirect(reflect.ValueOf(record)).FieldByName(name).Interface())
		}
	}
}
