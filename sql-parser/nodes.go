package sql_parser

import "github.com/golang-collections/collections/queue"

type OuterTag int
type BuiltinType int

func (BuiltinType) GetTokenName() string {
	return "builtin"
}

const (
	LeftOn     OuterTag = 1
	LeftUsing  OuterTag = 2
	RightOn    OuterTag = 3
	RightUsing OuterTag = 4
)

const (
	Default BuiltinType = -1
)

type SelectStatement struct {
	Id   int
	Tree *State
}

// Token 类型定义
type SqlToken interface {
	GetTokenName() string
}

type Tables struct {
	Values []SqlToken // As 或 Var
}

func (self *Tables) GetTokenName() string {
	return "tables"
}

type As struct {
	NewName *Const
	Value   SqlToken
}

func (self *As) GetTokenName() string {
	return "as"
}

type Func struct {
	Name string
	Args []interface{}
}

func (self *Func) GetTokenName() string {
	return "func"
}

type Case struct {
	Values []*CaseItem
}

func (self *Case) GetTokenName() string {
	return "case"
}

type CaseItem struct {
	Condition SqlToken  // *Func 或 *Var
	Value     SqlToken
}

func (self *CaseItem) GetTokenName() string {
	return "case-item"
}

type Inner struct {
	Table     *Var
	Condition *Func
}

func (self *Inner) GetTokenName() string {
	return "inner"
}

type Outer struct {
	Tag       OuterTag
	Table     SqlToken // As 或 Var
	Condition *Func
	Using     []SqlToken
}

func (self *Outer) GetTokenName() string {
	return "outer"
}

type Desc struct {
	Value *Var
}

func (self *Desc) GetTokenName() string {
	return "desc"
}

type Limit struct {
	Limit  *Const
	Offset *Const
}

func (self *Limit) GetTokenName() string {
	return "limit"
}

type Link struct {
	Id int
}

func (self *Link) GetTokenName() string {
	return "link"
}

type From struct {
	Tables *Tables
	Where  *Func
	Group  []*Var
	Having *Func
}

func (self *From) GetTokenName() string {
	return "from"
}

type State struct {
	From   *From
	Order  []SqlToken
	Limit  *Limit
	Select []SqlToken
}

func (self *State) GetTokenName() string {
	return "state"
}

type Var struct {
	Value *Const
	Next  *Var
}

func (self *Var) GetTokenName() string {
	return "var"
}

type Const struct {
	Value interface{}
}

func (self *Const) GetTokenName() string {
	return "const"
}

type EXPR struct {
	FromExpr   *Tables
	WhereExpr  *Func
	GroupExpr  []*Var
	HavingExpr *Func
	OrderExpr  []SqlToken
	LimitExpr  *Limit
	SelectExpr []SqlToken
}

func NewExpr(tree *State) *EXPR {
	frontPart := tree.From
	return &EXPR{
		FromExpr:   frontPart.Tables,
		WhereExpr:  frontPart.Where,
		GroupExpr:  frontPart.Group,
		HavingExpr: frontPart.Having,
		OrderExpr:  tree.Order,
		LimitExpr:  tree.Limit,
		SelectExpr: tree.Select,
	}
}

type InsertExpr struct {
	Ignore  bool
	Table   string
	Columns []SqlToken      // 可以是空
	Link    *Link           // 可以是空
	Values  [][]SqlToken // 可以是空
}

type ColumnValue struct {
	Column *Var
	Value  SqlToken
}

type UpdateExpr struct {
	Ignore bool
	Table  SqlToken
	Values []*ColumnValue
	Where  *Func
	Order  []SqlToken
	Limit  *Limit
}

type DeleteExpr struct {
	Ignore bool
	Table  string
	Where  *Func
	Order  []SqlToken
	Limit  *Limit
}

type Tree struct {
	SelectStatementsQueue *queue.Queue // 存 SQL 语句的语法树队列
	UpdateStatement       *UpdateExpr
	InsertStatement       *InsertExpr
	DeleteStatement       *DeleteExpr
}
