package golinq

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go/ast"
	"reflect"
	"strings"
	"sync"
	"time"
)

type Company struct {
	Id         int    `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	CreateTime time.Time
}

func (Company) TableName() string {
	return "company"
}

type Mobile struct {
	Id        int    `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	CompanyId int    `json:"company_id"`
}

func (Mobile) TableName() string {
	return "mobile"
}

type MobilePrice struct {
	Id       int     `gorm:"primary_key" json:"id"`
	MobileId int     `json:"mobile_id"`
	Type     int     `json:"type"`
	Price    float32 `json:"price"`
}

func (MobilePrice) TableName() string {
	return "mobile_price"
}

func MessageDBError(message string) error {
	if message == "" {
		return nil
	}
	return &DBError{message: message}
}

type DBError struct {
	error
	message string
}

func MessageDBErrorf(message string, args ...interface{}) error {
	if message == "" {
		return nil
	}
	return &DBError{message: fmt.Sprintf(message, args...)}
}

func reflectValue(values interface{}) (reflect.Value, error) {
	if values == nil {
		return reflect.Value{}, errors.WithStack(MessageDBError("values is nil"))
	}
	refValues := reflect.ValueOf(values)
	if refValues.IsNil() {
		return reflect.Value{}, errors.WithStack(MessageDBError("values is nil"))
	}
	//判断是否为切片类型
	if refValues.Kind() != reflect.Slice {
		return reflect.Value{}, errors.WithStack(MessageDBErrorf("values type `%v`, not `Slice`", refValues.Kind()))
	}
	if refValues.Len() == 0 {
		return reflect.Value{}, errors.WithStack(MessageDBError("values is empty"))
	}
	return refValues, nil
}

func (r *Company) BatchCreate(db *gorm.DB, values interface{}) (err error) {
	refValues, err := reflectValue(values)
	if err != nil {
		return err
	}
	db = db.Table(r.TableName())
	//获取回调函数
	beginTransaction := db.Callback().Create().Get("gorm:begin_transaction")
	beforeCreate := db.Callback().Create().Get("gorm:before_create")
	updateTimeStamp := db.Callback().Create().Get("gorm:update_time_stamp")
	afterCreate := db.Callback().Create().Get("gorm:after_create")
	commitOrRollbackTransaction := db.Callback().Create().Get("gorm:commit_or_rollback_transaction")

	scope := db.NewScope(refValues.Index(0).Interface())
	beginTransaction(scope)
	beforeCreate(scope)
	updateTimeStamp(scope)

	columns := make([]string, 0)
	placeholders := make([]string, 0)
	var blankColumnsWithDefaultValue []string
	for _, field := range scope.Fields() {
		if field.IsNormal {
			if field.IsBlank && field.HasDefaultValue {
				blankColumnsWithDefaultValue = append(blankColumnsWithDefaultValue, scope.Quote(field.DBName))
				scope.InstanceSet("gorm:blank_columns_with_default_value", blankColumnsWithDefaultValue)
			} else if !field.IsPrimaryKey || !field.IsBlank {
				columns = append(columns, scope.Quote(field.DBName))
				if field.Name == "Status" {
					placeholders = append(placeholders, scope.AddToVars(1))
				} else {
					placeholders = append(placeholders, scope.AddToVars(field.Field.Interface()))
				}
			}
		}
	}

	placeholdersList := make([]string, refValues.Len())
	placeholdersStr := fmt.Sprintf("(%s)", strings.Join(placeholders, ","))
	placeholdersList[0] = placeholdersStr

	for i := 1; i < refValues.Len(); i++ {
		elemScope := db.NewScope(refValues.Index(i).Interface())
		beforeCreate(elemScope)
		updateTimeStamp(elemScope)
		for _, field := range elemScope.Fields() {
			if field.IsNormal {
				if field.Name == "Status" {
					scope.AddToVars(1)
				} else if !field.IsPrimaryKey || !field.IsBlank {
					scope.AddToVars(field.Field.Interface())
				}
			}
		}
		placeholdersList[i] = placeholdersStr
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
		scope.QuotedTableName(),
		strings.Join(columns, ","),
		strings.Join(placeholdersList, ","),
	)
	err = scope.Raw(sql).Exec().DB().Error
	if err != nil {
		return errors.WithStack(err)
	}
	afterCreate(scope)
	commitOrRollbackTransaction(scope)
	err = errors.WithStack(scope.DB().Error)
	return err
}

// 字段数组传参为属性名，非表字段名
func (r *Company) BatchUpdate(db *gorm.DB, propertyNames []string, values interface{}) (err error) {
	refValues, err := reflectValue(values)
	if err != nil {
		return err
	}
	indirectValue := reflect.Indirect(refValues)
	if indirectValue.Kind() != reflect.Slice {
		return errors.WithStack(MessageDBError("values is not `slice` type"))
	}
	v := indirectValue.Index(0).Elem()
	fieldName2Index := depFieldName2Index(v, 2)
	ids := make([]interface{}, 0)
	for i := 0; i < indirectValue.Len(); i++ { //批量更新field
		ids = append(ids, reflect.Indirect(indirectValue.Index(i)).FieldByIndex(fieldName2Index["Id"]).Interface())
	}
	sql := "UPDATE " + r.TableName() + " SET "
	args := make([]interface{}, 0)
	for j, fieldName := range propertyNames {
		field := gorm.ToDBName(fieldName)
		var set, when string
		set += field + " = CASE id"
		for i := 0; i < indirectValue.Len(); i++ {
			v := reflect.Indirect(indirectValue.Index(i))
			when += " WHEN ? THEN ?"
			args = append(args, ids[i], v.FieldByIndex(fieldName2Index[fieldName]).Interface())
		}
		sql += set + when + " END"
		if j < len(propertyNames)-1 {
			sql += ","
		}
	}
	sql += " WHERE id in (?);"
	args = append(args, ids)
	err = db.Exec(sql, args...).Error
	return errors.WithStack(err)
}

func depFieldName2Index(value reflect.Value, depth int) map[string][]int {
	k := modelStructKey{t: value.Type(), d: depth}
	if v, ok := modelStructs.Load(k); ok {
		return v.(map[string][]int)
	}
	name2Index := make(map[string][]int)
	fieldName2Index(value, name2Index, []int{}, depth)
	modelStructs.Store(k, name2Index)
	return name2Index
}

type modelStructKey struct {
	t reflect.Type
	d int
}

var modelStructs = sync.Map{}

func fieldName2Index(value reflect.Value, name2Index map[string][]int, pre []int, depth int) {
	value = reflect.Indirect(value)
	kind := value.Kind()
	if kind != reflect.Struct || depth == 0 {
		return
	}
	v := reflect.Indirect(value)
	filedCount := v.NumField()
	for i := 0; i < filedCount; i++ {
		field := v.Type().Field(i)
		if !ast.IsExported(field.Name) {
			continue
		}
		fieldName := field.Name
		f := gorm.ToDBName(fieldName)
		name2Index[f] = append(pre, i)
		name2Index[fieldName] = append(pre, i)
		fieldName2Index(value.Field(i), name2Index, name2Index[fieldName], depth-1)
	}
}
