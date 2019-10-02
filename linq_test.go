package golinq

import (
	"encoding/json"
	"github.com/Aiyane/golinq/types"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

var DB *gorm.DB

func TestLinq(t *testing.T) {
	// 注册 sql 驱动
	Catcher.Register()
	// 创建数据库对象
	createDataBase()
	// 获取 db 对象
	db, _ := gorm.Open(DriverName, "test.db")
	DB = db
	RegisterDB2DataBase(map[string]types.DataSources{
		"test.db": DataBase,
	})

	// 测试普通 select 语句
	var company Company
	DB.First(&company, "id=?", 2)
	assert.Equal(t, company.Id, 2)
	assert.Equal(t, company.Name, "huawei")

	// 测试普通 insert 语句
	company = Company{
		Name: "oppo",
	}
	DB.Create(&company)
	var company2 Company
	DB.First(&company2, "id=?", 5)
	assert.Equal(t, company2.Name, "oppo")
	assert.Equal(t, company2.Id, 5)

	// 测试批量 insert 语句
	companies := []*Company{
		{
			Name: "vivo",
		},
		{
			Name: "hello",
		},
	}
	err := company.BatchCreate(DB, companies)
	if err != nil {
		t.Error(err)
	}
	var companies2 []*Company
	DB.Where("id>?", 4).Order("id").Find(&companies2)
	assert.Equal(t, len(companies2), 3)
	assert.Equal(t, companies2[0].Id, 5)
	assert.Equal(t, companies2[0].Name, "oppo")
	assert.Equal(t, companies2[1].Id, 6)
	assert.Equal(t, companies2[1].Name, "vivo")
	assert.Equal(t, companies2[2].Id, 7)
	assert.Equal(t, companies2[2].Name, "hello")

	// 测试普通 update 语句
	DB.Table("company").Where("id > ?", 5).Update(map[string]interface{}{"name": "haha"})
	var companies3 []*Company
	DB.Where("id>?", 5).Order("id").Find(&companies3)
	assert.Equal(t, len(companies3), 2)
	assert.Equal(t, companies3[0].Id, 6)
	assert.Equal(t, companies3[0].Name, "haha")
	assert.Equal(t, companies3[1].Id, 7)
	assert.Equal(t, companies3[1].Name, "haha")

	// 测试批量 update 语句
	companies = []*Company{
		{
			Id:   6,
			Name: "vivo",
		},
		{
			Id:   7,
			Name: "hello",
		},
	}
	err = company.BatchUpdate(DB, []string{"name"}, companies)
	if err != nil {
		t.Error(err)
	}
	var companies5 []*Company
	DB.Where("id>?", 5).Order("id").Find(&companies5)
	assert.Equal(t, len(companies5), 2)
	assert.Equal(t, companies5[0].Id, 6)
	assert.Equal(t, companies5[0].Name, "vivo")
	assert.Equal(t, companies5[1].Id, 7)
	assert.Equal(t, companies5[1].Name, "hello")

	DB.Table("company").Where("id > ?", 5).Update(map[string]interface{}{"name": "haha"})
	var companies6 []*Company
	DB.Where("id>?", 5).Order("id").Find(&companies6)
	assert.Equal(t, len(companies6), 2)
	assert.Equal(t, companies6[0].Id, 6)
	assert.Equal(t, companies6[0].Name, "haha")
	assert.Equal(t, companies6[1].Id, 7)
	assert.Equal(t, companies6[1].Name, "haha")

	// 测试普通 delete 语句
	DB.Where("name=?", "haha").Delete(Company{})
	var companies4 []*Company
	DB.Where("id>?", 5).Order("id").Find(&companies4)
	assert.Equal(t, len(companies4), 0)

	var mobiles []*Mobile
	DB.Where("company_id=? and name=?", 4, "Mix2S").Find(&mobiles)
	assert.Equal(t, len(mobiles), 1)
	assert.Equal(t, mobiles[0].Id, 3)
	assert.Equal(t, mobiles[0].Name, "Mix2S")
	assert.Equal(t, mobiles[0].CompanyId, 4)

	// 测试 order by, limit 语句
	var mobilePrice []*MobilePrice
	DB.Select("mobile_id, price").Where("id > ?", 1).Order("id desc").Limit(
		2).Find(&mobilePrice)
	assert.Equal(t, len(mobilePrice), 2)
	assert.Equal(t, mobilePrice[0].MobileId, 6)
	assert.Equal(t, mobilePrice[0].Price, float32(7000))
	assert.Equal(t, mobilePrice[1].MobileId, 5)
	assert.Equal(t, mobilePrice[1].Price, float32(6000))
}

func createDataBase() {
	// 可以使用json对象表示数据库对象
	// 或者 `string : struct列表` 的字典表示数据库对象
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
		}`), &DataBase)
	if err != nil {
		panic(err)
	}
}
