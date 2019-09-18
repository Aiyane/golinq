package test

import (
	"github.com/Aiyane/golinq"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestModelInsertExpr(t *testing.T) {
	sql := "INSERT company (`id`, `name`) VALUES (5, 'oppo')"
	dataSources := golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices())
	res := golinq.SqlRun(sql, dataSources, "")
	assert.Equal(t, res, 5)
	sql2 := "select * from company where id = 5"
	res2 := golinq.SqlRun(sql2, dataSources, "").([]interface{})
	if !isEqual(res2, []interface{}{
		map[string]interface{}{"id": 5, "name": "oppo"},
	}) {
		t.Error(res2)
	}
}

func TestModelInsertExpr2(t *testing.T) {
	sql := "INSERT company (`id`, `name`) VALUES (5, 'oppo')"
	dataSources := golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices())
	golinq.RegisterType((*Company)(nil))
	res := golinq.SqlRun(sql, dataSources, "Company")
	assert.Equal(t, res, 5)
	sql2 := "select * from company where id = 5"
	res2 := golinq.SqlRun(sql2, dataSources, "Company").([]*Company)
	assert.Equal(t, res2[0].Id, 5)
	assert.Equal(t, res2[0].Name, "oppo")
}

func TestModelInsertSelect(t *testing.T) {
	sql := "INSERT company (`id`, `name`) SELECT id, name FROM company WHERE id = 1"
	dataSources := golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices())
	res := golinq.SqlRun(sql, dataSources, "")
	assert.Equal(t, res, 1)
	sql2 := "select * from company where company.id = 1"
	res2 := golinq.SqlRun(sql2, dataSources, "").([]interface{})
	if !isEqual(res2, []interface{}{
		map[string]interface{}{"id": 1, "name": "apple"},
		map[string]interface{}{"id": 1, "name": "apple"},
	}) {
		t.Error(res2)
	}
}

func TestModelJoinOn(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
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

type Res1 struct {
	Name       string `sql:"name"`
	MobileName string `sql:"mobile_name"`
	CeoName    string `sql:"ceo_name"`
}

func TestModelJoinOn2(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"order by company.name, mobile.name"
	golinq.RegisterType((*Res1)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res1").([]*Res1)
	r1 := &Res1{"apple", "iphone4", "Tim Cook"}
	r2 := &Res1{"apple", "iphone5", "Tim Cook"}
	r3 := &Res1{"huawei", "Mate20", "renzhengfei"}
	r4 := &Res1{"xiaomi", "Mix2", "leijun"}
	r5 := &Res1{"xiaomi", "Mix2S", "leijun"}
	r6 := &Res1{"xiaomi", "xiaomi2", "leijun"}
	r7 := &Res1{"xiaomi", "xiaomi3", "leijun"}
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[0].MobileName, r1.MobileName)
	assert.Equal(t, res[0].CeoName, r1.CeoName)
	assert.Equal(t, res[1].Name, r2.Name)
	assert.Equal(t, res[1].MobileName, r2.MobileName)
	assert.Equal(t, res[1].CeoName, r2.CeoName)
	assert.Equal(t, res[2].Name, r3.Name)
	assert.Equal(t, res[2].MobileName, r3.MobileName)
	assert.Equal(t, res[2].CeoName, r3.CeoName)
	assert.Equal(t, res[3].Name, r4.Name)
	assert.Equal(t, res[3].MobileName, r4.MobileName)
	assert.Equal(t, res[3].CeoName, r4.CeoName)
	assert.Equal(t, res[4].Name, r5.Name)
	assert.Equal(t, res[4].MobileName, r5.MobileName)
	assert.Equal(t, res[4].CeoName, r5.CeoName)
	assert.Equal(t, res[5].Name, r6.Name)
	assert.Equal(t, res[5].MobileName, r6.MobileName)
	assert.Equal(t, res[5].CeoName, r6.CeoName)
	assert.Equal(t, res[6].Name, r7.Name)
	assert.Equal(t, res[6].MobileName, r7.MobileName)
	assert.Equal(t, res[6].CeoName, r7.CeoName)
}

func TestModelJoinOnWhere(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\""
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"mobile_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "apple", "mobile_name": "iphone4", "ceo_name": "Tim Cook"},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone5", "ceo_name": "Tim Cook"},
	}) {
		t.Error(res)
	}
}

func TestModelJoinOnWhere2(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\" order by mobile.name"
	golinq.RegisterType((*Res1)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res1").([]*Res1)
	r1 := &Res1{"apple", "iphone4", "Tim Cook"}
	r2 := &Res1{"apple", "iphone5", "Tim Cook"}
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[0].MobileName, r1.MobileName)
	assert.Equal(t, res[0].CeoName, r1.CeoName)
	assert.Equal(t, res[1].Name, r2.Name)
	assert.Equal(t, res[1].MobileName, r2.MobileName)
	assert.Equal(t, res[1].CeoName, r2.CeoName)
}

func TestModelJoinOnWhereDoubleNames(t *testing.T) {
	sql := "select * " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\""
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"mobile_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1, "company_name": "apple", "ceo_id": 1, "ceo_name": "Tim Cook",
			"ceo_company_id": 1, "mobile_id": 4, "mobile_name": "iphone4", "mobile_company_id": 1},
		map[string]interface{}{"company_id": 1, "company_name": "apple", "ceo_id": 1, "ceo_name": "Tim Cook",
			"ceo_company_id": 1, "mobile_id": 5, "mobile_name": "iphone5", "mobile_company_id": 1},
	}) {
		t.Error(res)
	}
}

type Res2 struct {
	CompanyId       int    `sql:"company_id"`
	CompanyName     string `sql:"company_name"`
	CeoId           int    `sql:"ceo_id"`
	CeoName         string `sql:"ceo_name"`
	CeoCompanyId    int    `sql:"ceo_company_id"`
	MobileId        int    `sql:"mobile_id"`
	MobileName      string `sql:"mobile_name"`
	MobileCompanyId int    `sql:"mobile_company_id"`
}

func TestModelJoinOnWhereDoubleNames2(t *testing.T) {
	sql := "select * " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\" order by mobile.name"
	golinq.RegisterType((*Res2)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res2").([]*Res2)
	r1 := &Res2{1, "apple", 1, "Tim Cook", 1, 4, "iphone4", 1}
	r2 := &Res2{1, "apple", 1, "Tim Cook", 1, 5, "iphone5", 1}
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
	assert.Equal(t, res[0].CompanyName, r1.CompanyName)
	assert.Equal(t, res[0].CeoId, r1.CeoId)
	assert.Equal(t, res[0].CeoName, r1.CeoName)
	assert.Equal(t, res[0].CeoCompanyId, r1.CeoCompanyId)
	assert.Equal(t, res[0].MobileId, r1.MobileId)
	assert.Equal(t, res[0].MobileName, r1.MobileName)
	assert.Equal(t, res[0].MobileCompanyId, r1.MobileCompanyId)
	assert.Equal(t, res[1].CompanyId, r2.CompanyId)
	assert.Equal(t, res[1].CompanyName, r2.CompanyName)
	assert.Equal(t, res[1].CeoId, r2.CeoId)
	assert.Equal(t, res[1].CeoName, r2.CeoName)
	assert.Equal(t, res[1].CeoCompanyId, r2.CeoCompanyId)
	assert.Equal(t, res[1].MobileId, r2.MobileId)
	assert.Equal(t, res[1].MobileName, r2.MobileName)
	assert.Equal(t, res[1].MobileCompanyId, r2.MobileCompanyId)
}

func TestModelWhereDiffPriorityCalc(t *testing.T) {
	sql := "select id+2*3 as id " +
		"from company " +
		"where name = \"apple\""
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"id": 7},
	}) {
		t.Error(res)
	}
}

type Res3 struct {
	Id int `sql:"id"`
}

func TestModelWhereDiffPriorityCalc2(t *testing.T) {
	sql := "select id+2*3 as id " +
		"from company " +
		"where name = \"apple\""
	golinq.RegisterType((*Res3)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys()), "Res3").([]*Res3)
	r1 := &Res3{7}
	assert.Equal(t, res[0].Id, r1.Id)
}

func TestModelWhereOr(t *testing.T) {
	sql := "select * " +
		"from mobile " +
		"where company_id = 1 OR id = 1"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"id"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"id": 1, "name": "xiaomi2", "company_id": 3},
		map[string]interface{}{"id": 4, "name": "iphone4", "company_id": 1},
		map[string]interface{}{"id": 5, "name": "iphone5", "company_id": 1},
	}) {
		t.Error(res)
	}
}

func TestModelWhereOr2(t *testing.T) {
	sql := "select * " +
		"from mobile " +
		"where company_id = 1 OR id = 1 order by id"
	golinq.RegisterType((*Mobile)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Mobile").([]*Mobile)
	r1 := &Mobile{1, "xiaomi2", 3}
	r2 := &Mobile{4, "iphone4", 1}
	r3 := &Mobile{5, "iphone5", 1}
	assert.Equal(t, res[0].Id, r1.Id)
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
	assert.Equal(t, res[1].Id, r2.Id)
	assert.Equal(t, res[1].Name, r2.Name)
	assert.Equal(t, res[1].CompanyId, r2.CompanyId)
	assert.Equal(t, res[2].Id, r3.Id)
	assert.Equal(t, res[2].Name, r3.Name)
	assert.Equal(t, res[2].CompanyId, r3.CompanyId)
}

func TestModelCalcWhere(t *testing.T) {
	sql := "select * " +
		"from mobile " +
		"where company_id + id > 5"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"id", "name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"id": 3, "company_id": 4, "name": "Mix2"},
		map[string]interface{}{"id": 3, "company_id": 4, "name": "Mix2S"},
		map[string]interface{}{"id": 5, "company_id": 1, "name": "iphone5"},
		map[string]interface{}{"id": 6, "company_id": 2, "name": "Mate20"},
	}) {
		t.Error(res)
	}
}

func TestModelCalcWhere2(t *testing.T) {
	sql := "select * " +
		"from mobile " +
		"where company_id + id > 5 order by id, name"
	golinq.RegisterType((*Mobile)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Mobile").([]*Mobile)
	r1 := &Mobile{3, "Mix2", 4}
	r2 := &Mobile{3, "Mix2S", 4}
	r3 := &Mobile{5, "iphone5", 1}
	r4 := &Mobile{6, "Mate20", 2}
	assert.Equal(t, res[0].Id, r1.Id)
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
	assert.Equal(t, res[1].Id, r2.Id)
	assert.Equal(t, res[1].Name, r2.Name)
	assert.Equal(t, res[1].CompanyId, r2.CompanyId)
	assert.Equal(t, res[2].Id, r3.Id)
	assert.Equal(t, res[2].Name, r3.Name)
	assert.Equal(t, res[2].CompanyId, r3.CompanyId)
	assert.Equal(t, res[3].Id, r4.Id)
	assert.Equal(t, res[3].Name, r4.Name)
	assert.Equal(t, res[3].CompanyId, r4.CompanyId)
}

func TestModelAndWhere(t *testing.T) {
	sql := "select * " +
		"from mobile " +
		"where company_id - id < 3 " +
		"and company_id - id > 0"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"id", "name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"id": 1, "name": "xiaomi2", "company_id": 3},
		map[string]interface{}{"id": 2, "name": "xiaomi3", "company_id": 3},
		map[string]interface{}{"id": 3, "name": "Mix2", "company_id": 4},
		map[string]interface{}{"id": 3, "name": "Mix2S", "company_id": 4},
	}) {
		t.Error(res)
	}
}

func TestModelAndWhere2(t *testing.T) {
	sql := "select * " +
		"from mobile " +
		"where company_id - id < 3 " +
		"and company_id - id > 0 order by id, name"
	golinq.RegisterType((*Mobile)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Mobile").([]*Mobile)
	r1 := &Mobile{1, "xiaomi2", 3}
	r2 := &Mobile{2, "xiaomi3", 3}
	r3 := &Mobile{3, "Mix2", 4}
	r4 := &Mobile{3, "Mix2S", 4}
	assert.Equal(t, res[0].Id, r1.Id)
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
	assert.Equal(t, res[1].Id, r2.Id)
	assert.Equal(t, res[1].Name, r2.Name)
	assert.Equal(t, res[1].CompanyId, r2.CompanyId)
	assert.Equal(t, res[2].Id, r3.Id)
	assert.Equal(t, res[2].Name, r3.Name)
	assert.Equal(t, res[2].CompanyId, r3.CompanyId)
	assert.Equal(t, res[3].Id, r4.Id)
	assert.Equal(t, res[3].Name, r4.Name)
	assert.Equal(t, res[3].CompanyId, r4.CompanyId)
}

type Res4 struct {
	CompanyName string `sql:"company_name"`
	MobileName  string `sql:"mobile_name"`
}

func TestModelWhereAs(t *testing.T) {
	sql := "select company.name as company_name, mobile.name as mobile_name " +
		"from mobile, company " +
		"where mobile.company_id = company.id " +
		"and mobile.name = 'xiaomi3'"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "xiaomi", "mobile_name": "xiaomi3"},
	}) {
		t.Error(res)
	}
}

func TestModelWhereAs2(t *testing.T) {
	sql := "select company.name as company_name, mobile.name as mobile_name " +
		"from mobile, company " +
		"where mobile.company_id = company.id " +
		"and mobile.name = 'xiaomi3'"
	golinq.RegisterType((*Res4)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res4").([]*Res4)
	r1 := &Res4{"xiaomi", "xiaomi3"}
	assert.Equal(t, res[0].CompanyName, r1.CompanyName)
	assert.Equal(t, res[0].MobileName, r1.MobileName)
}

type Res5 struct {
	CompanyName string `sql:"company_name"`
	MobileName  string `sql:"mobile_name"`
	CeoName     string `sql:"ceo_name"`
}

func TestModelAndWhereAs(t *testing.T) {
	sql := "select mobile.name as mobile_name, company.name as company_name, ceo.name as ceo_name " +
		"from mobile, company, ceo " +
		"where mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"and ceo.name = 'renzhengfei'"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "huawei", "mobile_name": "Mate20", "ceo_name": "renzhengfei"},
	}) {
		t.Error(res)
	}
}

func TestModelAndWhereAs2(t *testing.T) {
	sql := "select mobile.name as mobile_name, company.name as company_name, ceo.name as ceo_name " +
		"from mobile, company, ceo " +
		"where mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"and ceo.name = 'renzhengfei'"
	golinq.RegisterType((*Res5)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res5").([]*Res5)
	r1 := &Res5{"huawei", "Mate20", "renzhengfei"}
	assert.Equal(t, res[0].CompanyName, r1.CompanyName)
	assert.Equal(t, res[0].MobileName, r1.MobileName)
	assert.Equal(t, res[0].CeoName, r1.CeoName)
}

func TestModelFromAs(t *testing.T) {
	sql := "select m.name as mobile_name, cp.name as company_name, ceo.name as ceo_name " +
		"from mobile as m, company as cp, ceo " +
		"where m.company_id = company.id " +
		"and cp.id = ceo.company_id " +
		"and ceo.name = 'renzhengfei'"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "huawei", "mobile_name": "Mate20", "ceo_name": "renzhengfei"},
	}) {
		t.Error(res)
	}
}

func TestModelFromAs2(t *testing.T) {
	sql := "select m.name as mobile_name, cp.name as company_name, ceo.name as ceo_name " +
		"from mobile as m, company as cp, ceo " +
		"where m.company_id = company.id " +
		"and cp.id = ceo.company_id " +
		"and ceo.name = 'renzhengfei'"
	golinq.RegisterType((*Res5)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res5").([]*Res5)
	r1 := &Res5{"huawei", "Mate20", "renzhengfei"}
	assert.Equal(t, res[0].CompanyName, r1.CompanyName)
	assert.Equal(t, res[0].MobileName, r1.MobileName)
	assert.Equal(t, res[0].CeoName, r1.CeoName)
}

type Res6 struct {
	SumCompanyId int `sql:"sum_company_id"`
	SumMobileId  int `sql:"sum_mobile_id"`
}

func TestModelMulAggrFuncAs(t *testing.T) {
	sql := "select int(sum(id)) as sum_mobile_id, int(sum(company_id)) as sum_company_id " +
		"from mobile"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"sum_company_id": 18, "sum_mobile_id": 24},
	}) {
		t.Error(res)
	}
}

func TestModelMulAggrFuncAs2(t *testing.T) {
	sql := "select int(sum(id)) as sum_mobile_id, int(sum(company_id)) as sum_company_id " +
		"from mobile"
	golinq.RegisterType((*Res6)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res6").([]*Res6)
	r1 := &Res6{18, 24}
	assert.Equal(t, res[0].SumCompanyId, r1.SumCompanyId)
	assert.Equal(t, res[0].SumMobileId, r1.SumMobileId)
}

type Res7 struct {
	Count float64 `sql:"count1"`
}

func TestModelAggrFuncAs(t *testing.T) {
	sql := "select count(*) as count1 from mobile"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"count1": 7},
	}) {
		t.Error(res)
	}
}

func TestModelAggrFuncAs2(t *testing.T) {
	sql := "select count(*) as count1 from mobile"
	golinq.RegisterType((*Res7)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res7").([]*Res7)
	r1 := &Res7{7}
	assert.Equal(t, res[0].Count, r1.Count)
}

func TestModelMulAndWhere(t *testing.T) {
	sql := "select mobile.name as mobile_name, company.name as company_name, ceo.name as ceo_name " +
		"from mobile, company, ceo " +
		"where mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"and mobile.name = 'Mate20'"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_name": "huawei", "mobile_name": "Mate20", "ceo_name": "renzhengfei"},
	}) {
		t.Error(res)
	}
}

func TestModelMulAndWhere2(t *testing.T) {
	sql := "select mobile.name as mobile_name, company.name as company_name, ceo.name as ceo_name " +
		"from mobile, company, ceo " +
		"where mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"and mobile.name = 'Mate20'"
	golinq.RegisterType((*Res5)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res5").([]*Res5)
	r1 := &Res5{"huawei", "Mate20", "renzhengfei"}
	assert.Equal(t, res[0].CompanyName, r1.CompanyName)
	assert.Equal(t, res[0].MobileName, r1.MobileName)
	assert.Equal(t, res[0].CeoName, r1.CeoName)
}

func TestModelConstCompareWhere(t *testing.T) {
	sql := "select mobile.name as mobile_name, company.name as company_name, ceo.name as ceo_name " +
		"from mobile, company, ceo " +
		"where mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"and 'hello' = 'Mate20'"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{}) {
		t.Error(res)
	}
}

type Res8 struct {
	Sum float64 `sql:"sum1"`
}

func TestModelSum(t *testing.T) {
	sql := "select sum(mobile.id) as sum1 from mobile"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"sum1": 24},
	}) {
		t.Error(res)
	}
}

func TestModelSum2(t *testing.T) {
	sql := "select sum(mobile.id) as sum1 from mobile"
	golinq.RegisterType((*Res8)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res8").([]*Res8)
	r1 := &Res8{24}
	assert.Equal(t, res[0].Sum, r1.Sum)
}

type Res9 struct {
	Name      string `sql:"name"`
	CompanyId int    `sql:"company_id"`
	Id        int    `sql:"id"`
}

func TestModelInWhere(t *testing.T) {
	sql := "select company_id, name " +
		"from mobile " +
		"where name in ('xiaomi2', 'xiaomi3', 'Mate20')"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "Mate20", "company_id": 2},
		map[string]interface{}{"name": "xiaomi2", "company_id": 3},
		map[string]interface{}{"name": "xiaomi3", "company_id": 3},
	}) {
		t.Error(res)
	}
}

func TestModelInWhere2(t *testing.T) {
	sql := "select company_id, name " +
		"from mobile " +
		"where name in ('xiaomi2', 'xiaomi3', 'Mate20') order by name"
	golinq.RegisterType((*Res9)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res9").([]*Res9)
	r1 := &Res9{"Mate20", 2, 0}
	r2 := &Res9{"xiaomi2", 3, 0}
	r3 := &Res9{"xiaomi3", 3, 0}
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
	assert.Equal(t, res[1].Name, r2.Name)
	assert.Equal(t, res[1].CompanyId, r2.CompanyId)
	assert.Equal(t, res[2].Name, r3.Name)
	assert.Equal(t, res[2].CompanyId, r3.CompanyId)
}

func TestModelNotInWhere(t *testing.T) {
	sql := "select company_id, name " +
		"from mobile " +
		"where name not in ('xiaomi2', 'xiaomi3', 'Mate20')"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "Mix2", "company_id": 4},
		map[string]interface{}{"name": "Mix2S", "company_id": 4},
		map[string]interface{}{"name": "iphone4", "company_id": 1},
		map[string]interface{}{"name": "iphone5", "company_id": 1},
	}) {
		t.Error(res)
	}
}

func TestModelSubSelectExpr(t *testing.T) {
	sql := "select company_id, name " +
		"from mobile " +
		"where company_id in " +
		"(select id from company)"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"company_id", "name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "iphone4", "company_id": 1},
		map[string]interface{}{"name": "iphone5", "company_id": 1},
		map[string]interface{}{"name": "Mate20", "company_id": 2},
		map[string]interface{}{"name": "xiaomi2", "company_id": 3},
		map[string]interface{}{"name": "xiaomi3", "company_id": 3},
		map[string]interface{}{"name": "Mix2", "company_id": 4},
		map[string]interface{}{"name": "Mix2S", "company_id": 4},
	}) {
		t.Error(res)
	}
}

// todo 子select查询有坑, 必选在 struct 中有子select的字段
func TestModelSubSelectExpr2(t *testing.T) {
	sql := "select company_id, name " +
		"from mobile " +
		"where company_id in " +
		"(select id from company) order by company_id, name"
	golinq.RegisterType((*Res9)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res9").([]*Res9)
	r1 := &Res9{"iphone4", 1, 0}
	r2 := &Res9{"iphone5", 1, 0}
	r3 := &Res9{"Mate20", 2, 0}
	r4 := &Res9{"xiaomi2", 3, 0}
	r5 := &Res9{"xiaomi3", 3, 0}
	r6 := &Res9{"Mix2", 4, 0}
	r7 := &Res9{"Mix2S", 4, 0}
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
	assert.Equal(t, res[1].Name, r2.Name)
	assert.Equal(t, res[1].CompanyId, r2.CompanyId)
	assert.Equal(t, res[2].Name, r3.Name)
	assert.Equal(t, res[2].CompanyId, r3.CompanyId)
	assert.Equal(t, res[3].Name, r4.Name)
	assert.Equal(t, res[3].CompanyId, r4.CompanyId)
	assert.Equal(t, res[4].Name, r5.Name)
	assert.Equal(t, res[4].CompanyId, r5.CompanyId)
	assert.Equal(t, res[5].Name, r6.Name)
	assert.Equal(t, res[5].CompanyId, r6.CompanyId)
	assert.Equal(t, res[6].Name, r7.Name)
	assert.Equal(t, res[6].CompanyId, r7.CompanyId)
}

func TestModelCaseWhenElse(t *testing.T) {
	sql := "select " +
		"case when name = 'xiaomi' then '小米' " +
		"when name ='huawei' then '华为' " +
		"when name ='apple' then '苹果' " +
		"else '' end as company_name " +
		"from company"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
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

func TestModelCaseWhenElse2(t *testing.T) {
	sql := "select " +
		"case when name = 'xiaomi' then '小米' " +
		"when name ='huawei' then '华为' " +
		"when name ='apple' then '苹果' " +
		"else '' end as company_name " +
		"from company order by name"
	golinq.RegisterType((*Res5)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res5").([]*Res5)
	r1 := &Res5{CompanyName: "苹果"}
	r2 := &Res5{CompanyName: "华为"}
	r3 := &Res5{CompanyName: "小米"}
	r4 := &Res5{CompanyName: "小米"}
	assert.Equal(t, res[0].CompanyName, r1.CompanyName)
	assert.Equal(t, res[1].CompanyName, r2.CompanyName)
	assert.Equal(t, res[2].CompanyName, r3.CompanyName)
	assert.Equal(t, res[3].CompanyName, r4.CompanyName)
}

func TestModelUnicodeCharAs(t *testing.T) {
	sql := "select company.name as `公司`, company.id as ID from company"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"公司", "ID"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"公司": "apple", "ID": 1},
		map[string]interface{}{"公司": "huawei", "ID": 2},
		map[string]interface{}{"公司": "xiaomi", "ID": 3},
		map[string]interface{}{"公司": "xiaomi", "ID": 4},
	}) {
		t.Error(res)
	}
}

func TestModelWhereOrderByDesc(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"xiaomi\" " +
		"order by mobile.id desc, mobile.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2S", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi3", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi2", "ceo_name": "leijun"},
	}) {
		t.Error(res)
	}
}

func TestModelWhereOrderByDesc2(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"xiaomi\" " +
		"order by mobile.id desc, mobile.name"
	golinq.RegisterType((*Res1)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res1").([]*Res1)
	r4 := &Res1{"xiaomi", "Mix2", "leijun"}
	r5 := &Res1{"xiaomi", "Mix2S", "leijun"}
	r6 := &Res1{"xiaomi", "xiaomi3", "leijun"}
	r7 := &Res1{"xiaomi", "xiaomi2", "leijun"}
	assert.Equal(t, res[0].Name, r4.Name)
	assert.Equal(t, res[0].MobileName, r4.MobileName)
	assert.Equal(t, res[0].CeoName, r4.CeoName)
	assert.Equal(t, res[1].Name, r5.Name)
	assert.Equal(t, res[1].MobileName, r5.MobileName)
	assert.Equal(t, res[1].CeoName, r5.CeoName)
	assert.Equal(t, res[2].Name, r6.Name)
	assert.Equal(t, res[2].MobileName, r6.MobileName)
	assert.Equal(t, res[2].CeoName, r6.CeoName)
	assert.Equal(t, res[3].Name, r7.Name)
	assert.Equal(t, res[3].MobileName, r7.MobileName)
	assert.Equal(t, res[3].CeoName, r7.CeoName)
}

func TestModelWhereOrderByDescLimit(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"xiaomi\" " +
		"order by mobile.id desc, mobile.name limit 1,2"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2S", "ceo_name": "leijun"},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi3", "ceo_name": "leijun"},
	}) {
		t.Error(res)
	}
}

func TestModelWhereOrderByDescLimit2(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"xiaomi\" " +
		"order by mobile.id desc, mobile.name limit 1,2"
	golinq.RegisterType((*Res1)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res1").([]*Res1)
	r5 := &Res1{"xiaomi", "Mix2S", "leijun"}
	r6 := &Res1{"xiaomi", "xiaomi3", "leijun"}
	assert.Equal(t, res[0].Name, r5.Name)
	assert.Equal(t, res[0].MobileName, r5.MobileName)
	assert.Equal(t, res[0].CeoName, r5.CeoName)
	assert.Equal(t, res[1].Name, r6.Name)
	assert.Equal(t, res[1].MobileName, r6.MobileName)
	assert.Equal(t, res[1].CeoName, r6.CeoName)
}

func TestModelWhereOrderBy(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name, mobile_price.price " +
		"from company join mobile join ceo on company.id = mobile.company_id " +
		"join mobile_price on company.id = ceo.company_id " +
		"where mobile_price.mobile_id = mobile.id " +
		"order by mobile_price.price, mobile.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi2", "ceo_name": "leijun", "price": float32(1000)},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2", "ceo_name": "leijun", "price": float32(2000)},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "Mix2S", "ceo_name": "leijun", "price": float32(2000)},
		map[string]interface{}{"name": "xiaomi", "mobile_name": "xiaomi3", "ceo_name": "leijun", "price": float32(2000)},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone4", "ceo_name": "Tim Cook", "price": float32(5000)},
		map[string]interface{}{"name": "apple", "mobile_name": "iphone5", "ceo_name": "Tim Cook", "price": float32(6000)},
		map[string]interface{}{"name": "huawei", "mobile_name": "Mate20", "ceo_name": "renzhengfei", "price": float32(7000)},
	}) {
		t.Error(res)
	}
}

func TestModelOrderBy(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"order by company_id asc"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1},
		map[string]interface{}{"company_id": 1},
		map[string]interface{}{"company_id": 2},
		map[string]interface{}{"company_id": 3},
		map[string]interface{}{"company_id": 3},
		map[string]interface{}{"company_id": 4},
		map[string]interface{}{"company_id": 4},
	}) {
		t.Error(res)
	}
}

func TestModelOrderByDesc(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"order by company_id desc"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 4},
		map[string]interface{}{"company_id": 4},
		map[string]interface{}{"company_id": 3},
		map[string]interface{}{"company_id": 3},
		map[string]interface{}{"company_id": 2},
		map[string]interface{}{"company_id": 1},
		map[string]interface{}{"company_id": 1},
	}) {
		t.Error(res)
	}
}

func TestModelOrderByDesc2(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"order by company_id desc"
	golinq.RegisterType((*Res9)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res9").([]*Res9)
	r1 := &Res9{CompanyId: 4}
	r2 := &Res9{CompanyId: 4}
	r3 := &Res9{CompanyId: 3}
	r4 := &Res9{CompanyId: 3}
	r5 := &Res9{CompanyId: 2}
	r6 := &Res9{CompanyId: 1}
	r7 := &Res9{CompanyId: 1}
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
	assert.Equal(t, res[1].CompanyId, r2.CompanyId)
	assert.Equal(t, res[2].CompanyId, r3.CompanyId)
	assert.Equal(t, res[3].CompanyId, r4.CompanyId)
	assert.Equal(t, res[4].CompanyId, r5.CompanyId)
	assert.Equal(t, res[5].CompanyId, r6.CompanyId)
	assert.Equal(t, res[6].CompanyId, r7.CompanyId)
}

func TestModelGroupBy(t *testing.T) {
	sql := "select company.name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"group by company.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "apple"},
		map[string]interface{}{"name": "huawei"},
		map[string]interface{}{"name": "xiaomi"},
	}) {
		t.Error(res)
	}
}

func TestModelGroupBy2(t *testing.T) {
	sql := "select company.name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"group by company.name order by company.name"
	golinq.RegisterType((*Res1)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res1").([]*Res1)
	r1 := &Res1{Name: "apple"}
	r2 := &Res1{Name: "huawei"}
	r3 := &Res1{Name: "xiaomi"}
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[1].Name, r2.Name)
	assert.Equal(t, res[2].Name, r3.Name)
}

type Res10 struct {
	CompanyName     string `sql:"company_name"`
	MobileId        int    `sql:"mobile_id"`
	MobileName      string `sql:"mobile_name"`
	CompanyId       int    `sql:"company_id"`
	CeoId           int    `sql:"ceo_id"`
	CeoName         string `sql:"ceo_name"`
	CeoCompanyId    int    `sql:"ceo_company_id"`
	MobileCompanyId int    `sql:"mobile_company_id"`
}

func TestMoedlGroupByDoubleNames(t *testing.T) {
	sql := "select * " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"group by company.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
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

func TestMoedlGroupByDoubleNames2(t *testing.T) {
	sql := "select * " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"group by company.name order by company.name"
	golinq.RegisterType((*Res10)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res10").([]*Res10)
	r1 := &Res10{
		CompanyName:     "apple",
		MobileId:        0,
		MobileName:      "",
		CompanyId:       0,
		CeoId:           0,
		CeoName:         "",
		CeoCompanyId:    0,
		MobileCompanyId: 0,
	}
	r2 := &Res10{
		CompanyName:     "huawei",
		MobileId:        0,
		MobileName:      "",
		CompanyId:       0,
		CeoId:           0,
		CeoName:         "",
		CeoCompanyId:    0,
		MobileCompanyId: 0,
	}
	r3 := &Res10{
		CompanyName:     "xiaomi",
		MobileId:        0,
		MobileName:      "",
		CompanyId:       0,
		CeoId:           0,
		CeoName:         "",
		CeoCompanyId:    0,
		MobileCompanyId: 0,
	}
	assert.Equal(t, res[0].CompanyName, r1.CompanyName)
	assert.Equal(t, res[0].MobileId, r1.MobileId)
	assert.Equal(t, res[0].MobileName, r1.MobileName)
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
	assert.Equal(t, res[0].CeoName, r1.CeoName)
	assert.Equal(t, res[0].CeoId, r1.CeoId)
	assert.Equal(t, res[0].CeoCompanyId, r1.CeoCompanyId)
	assert.Equal(t, res[0].MobileCompanyId, r1.MobileCompanyId)
	assert.Equal(t, res[1].CompanyName, r2.CompanyName)
	assert.Equal(t, res[1].MobileId, r2.MobileId)
	assert.Equal(t, res[1].MobileName, r2.MobileName)
	assert.Equal(t, res[1].CompanyId, r2.CompanyId)
	assert.Equal(t, res[1].CeoName, r2.CeoName)
	assert.Equal(t, res[1].CeoId, r2.CeoId)
	assert.Equal(t, res[1].CeoCompanyId, r2.CeoCompanyId)
	assert.Equal(t, res[1].MobileCompanyId, r2.MobileCompanyId)
	assert.Equal(t, res[2].CompanyName, r3.CompanyName)
	assert.Equal(t, res[2].MobileId, r3.MobileId)
	assert.Equal(t, res[2].MobileName, r3.MobileName)
	assert.Equal(t, res[2].CompanyId, r3.CompanyId)
	assert.Equal(t, res[2].CeoName, r3.CeoName)
	assert.Equal(t, res[2].CeoId, r3.CeoId)
	assert.Equal(t, res[2].CeoCompanyId, r3.CeoCompanyId)
	assert.Equal(t, res[2].MobileCompanyId, r3.MobileCompanyId)
}

type Res11 struct {
	CompanyId int     `sql:"company_id"`
	Count     float64 `sql:"count1"`
}

func TestModelGroupByFunc(t *testing.T) {
	sql := "select company_id, gcount(*) as count1 " +
		"from mobile " +
		"group by company_id"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"company_id"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1, "count1": 2},
		map[string]interface{}{"company_id": 2, "count1": 1},
		map[string]interface{}{"company_id": 3, "count1": 2},
		map[string]interface{}{"company_id": 4, "count1": 2},
	}) {
		t.Error(res)
	}
}

func TestModelGroupByFunc2(t *testing.T) {
	sql := "select company_id, gcount(*) as count1 " +
		"from mobile " +
		"group by company_id order by company_id"
	golinq.RegisterType((*Res11)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res11").([]*Res11)
	r1 := &Res11{1, 2}
	r2 := &Res11{2, 1}
	r3 := &Res11{3, 2}
	r4 := &Res11{4, 2}
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
	assert.Equal(t, res[0].Count, r1.Count)
	assert.Equal(t, res[1].CompanyId, r2.CompanyId)
	assert.Equal(t, res[1].Count, r2.Count)
	assert.Equal(t, res[2].CompanyId, r3.CompanyId)
	assert.Equal(t, res[2].Count, r3.Count)
	assert.Equal(t, res[3].CompanyId, r4.CompanyId)
	assert.Equal(t, res[3].Count, r4.Count)
}

func TestModelGroupByHavingFunc(t *testing.T) {
	sql := "select company_id, gcount(*) as count1 " +
		"from mobile group by company_id " +
		"having count(*) >= 2 and company_id != 4"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"company_id"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1, "count1": 2},
		map[string]interface{}{"company_id": 3, "count1": 2},
	}) {
		t.Error(res)
	}
}

func TestModelGroupByNestingFunc(t *testing.T) {
	sql := "select gsum(mobile.id) as sum1, company_id " +
		"from mobile " +
		"group by company_id"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"company_id"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1, "sum1": 9},
		map[string]interface{}{"company_id": 2, "sum1": 6},
		map[string]interface{}{"company_id": 3, "sum1": 3},
		map[string]interface{}{"company_id": 4, "sum1": 6},
	}) {
		t.Error(res)
	}
}

func TestModelGroupByHavingOrderByDesc(t *testing.T) {
	sql := "select company.name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"group by company.name " +
		"having len(company.name) = 6 " +
		"order by company.name desc"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "xiaomi"},
		map[string]interface{}{"name": "huawei"},
	}) {
		t.Error(res)
	}
}

func TestModelGroupByHavingOrderByDesc2(t *testing.T) {
	sql := "select company.name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"group by company.name " +
		"having len(company.name) = 6 " +
		"order by company.name desc"
	golinq.RegisterType((*Res1)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res1").([]*Res1)
	r1 := &Res1{Name: "xiaomi"}
	r2 := &Res1{Name: "huawei"}
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[1].Name, r2.Name)
}

func TestModelGroupByOrderBy(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"group by company_id " +
		"order by company_id"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 1},
		map[string]interface{}{"company_id": 2},
		map[string]interface{}{"company_id": 3},
		map[string]interface{}{"company_id": 4},
	}) {
		t.Error(res)
	}
}

func TestModelGroupByOrderByLimit(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"group by company_id " +
		"order by company_id limit 3,2"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"company_id": 4},
	}) {
		t.Error(res)
	}
}

func TestModelGroupByOrderByLimit2(t *testing.T) {
	sql := "select company_id " +
		"from mobile " +
		"group by company_id " +
		"order by company_id limit 3,2"
	golinq.RegisterType((*Res2)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res2").([]*Res2)
	r1 := &Res2{CompanyId: 4}
	assert.Equal(t, res[0].CompanyId, r1.CompanyId)
}

func TestModelGroupByWhereUnicodeCharAsFunc(t *testing.T) {
	sql := "select ceo.name as ceo_name, deflate(mobile_price.type, 0, mobile_price.price) as `普通`,  " +
		"deflate(mobile_price.type, 1, mobile_price.price) as `旗舰` " +
		"from mobile, company, ceo, mobile_price " +
		"where mobile.id = mobile_price.mobile_id " +
		"and mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"group by ceo.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	sorted(&res, []string{"ceo_name"})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"旗舰": float32(11000), "普通": 0, "ceo_name": "Tim Cook"},
		map[string]interface{}{"旗舰": float32(4000), "普通": float32(3000), "ceo_name": "leijun"},
		map[string]interface{}{"旗舰": float32(7000), "普通": 0, "ceo_name": "renzhengfei"},
	}) {
		t.Error(res)
	}
}

func TestModelBlankResult(t *testing.T) {
	sql := "select ceo.name as ceo_name, deflate(mobile_price.type, 0, mobile_price.price) as `普通`,  " +
		"deflate(mobile_price.type, 1, mobile_price.price) as `旗舰` " +
		"from mobile, company, ceo, mobile_price " +
		"where mobile.id = mobile_price.mobile_id " +
		"and mobile.company_id = company.id " +
		"and company.id = ceo.company_id " +
		"group by ceo.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCeos(), getMobiles(), getMobilePrices()), "").([]interface{})
	if !isEqual(res, []interface{}{}) {
		t.Error(res)
	}
}

func TestModelMulDotOperator(t *testing.T) {
	sql := "select m.mobile.company.ceo as ceo_name from mobile_price as m " +
		"join ceo where m.mobile.company.id = ceo.company_id " +
		"group by ceo.name having ceo.name != \"renzhengfei\" order by ceo.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getMobilePrices2(), getCeos()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"ceo_name": "Tim Cook"},
		map[string]interface{}{"ceo_name": "leijun"},
	}) {
		t.Error(res)
	}
}

func TestModelMulDotOperator2(t *testing.T) {
	sql := "select m.mobile.company.ceo as ceo_name from mobile_price as m " +
		"join ceo where m.mobile.company.id = ceo.company_id " +
		"group by ceo.name having ceo.name != \"renzhengfei\" order by ceo.name"
	golinq.RegisterType((*Res1)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getMobilePrices2(), getCeos()), "Res1").([]*Res1)
	r1 := &Res1{CeoName: "Tim Cook"}
	r2 := &Res1{CeoName: "leijun"}
	assert.Equal(t, res[0].CeoName, r1.CeoName)
	assert.Equal(t, res[1].CeoName, r2.CeoName)
}

func TestModelDictFieldOrderMulDot(t *testing.T) {
	sql := "select m.mobile.company.name from mobile_price as m " +
		"join ceo where m.mobile.company.id = ceo.company_id " +
		"and m.id not in (1, 2, 4) " +
		"group by ceo.name having ceo.name != \"renzhengfei\" order by m.mobile.company.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getMobilePrices2(), getCeos()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "apple"},
		map[string]interface{}{"name": "xiaomi"},
	}) {
		t.Error(res)
	}
}

func TestModelDictFieldOrderMulDot2(t *testing.T) {
	sql := "select m.mobile.company.name from mobile_price as m " +
		"join ceo where m.mobile.company.id = ceo.company_id " +
		"and m.id not in (1, 2, 4) " +
		"group by ceo.name having ceo.name != \"renzhengfei\" order by m.mobile.company.name"
	golinq.RegisterType((*Res1)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getMobilePrices2(), getCeos()), "Res1").([]*Res1)
	r1 := &Res1{Name: "apple"}
	r2 := &Res1{Name: "xiaomi"}
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[1].Name, r2.Name)
}

func TestModelGroupByHavingMulDot(t *testing.T) {
	sql := "select m.mobile.company.name from mobile_price as m " +
		"join ceo where m.mobile.company.id = ceo.company_id " +
		"group by m.mobile.company.name having m.mobile.company.name != \"apple\" order by m.mobile.company.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getMobilePrices2(), getCeos()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": "huawei"},
		map[string]interface{}{"name": "xiaomi"},
	}) {
		t.Error(res)
	}
}

func TestModelGroupByHavingMulDot2(t *testing.T) {
	sql := "select m.mobile.company.name from mobile_price as m " +
		"join ceo where m.mobile.company.id = ceo.company_id " +
		"group by m.mobile.company.name having m.mobile.company.name != \"apple\" order by m.mobile.company.name"
	golinq.RegisterType((*Res1)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getMobilePrices2(), getCeos()), "Res1").([]*Res1)
	r1 := &Res1{Name: "huawei"}
	r2 := &Res1{Name: "xiaomi"}
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[1].Name, r2.Name)
}

func TestModelListField(t *testing.T) {
	sql := "select t.name, s.number from teacher as t, student as s where t.id = s.id " +
		"group by t.name, student.number order by t.name"
	res := golinq.SqlRun(sql, golinq.NewDataSources(getTeachers(), getStudents()), "").([]interface{})
	if !isEqual(res, []interface{}{
		map[string]interface{}{"name": []string{"jay", "tom", "will"}, "number": 10},
		map[string]interface{}{"name": []string{"zhang", "li", "wang"}, "number": 10},
	}) {
		t.Error(res)
	}
}

type Res12 struct {
	Name   []string `sql:"name"`
	Number int      `sql:"number"`
}

func TestModelListField2(t *testing.T) {
	sql := "select t.name, s.number from teacher as t, student as s where t.id = s.id " +
		"group by t.name, student.number order by t.name"
	golinq.RegisterType((*Res12)(nil))
	res := golinq.SqlRun(sql, golinq.NewDataSources(getTeachers(), getStudents()), "Res12").([]*Res12)
	r1 := &Res12{Name: []string{"jay", "tom", "will"}, Number: 10}
	r2 := &Res12{Name: []string{"zhang", "li", "wang"}, Number: 10}
	assert.Equal(t, res[0].Name, r1.Name)
	assert.Equal(t, res[0].Number, r1.Number)
	assert.Equal(t, res[1].Name, r2.Name)
	assert.Equal(t, res[1].Number, r2.Number)
}

var (
	companys      []*Company
	mobiles       []*Mobile
	ceos          []*Ceo
	mobilePrices  []*MobilePrice
	once4         sync.Once
	once5         sync.Once
	once6         sync.Once
	once7         sync.Once
	mobilePrices2 []*MobilePrice2
	once8         sync.Once
)

type Company struct {
	Id   int    `sql:"id"`
	Name string `sql:"name"`
}

func (Company) TableName() string {
	return "company"
}

type Mobile struct {
	Id        int    `sql:"id"`
	Name      string `sql:"name"`
	CompanyId int    `sql:"company_id"`
}

func (Mobile) TableName() string {
	return "mobile"
}

type Ceo struct {
	Id        int    `sql:"id"`
	Name      string `sql:"name"`
	CompanyId int    `sql:"company_id"`
}

func (Ceo) TableName() string {
	return "ceo"
}

type MobilePrice struct {
	Id       int     `sql:"id"`
	MobileId int     `sql:"mobile_id"`
	Type     int     `sql:"type"`
	Price    float32 `sql:"price"`
}

func (MobilePrice) TableName() string {
	return "mobile_price"
}

func getCompanys() []*Company {
	once4.Do(func() {
		c1 := &Company{
			Id:   1,
			Name: "apple",
		}
		c2 := &Company{
			Id:   2,
			Name: "huawei",
		}
		c3 := &Company{
			Id:   3,
			Name: "xiaomi",
		}
		c4 := &Company{
			Id:   4,
			Name: "xiaomi",
		}
		companys = []*Company{c1, c2, c3, c4}
	})
	return companys
}

func getMobiles() []*Mobile {
	once5.Do(func() {
		m1 := &Mobile{
			Id:        1,
			Name:      "xiaomi2",
			CompanyId: 3,
		}
		m2 := &Mobile{
			Id:        2,
			Name:      "xiaomi3",
			CompanyId: 3,
		}
		m3 := &Mobile{
			Id:        3,
			Name:      "Mix2",
			CompanyId: 4,
		}
		m4 := &Mobile{
			Id:        3,
			Name:      "Mix2S",
			CompanyId: 4,
		}
		m5 := &Mobile{
			Id:        4,
			Name:      "iphone4",
			CompanyId: 1,
		}
		m6 := &Mobile{
			Id:        5,
			Name:      "iphone5",
			CompanyId: 1,
		}
		m7 := &Mobile{
			Id:        6,
			Name:      "Mate20",
			CompanyId: 2,
		}
		mobiles = []*Mobile{m1, m2, m3, m4, m5, m6, m7}
	})
	return mobiles
}

func getCeos() []*Ceo {
	once6.Do(func() {
		c1 := &Ceo{
			Id:        1,
			Name:      "Tim Cook",
			CompanyId: 1,
		}
		c2 := &Ceo{
			Id:        2,
			Name:      "renzhengfei",
			CompanyId: 2,
		}
		c3 := &Ceo{
			Id:        3,
			Name:      "leijun",
			CompanyId: 3,
		}
		c4 := &Ceo{
			Id:        4,
			Name:      "leijun",
			CompanyId: 4,
		}
		ceos = []*Ceo{c1, c2, c3, c4}
	})
	return ceos
}

func getMobilePrices() []*MobilePrice {
	once7.Do(func() {
		m1 := &MobilePrice{
			Id:       1,
			MobileId: 1,
			Type:     0,
			Price:    1000,
		}
		m2 := &MobilePrice{
			Id:       2,
			MobileId: 2,
			Type:     0,
			Price:    2000,
		}
		m3 := &MobilePrice{
			Id:       3,
			MobileId: 3,
			Type:     1,
			Price:    2000,
		}
		m4 := &MobilePrice{
			Id:       4,
			MobileId: 4,
			Type:     1,
			Price:    5000,
		}
		m5 := &MobilePrice{
			Id:       5,
			MobileId: 5,
			Type:     1,
			Price:    6000,
		}
		m6 := &MobilePrice{
			Id:       6,
			MobileId: 6,
			Type:     1,
			Price:    7000,
		}
		mobilePrices = []*MobilePrice{m1, m2, m3, m4, m5, m6}
	})
	return mobilePrices
}

type MobilePrice2 struct {
	Id     int      `sql:"id"`
	Mobile *Mobile2 `sql:"mobile"`
	Type   int      `sql:"type"`
	Price  float32  `sql:"price"`
}

func (MobilePrice2) TableName() string {
	return "mobile_price"
}

type Mobile2 struct {
	Id      int       `sql:"id"`
	Name    string    `sql:"name"`
	Company *Company2 `sql:"company"`
}

func (Mobile2) TableName() string {
	return "mobile"
}

type Company2 struct {
	Id   int    `sql:"id"`
	Name string `sql:"name"`
	Ceo  string `sql:"ceo"`
}

func (Company2) TableName() string {
	return "company"
}

func getMobilePrices2() []*MobilePrice2 {
	once8.Do(func() {
		c1 := &Company2{
			Id:   3,
			Name: "xiaomi",
			Ceo:  "leijun",
		}
		c2 := &Company2{
			Id:   4,
			Name: "xiaomi",
			Ceo:  "leijun",
		}
		c3 := &Company2{
			Id:   1,
			Name: "apple",
			Ceo:  "Tim Cook",
		}
		c4 := &Company2{
			Id:   2,
			Name: "huawei",
			Ceo:  "renzhengfei",
		}
		m1 := &Mobile2{
			Id:      1,
			Name:    "xiaomi2",
			Company: c1,
		}
		m2 := &Mobile2{
			Id:      2,
			Name:    "xiaomi3",
			Company: c1,
		}
		m3 := &Mobile2{
			Id:      3,
			Name:    "Mix2",
			Company: c2,
		}
		m4 := &Mobile2{
			Id:      4,
			Name:    "iphone4",
			Company: c3,
		}
		m5 := &Mobile2{
			Id:      5,
			Name:    "iphone5",
			Company: c3,
		}
		m6 := &Mobile2{
			Id:      6,
			Name:    "Mate20",
			Company: c4,
		}
		p1 := &MobilePrice2{
			Id:     1,
			Mobile: m1,
			Type:   0,
			Price:  1000,
		}
		p2 := &MobilePrice2{
			Id:     2,
			Mobile: m2,
			Type:   0,
			Price:  2000,
		}
		p3 := &MobilePrice2{
			Id:     3,
			Mobile: m3,
			Type:   1,
			Price:  2000,
		}
		p4 := &MobilePrice2{
			Id:     4,
			Mobile: m4,
			Type:   1,
			Price:  5000,
		}
		p5 := &MobilePrice2{
			Id:     5,
			Mobile: m5,
			Type:   1,
			Price:  6000,
		}
		p6 := &MobilePrice2{
			Id:     6,
			Mobile: m6,
			Type:   1,
			Price:  7000,
		}
		mobilePrices2 = []*MobilePrice2{p1, p2, p3, p4, p5, p6}
	})
	return mobilePrices2
}

type Teacher struct {
	Name []string `sql:"name"`
	Year int      `sql:"year"`
	Id   int      `sql:"id"`
}

func (Teacher) TableName() string {
	return "teacher"
}

type Student struct {
	Number int `sql:"number"`
	Age    int `sql:"age"`
	Id     int `sql:"id"`
}

func (Student) TableName() string {
	return "student"
}

func getTeachers() []*Teacher {
	t1 := &Teacher{
		Name: []string{"zhang", "li", "wang"},
		Year: 5,
		Id:   1,
	}
	t2 := &Teacher{
		Name: []string{"jay", "tom", "will"},
		Year: 3,
		Id:   2,
	}
	return []*Teacher{t1, t2}
}

func getStudents() []*Student {
	s1 := &Student{
		Number: 10,
		Age:    18,
		Id:     1,
	}
	s2 := &Student{
		Number: 10,
		Age:    19,
		Id:     2,
	}
	s3 := &Student{
		Number: 10,
		Age:    20,
		Id:     3,
	}
	s4 := &Student{
		Number: 20,
		Age:    21,
		Id:     4,
	}
	return []*Student{s1, s2, s3, s4}
}
