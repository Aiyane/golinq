## GoLinq

### 介绍

一个执行 SQL 语句的 Go 工具。内置 sql 驱动，可作为 gorm 的 mock 工具。支持 model 中嵌套 struct，默认驼峰命名转蛇形命名，兼容`ID`, 支持 Time 类型，支持自定义类型(需要实现 sql.Scanner 接口与 driver.Valuer 接口)

将数据库以 Go 中 map 格式保存在内存，只需 mock 数据库数据，sql 语句执行结果与数据库执行基本一致。

### gorm 使用示例
```go
var DB *gorm.DB

func TestLinq(t *testing.T) {
	Catcher.Register()
	createDataBase()
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
}

func createDataBase() {
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
		}`), &DataBase)
	if err != nil {
		panic(err)
	}
}

type Company struct {
	Id   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (Company) TableName() string {
	return "company"
}
```

### 更多示例

请看 `linq_test.go` 文件 与 `tests` 文件夹的示例。

### 更多特性

#### 使用索引

在 join 多张表时，相当于做多重循环。如果 join 的表数量很多，耗时会过大。如果对表的某个字段建立索引，则只进行一次循环。

对 JOIN ON 语法建立索引，在 ON 条件中是使用先建索引后筛选的方式，所以在 JOIN 多张表时，请使用 JOIN ON 语句，示例：
```go
sql := "SELECT company.name, mobile.name AS mobile_name, ceo.name AS ceo_name, mobile_price.price " +
       "FROM company JOIN mobile JOIN ceo ON company.id = mobile.company_id " +
       "JOIN mobile_price ON company.id = ceo.company_id " +
       "WHERE mobile_price.mobile_id = mobile.id " +
       "ORDER BY mobile_price.price"
```

#### 多点操作

在 map 结构中，可能会存在嵌套的字典，可以在 **SELECT 语句、GROUP BY 语句 HAVING 语句、WHERE 语句与 ORDER BY 语句**中使用多点取出嵌套的字段。例如

```go
func main() {
    data := make(types.DataSources)
    err := json.Unmarshal([]byte(`{
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
    }`), &data)
    sql := "select m.mobile.company.ceo as ceo_name from mobile_price as m " +
           "join ceo where m.mobile.company.id = ceo.company_id " +
           "group by ceo.name having ceo.name != \"renzhengfei\" order by ceo.name"
    res := golinq.SqlRun(sql, data, "")
    fmt.Println(res)
}
```
结果为
```go
[map[ceo_name:Tim Cook] map[ceo_name:leijun]]
```

#### 结构体支持

可以在使用结构体列表的字典类比数据库，当然查询结果也可以包装在结构体中。所以有四种场景都提供支持

- 入参是结构体列表字典，出参是结构体列表
- 入参是结构体列表字典，出参是简单map
- 入参是简单map，出参是简单map (例如上述列子)
- 入参是简单map，出参是结构体列表

这里仅写出第一种情况示例
```go
type Res1 struct {
	Name       string `sql:"name"`
	MobileName string `sql:"mobile_name"`
	CeoName    string `sql:"ceo_name"`
}

func TestModelJoinOnWhere2(t *testing.T) {
	sql := "select company.name, mobile.name as mobile_name, ceo.name as ceo_name " +
		"from company join mobile on company.id = mobile.company_id join ceo on company.id = ceo.company_id " +
		"where company.name = \"apple\" order by mobile.name"
    // 此处注册查询结果结构体
	golinq.RegisterType((*Res1)(nil))
    // 第三个参数是结果体名称
    // golinq.NewDataSources 方法会自动组装结构体列表成 interpreters.DataSources 类型，可接受多个参数
	res := golinq.SqlRun(sql, golinq.NewDataSources(getCompanys(), getCeos(), getMobiles(), getMobilePrices()), "Res1")
}
```
结构体定义如下
```go
// 结构体定义, 必须实现 TableName 方法定义表名, 字段可以使用 `sql:"xx"` 注释改名
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

// 构造结构体列表
func getCompanys() []*Company {
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
	return companys
}

func getMobiles() []*Mobile {
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
	return mobiles
}

func getCeos() []*Ceo {
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
	return ceos
}
```
