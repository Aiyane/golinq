package test

import (
	"encoding/json"
	"github.com/Aiyane/golinq"
	"github.com/Aiyane/golinq/interpreters"
	"github.com/Aiyane/golinq/types"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var (
	DATAS1 types.DataSources
	DATAS2 types.DataSources
	DATAS3 types.DataSources
	once1  sync.Once
	once2  sync.Once
	once3  sync.Once
)

func TestMain(t *testing.M) {
	interpreters.TagString = "sql"
	t.Run()
}

func TestJoinOn(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"name", "mobile_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "apple", "mobile_name": "iphone4", "ceo_name": "Tim Cook"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone5", "ceo_name": "Tim Cook"},
		map[string]interface{}{"name": "huawei", "mobile_name": "Mate20", "ceo_name": "renzhengfei"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2S", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi2", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi3", "ceo_name": "leijun"},
	}) {
		t.Error(res)
	}
}

func TestJoinOnWhere(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\""
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"mobile_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "apple", "mobile_name": "iphone4", "ceo_name": "Tim Cook"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone5", "ceo_name": "Tim Cook"},
	}) {
		t.Error(res)
	}
}

func TestPartJoinOnWhere(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company, ceo join mobile on company.id = mobile.company_id " +
		"where company.name = \"apple\""
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"mobile_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "apple", "mobile_name": "iphone4", "ceo_name": "leijun"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone4", "ceo_name": "leijun"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone4", "ceo_name": "Tim Cook"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone4", "ceo_name": "renzhengfei"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone5", "ceo_name": "leijun"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone5", "ceo_name": "leijun"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone5", "ceo_name": "Tim Cook"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone5", "ceo_name": "renzhengfei"},
	}) {
		t.Error(res)
	}
}

func TestJoinOnWhereAgg(t *testing.T) {
	sql := "select count(mobile.name) as mobile_count " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\" and mobile.name = \"iphone8\""
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"mobile_name"})
	if !isEqual(res, []interface{}{}) {
		t.Error(res)
	}
}

func TestJoinOnWhereOrderAgg(t *testing.T) {
	sql := "select count(mobile.name) as mobile_count " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\" order by mobile.name"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"mobile_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"mobile_count": 2},
	}) {
		t.Error(res)
	}
}

func TestJoinOnWhereOrderNoAgg(t *testing.T) {
	sql := "select count(mobile.name) as mobile_count " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\" and mobile.name = \"iphone8\" order by mobile.name"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"mobile_name"})
	if !isEqual(res, []interface{}{}) {
		t.Error(res)
	}
}

func TestJoinOnWhereDoubleNames(t *testing.T) {
	sql := "select * " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\""
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"mobile_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1.0, "company_name": "apple", "ceo_id": 1.0, "ceo_name": "Tim Cook",
			"ceo_company_id": 1.0, "mobile_id": 4.0, "mobile_name": "iphone4", "mobile_company_id": 1.0},
		map[string]interface{}{"company_id": 1.0, "company_name": "apple", "ceo_id": 1.0, "ceo_name": "Tim Cook",
			"ceo_company_id": 1.0, "mobile_id": 5.0, "mobile_name": "iphone5", "mobile_company_id": 1.0},
	}) {
		t.Error(res)
	}
}

func TestWhereDiffPriorityCalc(t *testing.T) {
	sql := "select id+2*3 as id " +
		"from company " +
		"where name = \"apple\""
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"id": 7.0},
	}) {
		t.Error(res)
	}
}

func TestWhereOr(t *testing.T) {
	sql := "select * " +
		"from mobile " +
		"where company_id = 1 OR id = 1"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"id"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"id": 1.0, "name": "xiaomi2", "company_id": 3.0},
		map[string]interface{}{"id": 4.0, "name": "iphone4", "company_id": 1.0},
		map[string]interface{}{"id": 5.0, "name": "iphone5", "company_id": 1.0},
	}) {
		t.Error(res)
	}
}

func TestCalcWhere(t *testing.T) {
	sql := "select * " +
		"from mobile " +
		"where company_id + id > 5"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"id", "name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"id": 3.0, "company_id": 4.0, "name": "Mix2"},
		map[string]interface{}{"id": 3.0, "company_id": 4.0, "name": "Mix2S"},
		map[string]interface{}{"id": 5.0, "company_id": 1.0, "name": "iphone5"},
		map[string]interface{}{"id": 6.0, "company_id": 2.0, "name": "Mate20"},
	}) {
		t.Error(res)
	}
}

func TestUpdate(t *testing.T) {
	sql := "update mobile set name='hello' " +
		"where company_id + id > 5"
	data := newData()
	res := golinq.SqlRun(sql, data, "").(int)
	assert.Equal(t, res, 4)
	sql2 := "select * " +
		"from mobile " +
		"where mobile.company_id + mobile.id > 5"
	res2 := golinq.SqlRun(sql2, data, "").([]interface{})
	sorted(&res2, []string{"id", "name"})
	if !isEqual(res2, []interface{}{
		map[string]interface{}{"id": 3.0, "company_id": 4.0, "name": "hello"},
		map[string]interface{}{"id": 3.0, "company_id": 4.0, "name": "hello"},
		map[string]interface{}{"id": 5.0, "company_id": 1.0, "name": "hello"},
		map[string]interface{}{"id": 6.0, "company_id": 2.0, "name": "hello"},
	}) {
		t.Error(res2)
	}
}

func TestUpdateOrder(t *testing.T) {
	sql := "update mobile set name='hello' " +
		"where id > 3 order by id limit 1,1"
	data := newData()
	res := golinq.SqlRun(sql, data, "").(int)
	assert.Equal(t, res, 1)
	sql2 := "select * " +
		"from mobile "
	res2 := golinq.SqlRun(sql2, data, "").([]interface{})
	sorted(&res2, []string{"id", "name"})
	if !isEqual(res2, []interface{}{
		map[string]interface{}{"id": 1.0, "company_id": 3.0, "name": "xiaomi2"},
		map[string]interface{}{"id": 2.0, "company_id": 3.0, "name": "xiaomi3"},
		map[string]interface{}{"id": 3.0, "company_id": 4.0, "name": "Mix2"},
		map[string]interface{}{"id": 3.0, "company_id": 4.0, "name": "Mix2S"},
		map[string]interface{}{"id": 4.0, "company_id": 1.0, "name": "iphone4"},
		map[string]interface{}{"id": 5.0, "company_id": 1.0, "name": "hello"},
		map[string]interface{}{"id": 6.0, "company_id": 2.0, "name": "Mate20"},
	}) {
		t.Error(res2)
	}
}

func TestUpdateOrderCase(t *testing.T) {
	sql := "update mobile set name=case when id = 4 then 'hello4' " +
		"when id = 5 then 'hello5' " +
		"else 'hello6' end " +
		"where mobile.id > 3 order by mobile.id"
	data := newData()
	res := golinq.SqlRun(sql, data, "").(int)
	assert.Equal(t, res, 3)
	sql2 := "select * " +
		"from mobile "
	res2 := golinq.SqlRun(sql2, data, "").([]interface{})
	sorted(&res2, []string{"id", "name"})
	if !isEqual(res2, []interface{}{
		map[string]interface{}{"id": 1.0, "company_id": 3.0, "name": "xiaomi2"},
		map[string]interface{}{"id": 2.0, "company_id": 3.0, "name": "xiaomi3"},
		map[string]interface{}{"id": 3.0, "company_id": 4.0, "name": "Mix2"},
		map[string]interface{}{"id": 3.0, "company_id": 4.0, "name": "Mix2S"},
		map[string]interface{}{"id": 4.0, "company_id": 1.0, "name": "hello4"},
		map[string]interface{}{"id": 5.0, "company_id": 1.0, "name": "hello5"},
		map[string]interface{}{"id": 6.0, "company_id": 2.0, "name": "hello6"},
	}) {
		t.Error(res2)
	}
}

func TestAndWhere(t *testing.T) {
	sql := "select * " +
		"from mobile " +
		"where company_id - id < 3 " +
		"and company_id - id > 0"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"id", "name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"id": 1.0, "name": "xiaomi2", "company_id": 3.0},
		map[string]interface{}{"id": 2.0, "name": "xiaomi3", "company_id": 3.0},
		map[string]interface{}{"id": 3.0, "name": "Mix2", "company_id": 4.0},
		map[string]interface{}{"id": 3.0, "name": "Mix2S", "company_id": 4.0},
	}) {
		t.Error(res)
	}
}

func TestWhereAs(t *testing.T) {
	sql := "select company.name as company_name, mobile.name as mobile_name " +
		"from mobile, company " +
		"where mobile.company_id = company.id " +
		"and mobile.name = 'xiaomi3'"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "xiaomi", "mobile_name": "xiaomi3"},
	}) {
		t.Error(res)
	}
}

func TestAndWhereAs(t *testing.T) {
	sql := "select mobile.name as mobile_name, company.name as company_name, ceo.name as ceo_name " +
		"from mobile, company, ceo " +
		"where mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"and ceo.name = 'renzhengfei'"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "huawei", "mobile_name": "Mate20", "ceo_name": "renzhengfei"},
	}) {
		t.Error(res)
	}
}

func TestFromAs(t *testing.T) {
	sql := "select m.name as mobile_name, cp.name as company_name, ceo.name as ceo_name " +
		"from mobile as m, company as cp, ceo " +
		"where m.company_id = company.id " +
		"and cp.id = ceo.company_id " +
		"and ceo.name = 'renzhengfei'"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "huawei", "mobile_name": "Mate20", "ceo_name": "renzhengfei"},
	}) {
		t.Error(res)
	}
}

func TestMulAggrFuncAs(t *testing.T) {
	sql := "select sum(id) as sum_mobile_id, sum(company_id) as sum_company_id " +
		"from mobile"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"sum_company_id": 18.0, "sum_mobile_id": 24.0},
	}) {
		t.Error(res)
	}
}

func TestAggrFuncAs(t *testing.T) {
	sql := "select count(*) as count1 from mobile"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"count1": 7},
	}) {
		t.Error(res)
	}
}

func TestMulAndWhere(t *testing.T) {
	sql := "select mobile.name as mobile_name, company.name as company_name, ceo.name as ceo_name " +
		"from mobile, company, ceo " +
		"where mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"and mobile.name = 'Mate20'"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "huawei", "mobile_name": "Mate20", "ceo_name": "renzhengfei"},
	}) {
		t.Error(res)
	}
}

func TestConstCompareWhere(t *testing.T) {
	sql := "select mobile.name as mobile_name, company.name as company_name, ceo.name as ceo_name " +
		"from mobile, company, ceo " +
		"where mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"and 'hello' = 'Mate20'"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{}) {
		t.Error(res)
	}
}

func TestSum(t *testing.T) {
	sql := "select sum(mobile.id) as sum1 from mobile"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"sum1": 24.0},
	}) {
		t.Error(res)
	}
}

func TestInWhere(t *testing.T) {
	sql := "select company_id, name " +
		"from mobile " +
		"where name in ('xiaomi2', 'xiaomi3', 'Mate20')"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "Mate20", "company_id": 2.0},
		map[string]interface{}{"name": "xiaomi2", "company_id": 3.0},
		map[string]interface{}{"name": "xiaomi3", "company_id": 3.0},
	}) {
		t.Error(res)
	}
}

func TestNotInWhere(t *testing.T) {
	sql := "select company_id, name " +
		"from mobile " +
		"where name not in ('xiaomi2', 'xiaomi3', 'Mate20')"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "Mix2", "company_id": 4.0},
		map[string]interface{}{"name": "Mix2S", "company_id": 4.0},
		map[string]interface{}{"name": "iphone4", "company_id": 1.0},
		map[string]interface{}{"name": "iphone5", "company_id": 1.0},
	}) {
		t.Error(res)
	}
}

func TestSubSelectExpr(t *testing.T) {
	sql := "select company_id, name " +
		"from mobile " +
		"where company_id in " +
		"(select id from company)"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"company_id", "name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "iphone4", "company_id": 1.0},
		map[string]interface{}{"name": "iphone5", "company_id": 1.0},
		map[string]interface{}{"name": "Mate20", "company_id": 2.0},
		map[string]interface{}{"name": "xiaomi2", "company_id": 3.0},
		map[string]interface{}{"name": "xiaomi3", "company_id": 3.0},
		map[string]interface{}{"name": "Mix2", "company_id": 4.0},
		map[string]interface{}{"name": "Mix2S", "company_id": 4.0},
	}) {
		t.Error(res)
	}
}

func TestCaseWhenElse(t *testing.T) {
	sql := "select " +
		"case when name = 'xiaomi' then '小米' " +
		"when name ='huawei' then '华为' " +
		"when name ='apple' then '苹果' " +
		"else '' end as company_name " +
		"from company"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"company_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "华为"},
		map[string]interface{}{"company_name": "小米"},
		map[string]interface{}{"company_name": "小米"},
		map[string]interface{}{"company_name": "苹果"},
	}) {
		t.Error(res)
	}
}

func TestUnicodeCharAs(t *testing.T) {
	sql := "select name as `公司`, id as ID from company"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"公司", "ID"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"公司": "apple", "ID": 1.0},
		map[string]interface{}{"公司": "huawei", "ID": 2.0},
		map[string]interface{}{"公司": "xiaomi", "ID": 3.0},
		map[string]interface{}{"公司": "xiaomi", "ID": 4.0},
	}) {
		t.Error(res)
	}
}

func TestWhereOrderByDesc(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"xiaomi\" " +
		"order by mobile.id desc, mobile.name"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2S", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi3", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi2", "ceo_name": "leijun"},
	}) {
		t.Error(res)
	}
}

func TestWhereOrderByDescLimit(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"xiaomi\" " +
		"order by mobile.id desc, mobile.name limit 1,2"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2S", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi3", "ceo_name": "leijun"},
	}) {
		t.Error(res)
	}
}

func TestWhereOrderBy(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name, mobile_price.price " +
		"from company join mobile join ceo on company.id = mobile.company_id " +
		"join mobile_price on company.id = ceo.company_id " +
		"where mobile_price.mobile_id = mobile.id " +
		"order by mobile_price.price, mobile.name"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi2", "ceo_name": "leijun", "price": 1000.0},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2", "ceo_name": "leijun", "price": 2000.0},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2S", "ceo_name": "leijun", "price": 2000.0},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi3", "ceo_name": "leijun", "price": 2000.0},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone4", "ceo_name": "Tim Cook", "price": 5000.0},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone5", "ceo_name": "Tim Cook", "price": 6000.0},
		map[string]interface{}{"name": "huawei", "mobile_name": "Mate20", "ceo_name": "renzhengfei", "price": 7000.0},
	}) {
		t.Error(res)
	}
}

func TestOrderBy(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"order by company_id asc"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1.0},
		map[string]interface{}{"company_id": 1.0},
		map[string]interface{}{"company_id": 2.0},
		map[string]interface{}{"company_id": 3.0},
		map[string]interface{}{"company_id": 3.0},
		map[string]interface{}{"company_id": 4.0},
		map[string]interface{}{"company_id": 4.0},
	}) {
		t.Error(res)
	}
}

func TestOrderByDesc(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"order by company_id desc"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 4.0},
		map[string]interface{}{"company_id": 4.0},
		map[string]interface{}{"company_id": 3.0},
		map[string]interface{}{"company_id": 3.0},
		map[string]interface{}{"company_id": 2.0},
		map[string]interface{}{"company_id": 1.0},
		map[string]interface{}{"company_id": 1.0},
	}) {
		t.Error(res)
	}
}

func TestDelete(t *testing.T) {
	sql := "delete from mobile " +
		"where company_id = 3"
	data := newData()
	res := golinq.SqlRun(sql, data, "").(int)
	assert.Equal(t, res, 2)
	res2 := golinq.SqlRun("select company_id "+
		"from mobile "+
		"order by company_id desc", data, "").([]interface{})
	if !isEqual(res2, []interface{}{
		map[string]interface{}{"company_id": 4.0},
		map[string]interface{}{"company_id": 4.0},
		map[string]interface{}{"company_id": 2.0},
		map[string]interface{}{"company_id": 1.0},
		map[string]interface{}{"company_id": 1.0},
	}) {
		t.Error(res2)
	}
}

func TestDeleteOrder(t *testing.T) {
	sql := "delete from mobile " +
		"where company_id = 3 or company_id = 2 or mobile.company_id = 4 " +
		"order by company_id desc limit 2"
	data := newData()
	res := golinq.SqlRun(sql, data, "").(int)
	assert.Equal(t, res, 2)
	res2 := golinq.SqlRun("select mobile.company_id "+
		"from mobile "+
		"order by company_id desc", data, "").([]interface{})
	if !isEqual(res2, []interface{}{
		map[string]interface{}{"company_id": 3.0},
		map[string]interface{}{"company_id": 3.0},
		map[string]interface{}{"company_id": 2.0},
		map[string]interface{}{"company_id": 1.0},
		map[string]interface{}{"company_id": 1.0},
	}) {
		t.Error(res2)
	}
}

func TestGroupBy(t *testing.T) {
	sql := "select company.name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"group by company.name"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "apple"},
		map[string]interface{}{"name": "huawei"},
		map[string]interface{}{"name": "xiaomi"},
	}) {
		t.Error(res)
	}
}

func TestGroupByDoubleNames(t *testing.T) {
	sql := "select * " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"group by company.name"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"company_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "apple", "mobile_id": nil, "mobile_name": nil, "company_id": nil,
			"ceo_id": nil, "ceo_name": nil, "ceo_company_id": nil, "mobile_company_id": nil},
		map[string]interface{}{"company_name": "huawei", "mobile_id": nil, "mobile_name": nil, "company_id": nil,
			"ceo_id": nil, "ceo_name": nil, "ceo_company_id": nil, "mobile_company_id": nil},
		map[string]interface{}{"company_name": "xiaomi", "mobile_id": nil, "mobile_name": nil, "company_id": nil,
			"ceo_id": nil, "ceo_name": nil, "ceo_company_id": nil, "mobile_company_id": nil},
	}) {
		t.Error(res)
	}
}

func TestGroupByFunc(t *testing.T) {
	sql := "select company_id, gcount(*) as count1 " +
		"from mobile " +
		"group by company_id"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"company_id"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1.0, "count1": 2},
		map[string]interface{}{"company_id": 2.0, "count1": 1},
		map[string]interface{}{"company_id": 3.0, "count1": 2},
		map[string]interface{}{"company_id": 4.0, "count1": 2},
	}) {
		t.Error(res)
	}
}

func TestGroupByHavingFunc(t *testing.T) {
	sql := "select company_id, gcount(*) as count1 " +
		"from mobile group by company_id " +
		"having count(*) >= 2 and company_id != 4"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"company_id"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1.0, "count1": 2},
		map[string]interface{}{"company_id": 3.0, "count1": 2},
	}) {
		t.Error(res)
	}
}

func TestGroupByNestingFunc(t *testing.T) {
	sql := "select gsum(mobile.id) as sum1, company_id " +
		"from mobile " +
		"group by company_id"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"company_id"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1.0, "sum1": 9.0},
		map[string]interface{}{"company_id": 2.0, "sum1": 6.0},
		map[string]interface{}{"company_id": 3.0, "sum1": 3.0},
		map[string]interface{}{"company_id": 4.0, "sum1": 6.0},
	}) {
		t.Error(res)
	}
}

func TestGroupByHavingOrderByDesc(t *testing.T) {
	sql := "select company.name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"group by company.name " +
		"having len(company.name) = 6 " +
		"order by company.name desc"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "xiaomi"},
		map[string]interface{}{"name": "huawei"},
	}) {
		t.Error(res)
	}
}

func TestGroupByOrderBy(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"group by company_id " +
		"order by company_id"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1.0},
		map[string]interface{}{"company_id": 2.0},
		map[string]interface{}{"company_id": 3.0},
		map[string]interface{}{"company_id": 4.0},
	}) {
		t.Error(res)
	}
}

func TestGroupByOrderByLimit(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"group by company_id " +
		"order by company_id limit 3,2"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 4.0},
	}) {
		t.Error(res)
	}
}

func TestGroupByWhereUnicodeCharAsFunc(t *testing.T) {
	sql := "select ceo.name as ceo_name, deflate(mobile_price.type, 0, mobile_price.price) as `普通`,  " +
		"deflate(mobile_price.type, 1, mobile_price.price) as `旗舰` " +
		"from mobile, company, ceo, mobile_price " +
		"where mobile.id = mobile_price.mobile_id " +
		"and mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"group by ceo.name"
	res := golinq.SqlRun(sql, getData(), "").([]interface{})
	sorted(&res, []string{"ceo_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"旗舰": 11000.0, "普通": 0, "ceo_name": "Tim Cook"},
		map[string]interface{}{"旗舰": 4000.0, "普通": 3000.0, "ceo_name": "leijun"},
		map[string]interface{}{"旗舰": 7000.0, "普通": 0, "ceo_name": "renzhengfei"},
	}) {
		t.Error(res)
	}
}

func TestBlankResult(t *testing.T) {
	sql := "select ceo.name as ceo_name, deflate(mobile_price.type, 0, mobile_price.price) as `普通`,  " +
		"deflate(mobile_price.type, 1, mobile_price.price) as `旗舰` " +
		"from mobile, company, ceo, mobile_price " +
		"where mobile.id = mobile_price.mobile_id " +
		"and mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"group by ceo.name"
	dataSources := getData()
	dataSources["company"] = []types.Record{}
	res := golinq.SqlRun(sql, dataSources, "").([]interface{})
	if !isEqual(res, []interface{}{}) {
		t.Error(res)
	}
}

func TestMulDotOperator(t *testing.T) {
	sql := "select m.mobile.company.ceo as ceo_name from mobile_price as m " +
		"join ceo where m.mobile.company.id = ceo.company_id " +
		"group by ceo.name having ceo.name != \"renzhengfei\" order by ceo.name"
	dataSources := getData2()
	res := golinq.SqlRun(sql, dataSources, "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"ceo_name": "Tim Cook"},
		map[string]interface{}{"ceo_name": "leijun"},
	}) {
		t.Error(res)
	}
}

func TestListField(t *testing.T) {
	sql := "select t.name, s.number from Teacher as t, Student as s where t.id = s.id " +
		"group by t.name, Student.number order by t.name"
	dataSources := getData3()
	res := golinq.SqlRun(sql, dataSources, "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": []interface{}{"jay", "tom", "will"}, "number": 10.0},
		map[string]interface{}{"name": []interface{}{"zhang", "li", "wang"}, "number": 10.0},
	}) {
		t.Error(res)
	}
}

func TestDictFieldOrderMulDot(t *testing.T) {
	sql := "select m.mobile.company from mobile_price as m " +
		"join ceo where m.mobile.company.id = ceo.company_id " +
		"and m.id not in (1, 2, 4) " +
		"group by ceo.name having ceo.name != \"renzhengfei\" order by m.mobile.company.name"
	dataSources := getData2()
	res := golinq.SqlRun(sql, dataSources, "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company": map[string]interface{}{"id": 1.0, "name": "apple", "ceo": "Tim Cook"}},
		map[string]interface{}{"company": map[string]interface{}{"id": 4.0, "name": "xiaomi", "ceo": "leijun"}},
	}) {
		t.Error(res)
	}
}

func TestGroupByHavingMulDot(t *testing.T) {
	sql := "select m.mobile.company from mobile_price as m " +
		"join ceo where m.mobile.company.id = ceo.company_id " +
		"group by m.mobile.company.name having m.mobile.company.name != \"apple\" order by m.mobile.company.name"
	dataSources := getData2()
	res := golinq.SqlRun(sql, dataSources, "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company": map[string]interface{}{"id": 2.0, "name": "huawei", "ceo": "renzhengfei"}},
		map[string]interface{}{"company": map[string]interface{}{"id": 4.0, "name": "xiaomi", "ceo": "leijun"}},
	}) {
		t.Error(res)
	}
}

func TestInsertExpr(t *testing.T) {
	sql := "INSERT company (`id`, `name`) VALUES (5, 'oppo')"
	dataSources := newData()
	res := golinq.SqlRun(sql, dataSources, "")
	assert.Equal(t, res, 5)
	sql2 := "select * from company where company.id = 5"
	res2 := golinq.SqlRun(sql2, dataSources, "").([]interface{})
	if !isEqual(res2, []interface{}{
		map[string]interface{}{"id": 5, "name": "oppo"},
	}) {
		t.Error(res2)
	}
}

func TestInsertSelect(t *testing.T) {
	golinq.ClearType()
	sql := "INSERT company (`id`, `name`) SELECT id, name FROM company WHERE id = 1"
	dataSources := newData()
	res := golinq.SqlRun(sql, dataSources, "")
	assert.Equal(t, res, 1)
	sql2 := "select * from company where company.id = 1"
	res2 := golinq.SqlRun(sql2, dataSources, "").([]interface{})
	if !isEqual(res2, []interface{}{
		map[string]interface{}{"id": 1.0, "name": "apple"},
		map[string]interface{}{"id": 1.0, "name": "apple"},
	}) {
		t.Error(res2)
	}
}

func getData() types.DataSources {
	once1.Do(func() {
		err := json.Unmarshal([]byte(
			`{
			"company": [
				{"id": 1, "name": "apple"},
				{"id": 2, "name": "huawei"},
				{"id": 3, "name": "xiaomi"},
				{"id": 4, "name": "xiaomi"}
			],
			"mobile": [
				{"id": 1, "name": "xiaomi2", "company_id": 3},
				{"id": 2, "name": "xiaomi3", "company_id": 3},
				{"id": 3, "name": "Mix2", "company_id": 4},
				{"id": 3, "name": "Mix2S", "company_id": 4},
				{"id": 4, "name": "iphone4", "company_id": 1},
				{"id": 5, "name": "iphone5", "company_id": 1},
				{"id": 6, "name": "Mate20", "company_id": 2}
			],
			"ceo": [
				{"id": 1, "name": "Tim Cook", "company_id": 1},
				{"id": 2, "name": "renzhengfei", "company_id": 2},
				{"id": 3, "name": "leijun", "company_id": 3},
				{"id": 4, "name": "leijun", "company_id": 4}
			],
			"mobile_price": [
				{"id": 1, "mobile_id": 1, "type": 0, "price": 1000},
				{"id": 2, "mobile_id": 2, "type": 0, "price": 2000},
				{"id": 3, "mobile_id": 3, "type": 1, "price": 2000},
				{"id": 4, "mobile_id": 4, "type": 1, "price": 5000},
				{"id": 5, "mobile_id": 5, "type": 1, "price": 6000},
				{"id": 6, "mobile_id": 6, "type": 1, "price": 7000}
			]
		}`), &DATAS1)
		if err != nil {
			panic("出现错误！")
		}
	})
	return DATAS1
}

func newData() types.DataSources {
	var DATAS types.DataSources
	err := json.Unmarshal([]byte(
		`{
			"company": [
				{"id": 1, "name": "apple"},
				{"id": 2, "name": "huawei"},
				{"id": 3, "name": "xiaomi"},
				{"id": 4, "name": "xiaomi"}
			],
			"mobile": [
				{"id": 1, "name": "xiaomi2", "company_id": 3},
				{"id": 2, "name": "xiaomi3", "company_id": 3},
				{"id": 3, "name": "Mix2", "company_id": 4},
				{"id": 3, "name": "Mix2S", "company_id": 4},
				{"id": 4, "name": "iphone4", "company_id": 1},
				{"id": 5, "name": "iphone5", "company_id": 1},
				{"id": 6, "name": "Mate20", "company_id": 2}
			],
			"ceo": [
				{"id": 1, "name": "Tim Cook", "company_id": 1},
				{"id": 2, "name": "renzhengfei", "company_id": 2},
				{"id": 3, "name": "leijun", "company_id": 3},
				{"id": 4, "name": "leijun", "company_id": 4}
			],
			"mobile_price": [
				{"id": 1, "mobile_id": 1, "type": 0, "price": 1000},
				{"id": 2, "mobile_id": 2, "type": 0, "price": 2000},
				{"id": 3, "mobile_id": 3, "type": 1, "price": 2000},
				{"id": 4, "mobile_id": 4, "type": 1, "price": 5000},
				{"id": 5, "mobile_id": 5, "type": 1, "price": 6000},
				{"id": 6, "mobile_id": 6, "type": 1, "price": 7000}
			]
		}`), &DATAS)
	if err != nil {
		panic("出现错误！")
	}
	return DATAS
}

func getData2() types.DataSources {
	once2.Do(func() {
		err := json.Unmarshal([]byte(
			`{
			"ceo": [
				{"id": 1, "name": "Tim Cook", "company_id": 1},
				{"id": 2, "name": "renzhengfei", "company_id": 2},
				{"id": 3, "name": "leijun", "company_id": 3},
				{"id": 4, "name": "leijun", "company_id": 4}
			],
			"mobile_price": [
				{"id": 1, "mobile": {"id": 1, "name": "xiaomi2", "company": {"id": 3, "name": "xiaomi", "ceo": "leijun"}}, "type": 0, "price": 1000},
				{"id": 2, "mobile": {"id": 2, "name": "xiaomi3", "company": {"id": 3, "name": "xiaomi", "ceo": "leijun"}}, "type": 0, "price": 2000},
				{"id": 3, "mobile": {"id": 3, "name": "Mix2", "company": {"id": 4, "name": "xiaomi", "ceo": "leijun"}}, "type": 1, "price": 2000},
				{"id": 4, "mobile": {"id": 4, "name": "iphone4", "company": {"id": 1, "name": "apple", "ceo": "Tim Cook"}}, "type": 1, "price": 5000},
				{"id": 5, "mobile": {"id": 5, "name": "iphone5", "company": {"id": 1, "name": "apple", "ceo": "Tim Cook"}}, "type": 1, "price": 6000},
				{"id": 6, "mobile": {"id": 6, "name": "Mate20", "company": {"id": 2, "name": "huawei", "ceo": "renzhengfei"}}, "type": 1, "price": 7000}
			]
		}`), &DATAS2)
		if err != nil {
			panic("出现错误！")
		}
	})
	return DATAS2
}

func getData3() types.DataSources {
	once3.Do(func() {
		err := json.Unmarshal([]byte(
			`{
			"Teacher": [
				{"name": ["zhang", "li", "wang"], "year": 5, "id": 1},
				{"name": ["jay", "tom", "will"], "year": 3, "id": 2}
			],
			"Student": [
				{"number": 10, "age": 18, "id": 1},
				{"number": 10, "age": 19, "id": 2},
				{"number": 10, "age": 20, "id": 3},
				{"number": 20, "age": 21, "id": 4}
			]
		}`), &DATAS3)
		if err != nil {
			panic("出现错误！")
		}
	})
	return DATAS3
}
