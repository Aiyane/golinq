package sql_parser

import (
	"github.com/Aiyane/golinq/gen"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang-collections/collections/queue"
)

// 新建对象
func NewVisitor() *visitor {
	return &visitor{
		SelectStatementsQueue: queue.New(),
		_id:                   0,
	}
}

type visitor struct {
	*parser.BaseMySqlParserVisitor
	SelectStatementsQueue *queue.Queue // 存 SQL 语句的语法树队列
	UpdateStatement       *UpdateExpr
	InsertStatement       *InsertExpr
	DeleteStatement       *DeleteExpr
	_id                   int // 每一个语句的 id
}

// 获取 select 语句 id
func (self *visitor) getId() int {
	self._id++
	return self._id
}

func (self *visitor) GetTree() *Tree {
	return &Tree{
		SelectStatementsQueue: self.SelectStatementsQueue,
		UpdateStatement:       self.UpdateStatement,
		InsertStatement:       self.InsertStatement,
		DeleteStatement:       self.DeleteStatement,
	}
}

func (self *visitor) Visit(root *parser.DmlStatementContext) {
	selectCtx := root.QuerySpecification()
	if selectCtx != nil {
		self.doVisitQuerySpecification(selectCtx.(*parser.QuerySpecificationContext))
		return
	}
	insertCtx := root.InsertStatement()
	if insertCtx != nil {
		self.doVisitInsertStatement(insertCtx.(*parser.InsertStatementContext))
		return
	}
	updateCtx := root.UpdateStatement()
	if updateCtx != nil {
		self.doVisitUpdateStatement(updateCtx.(*parser.UpdateStatementContext))
		return
	}
	deleteCtx := root.DeleteStatement()
	if deleteCtx != nil {
		self.doVisitDeleteStatement(deleteCtx.(*parser.DeleteStatementContext))
	}
}

func (self *visitor) doVisitInsertStatement(ctx *parser.InsertStatementContext) {
	/*
		INSERT
		IGNORE? INTO? ID
		  ('(' columns=uidList ')')? insertStatementValue
	*/
	ignore := false
	if ctx.IGNORE() != nil {
		ignore = true
	}
	table := ctx.ID().GetText()
	var uidList []SqlToken
	if columns := ctx.UidList(); columns != nil {
		uidList = self.doVisitUidList(columns.(*parser.UidListContext))
	}
	insertStatementValue := self.doVisitInsertStatementValue(ctx.InsertStatementValue().(*parser.InsertStatementValueContext))
	var values [][]SqlToken
	var link *Link
	switch v := insertStatementValue.(type) {
	case [][]SqlToken:
		values = v
	case *Link:
		link = v
	default:
		logrus.Errorf("[doVisitInsertStatement] insertStatementValue no such type [%v]", insertStatementValue)
		return
	}
	self.InsertStatement = &InsertExpr{
		Ignore:  ignore,
		Table:   doVisitID(table),
		Columns: uidList,
		Link:    link,
		Values:  values,
	}
}

func (self *visitor) doVisitInsertStatementValue(ctx *parser.InsertStatementValueContext) interface{} {
	/*
	   selectStatement
	    | insertFormat=(VALUES | VALUE)
	      '(' expressionsWithDefaults ')'
	        (',' '(' expressionsWithDefaults ')')*
	*/
	if selectStatement := ctx.SelectStatement(); selectStatement != nil {
		return self.visitSelectStatement(selectStatement)
	}
	var ret [][]SqlToken
	for _, expressionsWithDefaults := range ctx.AllExpressionsWithDefaults() {
		ret = append(ret, self.doVisitExpressionsWithDefaults(expressionsWithDefaults.(*parser.ExpressionsWithDefaultsContext)))
	}
	return ret
}

func (self *visitor) doVisitUpdatedElement(ctx *parser.UpdatedElementContext) *ColumnValue {
	/*
		fullColumnName '=' expressionOrDefault
	*/
	fullColumnName := self.doVisitFullColumnName(ctx.FullColumnName().(*parser.FullColumnNameContext))
	expressionOrDefault := self.doVisitExpressionOrDefault(ctx.ExpressionOrDefault().(*parser.ExpressionOrDefaultContext))
	return &ColumnValue{
		Column: fullColumnName,
		Value:  expressionOrDefault,
	}
}

func (self *visitor) doVisitExpressionsWithDefaults(ctx *parser.ExpressionsWithDefaultsContext) []SqlToken {
	/*
		expressionOrDefault (',' expressionOrDefault)*
	*/
	var ret []SqlToken
	for _, expressionOrDefault := range ctx.AllExpressionOrDefault() {
		ret = append(ret, self.doVisitExpressionOrDefault(expressionOrDefault.(*parser.ExpressionOrDefaultContext)))
	}
	return ret
}

func (self *visitor) doVisitExpressionOrDefault(ctx *parser.ExpressionOrDefaultContext) SqlToken {
	/*
		expression | DEFAULT
	*/
	if expression := ctx.Expression(); expression != nil {
		return self.visitExpression(expression).(SqlToken)
	}
	return Default
}

func (self *visitor) doVisitDeleteStatement(ctx *parser.DeleteStatementContext) {
	/*
		DELETE IGNORE?
		FROM ID
		  (WHERE expression)?
		  orderByClause? limitClause?
	*/
	ignore := false
	if ctx.IGNORE() != nil {
		ignore = true
	}
	table := ctx.ID().GetText()
	var where *Func
	if expression := ctx.Expression(); expression != nil {
		_where := self.visitExpression(expression)
		where = dealWhere(_where)
	}
	var order []SqlToken
	if orderByClause := ctx.OrderByClause(); orderByClause != nil {
		order = self.doVisitOrderByClause(orderByClause.(*parser.OrderByClauseContext))
	}
	var limit *Limit
	if limitClause := ctx.LimitClause(); limitClause != nil {
		limit = self.doVisitLimitClause(limitClause.(*parser.LimitClauseContext))
	}
	self.DeleteStatement = &DeleteExpr{
		Ignore: ignore,
		Table:  doVisitID(table),
		Where:  where,
		Order:  order,
		Limit:  limit,
	}
}

func (self *visitor) doVisitUpdateStatement(ctx *parser.UpdateStatementContext) {
	/*
		UPDATE IGNORE? ID (AS? ID)?
		  SET updatedElement (',' updatedElement)*
		  (WHERE expression)? orderByClause? limitClause?
	*/
	ignore := false
	if ctx.IGNORE() != nil {
		ignore = true
	}
	_table := &Var{
		Value: &Const{
			Value: doVisitID(ctx.ID(0).GetText()),
		},
		Next: nil,
	}
	var table SqlToken
	if as := ctx.AS(); as != nil {
		table = &As{
			NewName: &Const{
				Value: doVisitID(ctx.ID(1).GetText()),
			},
			Value: _table,
		}
	} else {
		table = _table
	}
	values := make([]*ColumnValue, 0, 50)
	for _, updatedElement := range ctx.AllUpdatedElement() {
		values = append(values, self.doVisitUpdatedElement(updatedElement.(*parser.UpdatedElementContext)))
	}
	var where *Func
	if expression := ctx.Expression(); expression != nil {
		_where := self.visitExpression(expression)
		where = dealWhere(_where)
	}
	var order []SqlToken
	if orderByClause := ctx.OrderByClause(); orderByClause != nil {
		order = self.doVisitOrderByClause(orderByClause.(*parser.OrderByClauseContext))
	}
	var limit *Limit
	if limitClause := ctx.LimitClause(); limitClause != nil {
		limit = self.doVisitLimitClause(limitClause.(*parser.LimitClauseContext))
	}
	self.UpdateStatement = &UpdateExpr{
		Ignore: ignore,
		Table:  table,
		Values: values,
		Where:  where,
		Order:  order,
		Limit:  limit,
	}
}

func (self *visitor) doVisitQueryExpression(ctx *parser.QueryExpressionContext) *Link {
	/*
		LR_BRACKET querySpecification RR_BRACKET | LR_BRACKET queryExpression RR_BRACKET
	*/
	querySpecification := ctx.QuerySpecification()
	if querySpecification == nil {
		queryExpression := ctx.QueryExpression()
		if queryExpression != nil {
			return self.doVisitQueryExpression(queryExpression.(*parser.QueryExpressionContext))
		}
	}
	return self.doVisitQuerySpecification(querySpecification.(*parser.QuerySpecificationContext))
}

func (self *visitor) doVisitQuerySpecification(ctx *parser.QuerySpecificationContext) *Link {
	/* 每个 select 子句都需要入队
	SELECT DISTINCT? selectElements
	fromClause? orderByClause? limitClause?
	*/
	var distinct SqlToken
	distinct = &Const{Value: false}
	if d := ctx.DISTINCT(); d != nil {
		distinct = &Const{Value: true}
	}

	selectElements := self.doVisitSelectElements(ctx.SelectElements().(*parser.SelectElementsContext))

	var fromClause *From
	if f := ctx.FromClause(); f != nil {
		fromClause = self.doVisitFromClause(f.(*parser.FromClauseContext))
	} else {
		fromClause = nil
	}

	var orderByClause []SqlToken
	if o := ctx.OrderByClause(); o != nil {
		orderByClause = self.doVisitOrderByClause(o.(*parser.OrderByClauseContext))
	} else {
		orderByClause = nil
	}

	var limitClause *Limit
	if l := ctx.LimitClause(); l != nil {
		limitClause = self.doVisitLimitClause(l.(*parser.LimitClauseContext))
	} else {
		limitClause = nil
	}

	var selectClause []SqlToken
	selectClause = append(selectClause, distinct)
	selectClause = append(selectClause, selectElements...)

	SqlToken := &State{
		From:   fromClause,
		Order:  orderByClause,
		Limit:  limitClause,
		Select: selectClause,
	}

	tree := &SelectStatement{
		Id:   self.getId(),
		Tree: SqlToken,
	}

	// 入队
	self.SelectStatementsQueue.Enqueue(tree)
	// 返回的是链接
	return &Link{tree.Id}
}

func (self *visitor) doVisitSelectColumnElement(ctx *parser.SelectColumnElementContext) SqlToken {
	/*
		fullColumnName (AS? ID)?
	*/
	name := self.doVisitFullColumnName(ctx.FullColumnName().(*parser.FullColumnNameContext))

	alis := ctx.ID()
	if alis == nil {
		return name
	}
	alisStr := alis.GetText()
	return &As{
		NewName: &Const{Value: doVisitID(alisStr)},
		Value:   name,
	}
}

func doVisitID(value string) string {
	if value[0] == '`' {
		return value[1 : len(value)-1]
	}
	return value
}

func (self *visitor) doVisitSelectFunctionElement(ctx *parser.SelectFunctionElementContext) SqlToken {
	/*
		functionCall (AS? ID)?
	*/
	name := self.visitFunctionCall(ctx.FunctionCall())

	alis := ctx.ID()
	if alis == nil {
		return name
	}

	alisStr := alis.GetText()
	return &As{
		NewName: &Const{Value: doVisitID(alisStr)},
		Value:   name,
	}
}

func (self *visitor) doVisitSelectExpressionElement(ctx *parser.SelectExpressionElementContext) SqlToken {
	/*
		expression (AS? ID)?
	*/
	name := self.visitExpression(ctx.Expression()).(SqlToken)

	alis := ctx.ID()
	if alis == nil {
		return name
	}

	alisStr := alis.GetText()
	return &As{
		NewName: &Const{Value: doVisitID(alisStr)},
		Value:   name,
	}
}

func (self *visitor) doVisitSelectElements(ctx *parser.SelectElementsContext) []SqlToken {
	/*
		(star='*' | selectElement ) (',' selectElement)*
	*/
	var args []SqlToken
	if s := ctx.STAR(); s != nil {
		args = append(args, &Const{true})
	} else {
		args = append(args, &Const{false})
	}
	selectElements := ctx.AllSelectElement()
	for _, selectElement := range selectElements {
		args = append(args, self.visitSelectElement(selectElement))
	}
	return args
}

func dealWhere(_whereExpr interface{}) *Func {
	var whereExpr *Func
	switch realWhereExpr := _whereExpr.(type) {
	case *Func:
		whereExpr = realWhereExpr
	case []SqlToken:
		whereExpr = realWhereExpr[0].(*Func)
	default:
		logrus.Errorf("[doVisitFromClause] no such type [%T]", whereExpr)
		return nil
	}
	return whereExpr
}

func (self *visitor) doVisitFromClause(ctx *parser.FromClauseContext) *From {
	/*
		FROM tableSources
		(WHERE whereExpr=expression)?
		(
		  GROUP BY
		  groupByItem (',' groupByItem)*
		  (HAVING havingExpr=expression)?
		)?
	*/
	var groupExpr []*Var
	var havingExpr *Func
	var whereExpr *Func
	tableSources := self.doVisitTableSources(ctx.TableSources().(*parser.TableSourcesContext))
	if g := ctx.GROUP(); g != nil {
		var items []*Var
		_items := ctx.AllGroupByItem()
		for _, item := range _items {
			items = append(items, self.doVisitGroupByItem(item.(*parser.GroupByItemContext)).(*Var))
		}
		groupExpr = items
	} else {
		groupExpr = nil
	}
	if w := ctx.WHERE(); w != nil {
		_whereExpr := self.visitExpression(ctx.Expression(0))
		whereExpr = dealWhere(_whereExpr)
		if h := ctx.HAVING(); h != nil {
			havingExpr = self.visitExpression(ctx.Expression(1)).(*Func)
		} else {
			havingExpr = nil
		}
	} else {
		whereExpr = nil
		if h := ctx.HAVING(); h != nil {
			havingExpr = self.visitExpression(ctx.Expression(0)).(*Func)
		} else {
			havingExpr = nil
		}
	}
	return &From{
		Tables: tableSources,
		Where:  whereExpr,
		Group:  groupExpr,
		Having: havingExpr,
	}

}

func (self *visitor) doVisitGroupByItem(ctx *parser.GroupByItemContext) interface{} {
	/*
		expression order=(ASC | DESC)?
	*/
	if d := ctx.DESC(); d != nil {
		return &Desc{Value: self.visitExpression(ctx.Expression()).(*Var)}
	} else {
		return self.visitExpression(ctx.Expression())
	}
}

func (self *visitor) doVisitTableSourceBase(ctx *parser.TableSourceBaseContext) *Tables {
	/*
		tableSourceItem joinPart*
	*/
	var args []SqlToken
	args = append(args, self.visitTableSourceItem(ctx.TableSourceItem()))
	joinParts := ctx.AllJoinPart()
	for _, joinPart := range joinParts {
		args = append(args, self.visitJoinPart(joinPart))
	}
	return &Tables{Values: args}
}

func (self *visitor) doVisitFunctionArgs(ctx *parser.FunctionArgsContext) []interface{} {
	/*
		(allFunctionArg)
		(',' (allFunctionArg) )*
	*/
	var args []interface{}
	allFunctionArgs := ctx.AllAllFunctionArg()
	for _, allFunctionArg := range allFunctionArgs {
		args = append(args, self.doVisitAllFunctionArg(allFunctionArg.(*parser.AllFunctionArgContext)))
	}
	return args
}

func (self *visitor) doVisitAllFunctionArg(ctx *parser.AllFunctionArgContext) interface{} {
	/*
		'*' | ALL | constant | fullColumnName | functionCall | expression
	*/
	if ctx.STAR() != nil {
		return &Const{Value: "*"}
	}
	if ctx.ALL() != nil {
		return &Const{Value: "*"}
	}
	if c := ctx.Constant(); c != nil {
		return self.doVisitConstant(c.(*parser.ConstantContext))
	}
	if f := ctx.FullColumnName(); f != nil {
		return self.doVisitFullColumnName(f.(*parser.FullColumnNameContext))
	}
	if fc := ctx.FunctionCall(); fc != nil {
		return self.visitFunctionCall(fc)
	}
	return self.visitExpression(ctx.Expression())
}

func (self *visitor) doVisitTableSourceNested(ctx *parser.TableSourceNestedContext) *Tables {
	/*
		'(' tableSourceItem joinPart* ')'
	*/
	var args []SqlToken
	args = append(args, self.visitTableSourceItem(ctx.TableSourceItem()))
	joinParts := ctx.AllJoinPart()
	for _, joinPart := range joinParts {
		args = append(args, self.visitJoinPart(joinPart))
	}
	return &Tables{Values: args}
}

func (self *visitor) doVisitInnerJoin(ctx *parser.InnerJoinContext) *Inner {
	/*
		(INNER | CROSS)? JOIN tableSourceItem (ON expression)?
	*/
	if ctx.ON() != nil {
		return &Inner{
			Table:     self.visitTableSourceItem(ctx.TableSourceItem()).(*Var),
			Condition: self.visitExpression(ctx.Expression()).(*Func),
		}
	}
	return &Inner{
		Table:     self.visitTableSourceItem(ctx.TableSourceItem()).(*Var),
		Condition: nil,
	}
}

func (self *visitor) doVisitOuterJoin(ctx *parser.OuterJoinContext) *Outer {
	/*
		(LEFT | RIGHT) OUTER? JOIN tableSourceItem
		(
		  ON expression
		  | USING '(' uidList ')'
		)
	*/
	expression := ctx.Expression()
	var tag OuterTag
	var table SqlToken
	var condition *Func
	var using []SqlToken
	if l := ctx.LEFT(); l != nil {
		if expression != nil {
			tag = LeftOn
			table = self.visitTableSourceItem(ctx.TableSourceItem())
			condition = self.visitExpression(expression).(*Func)
		} else {
			tag = LeftUsing
			table = self.visitTableSourceItem(ctx.TableSourceItem())
			using = self.doVisitUidList(ctx.UidList().(*parser.UidListContext))
		}
	} else {
		if expression != nil {
			tag = RightOn
			table = self.visitTableSourceItem(ctx.TableSourceItem())
			condition = self.visitExpression(expression).(*Func)
		} else {
			tag = RightUsing
			table = self.visitTableSourceItem(ctx.TableSourceItem())
			using = self.doVisitUidList(ctx.UidList().(*parser.UidListContext))
		}
	}
	return &Outer{
		Tag:       tag,
		Table:     table,
		Condition: condition,
		Using:     using,
	}
}

func (self *visitor) doVisitUidList(ctx *parser.UidListContext) []SqlToken {
	/*
		ID (',' ID)*
	*/
	var args []SqlToken
	ids := ctx.AllID()
	for _, id := range ids {
		args = append(args, &Const{doVisitID(id.GetText())})
	}
	return args
}

func (self *visitor) doVisitAtomTableItem(ctx *parser.AtomTableItemContext) SqlToken {
	/*
		fullColumnName (AS? alias=ID)?
	*/
	name := self.doVisitFullColumnName(ctx.FullColumnName().(*parser.FullColumnNameContext))
	if alis := ctx.ID(); alis != nil {
		return &As{
			NewName: &Const{Value: doVisitID(alis.GetText())},
			Value:   name,
		}
	} else {
		return name
	}
}

func (self *visitor) doVisitSubqueryTableItem(ctx *parser.SubqueryTableItemContext) *As {
	/*
		selectStatement
		| '(' parenthesisSubquery=selectStatement ')'
		)
		AS? alias=ID
	*/
	return &As{
		NewName: &Const{Value: doVisitID(ctx.ID().GetText())},
		Value:   self.visitSelectStatement(ctx.SelectStatement()),
	}
}

func (self *visitor) doVisitTableSourcesItem(ctx *parser.TableSourcesItemContext) *Tables {
	/*
		'(' tableSources ')'
	*/
	return self.doVisitTableSources(ctx.TableSources().(*parser.TableSourcesContext))
}

func (self *visitor) doVisitTableSources(ctx *parser.TableSourcesContext) *Tables {
	/*
		tableSource (',' tableSource)*
	*/
	firstSource := self.visitTableSource(ctx.TableSource(0))
	tables := ctx.AllTableSource()
	for _, table := range tables[1:] {
		firstSource.Values = append(firstSource.Values,
			self.visitTableSource(table).Values...)
	}
	return firstSource
}

func (self *visitor) doVisitOrderByClause(ctx *parser.OrderByClauseContext) []SqlToken {
	/*
		ORDER BY orderByExpression (',' orderByExpression)*
	*/
	var args []SqlToken
	expressions := ctx.AllOrderByExpression()
	for _, expression := range expressions {
		args = append(args, self.doVisitOrderByExpression(expression.(*parser.OrderByExpressionContext)).(SqlToken))
	}
	return args
}

func (self *visitor) doVisitOrderByExpression(ctx *parser.OrderByExpressionContext) interface{} {
	/*
		expression order=(ASC | DESC)?
	*/
	if d := ctx.DESC(); d != nil {
		return &Desc{Value: self.visitExpression(ctx.Expression()).(*Var)}
	} else {
		return self.visitExpression(ctx.Expression())
	}
}

func (self *visitor) doVisitLimitClause(ctx *parser.LimitClauseContext) *Limit {
	/*
		LIMIT
		(
		  (offset=DECIMAL_LITERAL ',')? limit=DECIMAL_LITERAL
		  | limit=DECIMAL_LITERAL OFFSET offset=DECIMAL_LITERAL
		)
	*/
	var offset *Const
	var limit *Const
	if o := ctx.OFFSET(); o != nil {
		offsetNum, _ := strconv.Atoi(ctx.DECIMAL_LITERAL(1).GetText())
		limitNum, _ := strconv.Atoi(ctx.DECIMAL_LITERAL(0).GetText())
		offset = &Const{Value: offsetNum}
		limit = &Const{Value: limitNum}
	} else if c := ctx.COMMA(); c != nil {
		offsetNum, _ := strconv.Atoi(ctx.DECIMAL_LITERAL(0).GetText())
		limitNum, _ := strconv.Atoi(ctx.DECIMAL_LITERAL(1).GetText())
		offset = &Const{Value: offsetNum}
		limit = &Const{Value: limitNum}
	} else {
		offset = &Const{Value: 0}
		limitNum, _ := strconv.Atoi(ctx.DECIMAL_LITERAL(0).GetText())
		limit = &Const{Value: limitNum}
	}
	return &Limit{
		Limit:  limit,
		Offset: offset,
	}
}

func (self *visitor) doVisitNotExpression(ctx *parser.NotExpressionContext) *Func {
	/*
		notOperator=(NOT | '!') expression
	*/
	return &Func{
		Name: "not",
		Args: []interface{}{self.visitExpression(ctx.Expression())},
	}
}

func (self *visitor) doVisitLogicalExpression(ctx *parser.LogicalExpressionContext) *Func {
	/*
		expression logicalOperator expression
	*/
	left := self.visitExpression(ctx.Expression(0)).(SqlToken)
	right := self.visitExpression(ctx.Expression(1)).(SqlToken)
	var args []interface{}
	args = append(args, left)
	args = append(args, right)
	return &Func{
		Name: self.doVisitLogicalOperator(ctx.LogicalOperator().(*parser.LogicalOperatorContext)),
		Args: args,
	}
}

func (self *visitor) doVisitIsExpression(ctx *parser.IsExpressionContext) *Func {
	/*
		predicate IS NOT? testValue=(TRUE | FALSE | UNKNOWN)
	*/
	left := self.visitPredicate(ctx.Predicate()).(SqlToken)
	val := ctx.GetTestValue().GetText()
	var right SqlToken
	var name string
	if n := ctx.NOT(); n != nil {
		name = "is_not"
		if val == "TRUE" {
			right = &Const{Value: true}
		} else if val == "FALSE" {
			right = &Const{Value: false}
		} else {
			right = &Const{Value: nil}
		}
	} else {
		name = "is"
		//args = append(args, left)
		if val == "TRUE" {
			right = &Const{Value: true}
		} else if val == "FALSE" {
			right = &Const{Value: false}
		} else {
			right = &Const{Value: nil}
		}
	}
	return &Func{
		Name: name,
		Args: []interface{}{left, right},
	}
}

func (self *visitor) doVisitInPredicate(ctx *parser.InPredicateContext) *Func {
	/*
		predicate NOT? IN '(' (selectStatement | expressions) ')':
	*/
	left := self.visitPredicate(ctx.Predicate()).(SqlToken)
	var name string
	var right interface{}
	selectStatement := ctx.SelectStatement()
	if n := ctx.NOT(); n != nil {
		name = "not_in"
		if selectStatement != nil {
			right = self.visitSelectStatement(selectStatement)
		} else {
			right = self.doVisitExpressions(ctx.Expressions().(*parser.ExpressionsContext))
		}
	} else {
		name = "in"
		if selectStatement != nil {
			right = self.visitSelectStatement(selectStatement)
		} else {
			right = self.doVisitExpressions(ctx.Expressions().(*parser.ExpressionsContext))
		}
	}
	return &Func{
		Name: name,
		Args: []interface{}{left, right},
	}
}

func (self *visitor) doVisitExpressions(ctx *parser.ExpressionsContext) []SqlToken {
	/*
		expression (',' expression)*
	*/
	var args []SqlToken
	for _, expression := range ctx.AllExpression() {
		args = append(args, self.visitExpression(expression).(SqlToken))
	}
	return args
}

func (self *visitor) doVisitIsNullPredicate(ctx *parser.IsNullPredicateContext) *Func {
	/*
		predicate IS nullNotnull
	*/
	left := self.visitPredicate(ctx.Predicate()).(SqlToken)
	right := self.doVisitNullNotnull(ctx.NullNotnull().(*parser.NullNotnullContext))
	name := "is"
	return &Func{
		Name: name,
		Args: []interface{}{left, right},
	}
}

func (self *visitor) doVisitNullNotnull(ctx *parser.NullNotnullContext) SqlToken {
	/*
		NOT? (NULL_LITERAL | NULL_SPEC_LITERAL)
	*/
	if n := ctx.NULL_LITERAL(); n != nil {
		if _not := ctx.NOT(); _not != nil {
			return &Func{
				Name: "not",
				Args: []interface{}{&Const{Value: nil}},
			}
		} else {
			return &Const{Value: nil}
		}
	} else {
		if _not := ctx.NOT(); _not != nil {
			return &Func{
				Name: "not",
				Args: []interface{}{&Const{Value: "\n"}},
			}
		} else {
			return &Const{Value: "\n"}
		}
	}
}

func (self *visitor) doVisitBinaryComparasionPredicate(ctx *parser.BinaryComparasionPredicateContext) *Func {
	/*
		left=predicate comparisonOperator right=predicate
	*/
	left := self.visitPredicate(ctx.Predicate(0)).(SqlToken)
	op := self.doVisitComparisonOperator(ctx.ComparisonOperator().(*parser.ComparisonOperatorContext))
	right := self.visitPredicate(ctx.Predicate(1)).(SqlToken)
	return &Func{
		Name: op,
		Args: []interface{}{left, right},
	}
}

func (self *visitor) doVisitComparisonOperator(ctx *parser.ComparisonOperatorContext) string {
	/*
		'=' | '>' | '<' | '<' '=' | '>' '='
		| '<' '>' | '!' '=' | '<' '=' '>'
	*/
	equalSymbol := ctx.EQUAL_SYMBOL()
	greaterSymbol := ctx.GREATER_SYMBOL()
	lessSymbol := ctx.LESS_SYMBOL()
	exclamationSymbol := ctx.EXCLAMATION_SYMBOL()
	if equalSymbol != nil {
		if lessSymbol == nil {
			if exclamationSymbol == nil {
				if greaterSymbol == nil {
					return "="
				} else {
					return ">="
				}
			} else {
				return "!="
			}
		} else {
			if greaterSymbol == nil {
				return "<=>"
			} else {
				return "<="
			}
		}
	} else if greaterSymbol != nil {
		if lessSymbol != nil {
			return "<>"
		} else {
			return ">"
		}
	} else {
		return "<"
	}
}

func (self *visitor) doVisitSubqueryComparasionPredicate(ctx *parser.SubqueryComparasionPredicateContext) *Func {
	/*
		predicate comparisonOperator quantifier=(ALL | ANY | SOME) '(' selectStatement ')'
	*/
	left := self.visitPredicate(ctx.Predicate()).(SqlToken)
	op := self.doVisitComparisonOperator(ctx.ComparisonOperator().(*parser.ComparisonOperatorContext))
	selectStatement := self.visitSelectStatement(ctx.SelectStatement())
	return &Func{
		Name: ctx.GetQuantifier().GetText() + op,
		Args: []interface{}{left, selectStatement},
	}
}

func (self *visitor) doVisitBetweenPredicate(ctx *parser.BetweenPredicateContext) *Func {
	/*
		predicate NOT? BETWEEN predicate AND predicate
	*/
	left := self.visitPredicate(ctx.Predicate(0)).(SqlToken)
	start := self.visitPredicate(ctx.Predicate(1)).(SqlToken)
	end := self.visitPredicate(ctx.Predicate(2)).(SqlToken)
	var name string
	if n := ctx.NOT(); n != nil {
		name = "not_between"
	} else {
		name = "between"
	}
	return &Func{
		Name: name,
		Args: []interface{}{left, start, end},
	}
}

func (self *visitor) doVisitLikePredicate(ctx *parser.LikePredicateContext) *Func {
	/*
		predicate NOT? LIKE predicate
	*/
	left := self.visitPredicate(ctx.Predicate(0)).(SqlToken)
	right := self.visitPredicate(ctx.Predicate(1)).(SqlToken)
	var name string
	if n := ctx.NOT(); n != nil {
		name = "not_like"
	} else {
		name = "like"
	}
	return &Func{
		Name: name,
		Args: []interface{}{left, right},
	}
}

func (self *visitor) doVisitRegexpPredicate(ctx *parser.RegexpPredicateContext) *Func {
	/*
		predicate NOT? regex=(REGEXP | RLIKE) predicate
	*/
	left := self.visitPredicate(ctx.Predicate(0)).(SqlToken)
	right := self.visitPredicate(ctx.Predicate(1)).(SqlToken)
	var name string
	if n := ctx.NOT(); n != nil {
		name = "not_regexp"
	} else {
		name = "regexp"
	}
	return &Func{
		Name: name,
		Args: []interface{}{left, right},
	}
}

func (self *visitor) doVisitConstant(ctx *parser.ConstantContext) SqlToken {
	/*
		STRING_LITERAL | DECIMAL_LITERAL
		| '-' DECIMAL_LITERAL
		| booleanLiteral
		| REAL_LITERAL
		| NOT? nullLiteral=(NULL_LITERAL | NULL_SPEC_LITERAL)
	*/
	str := ctx.STRING_LITERAL()
	if str != nil {
		s := str.GetText()
		return &Const{Value: s[1 : len(s)-1]}
	}
	decimal := ctx.DECIMAL_LITERAL()
	minus := ctx.MINUS()
	if decimal != nil {
		num, _ := strconv.Atoi(decimal.GetText())
		var value int
		if minus != nil {
			value = -1 * num
		} else {
			value = num
		}
		return &Const{Value: value}
	}
	boolean := ctx.BooleanLiteral()
	if boolean != nil {
		return self.doVisitBooleanLiteral(boolean.(*parser.BooleanLiteralContext))
	}
	realNum := ctx.REAL_LITERAL()
	if realNum != nil {
		var value float64
		value, _ = strconv.ParseFloat(realNum.GetText(), 64)
		return &Const{Value: value}
	}
	if ctx.NOT() != nil {
		if ctx.NULL_LITERAL() != nil {
			return &Func{
				Name: "not",
				Args: []interface{}{&Const{Value: nil}},
			}
		} else {
			return &Const{Value: "\n"}
		}
	}
	if ctx.NULL_LITERAL() != nil {
		return &Const{Value: nil}
	} else {
		return &Const{Value: "\n"}
	}
}

func (self *visitor) doVisitBooleanLiteral(ctx *parser.BooleanLiteralContext) *Const {
	if ctx.TRUE() != nil {
		return &Const{Value: true}
	} else {
		return &Const{Value: false}
	}
}

func (self *visitor) doVisitFullColumnName(ctx *parser.FullColumnNameContext) *Var {
	/*
		ID DOT_ID*
	*/
	name := &Const{Value: ctx.ID().GetText()}
	columnName := &Var{Value: name}
	var curColumnName *Var
	curColumnName = columnName
	for _, next := range ctx.AllDOT_ID() {
		nextColumnName := &Var{Value: &Const{Value: next.GetText()[1:]}}
		curColumnName.Next = nextColumnName
		curColumnName = nextColumnName
	}
	return columnName
}

func (self *visitor) doVisitCaseFunctionCall(ctx *parser.CaseFunctionCallContext) *Case {
	/*
		CASE caseFuncAlternative+
		(ELSE elseArg=functionArg)? END
	*/
	var args []*CaseItem
	for _, funcAlternative := range ctx.AllCaseFuncAlternative() {
		args = append(args, self.doVisitCaseFuncAlternative(
			funcAlternative.(*parser.CaseFuncAlternativeContext)))
	}
	if ctx.ELSE() != nil {
		args = append(args, &CaseItem{
			Condition: nil,
			Value:     self.doVisitFunctionArg(ctx.FunctionArg().(*parser.FunctionArgContext)).(SqlToken),
		})
	}

	return &Case{Values: args}
}

func (self *visitor) doVisitCaseVarFunctionCall(ctx *parser.CaseVarFunctionCallContext) *Case {
	/*
		CASE fullColumnName caseFuncAlternative+
		   (ELSE elseArg=functionArg)? END
	*/
	var args []*CaseItem
	expr := self.doVisitFullColumnName(ctx.FullColumnName().(*parser.FullColumnNameContext))
	for _, funcAlternative := range ctx.AllCaseFuncAlternative() {
		caseItem := self.doVisitCaseFuncAlternative(
			funcAlternative.(*parser.CaseFuncAlternativeContext))
		caseItem.Condition = &Func{
			Name: "=",
			Args: []interface{}{caseItem.Condition, expr},
		}
		args = append(args, caseItem)
	}
	if ctx.ELSE() != nil {
		args = append(args, &CaseItem{
			Condition: nil,
			Value:     self.doVisitFunctionArg(ctx.FunctionArg().(*parser.FunctionArgContext)).(SqlToken),
		})
	}

	return &Case{Values: args}
}

func (self *visitor) doVisitFunctionArg(ctx *parser.FunctionArgContext) interface{} {
	/*
		constant | fullColumnName | functionCall | expression
	*/
	if c := ctx.Constant(); c != nil {
		return self.doVisitConstant(c.(*parser.ConstantContext))
	}
	if f := ctx.FullColumnName(); f != nil {
		return self.doVisitFullColumnName(f.(*parser.FullColumnNameContext))
	}
	if fc := ctx.FunctionCall(); fc != nil {
		return self.visitFunctionCall(fc)
	}
	return self.visitExpression(ctx.Expression())
}

func (self *visitor) doVisitCaseFuncAlternative(ctx *parser.CaseFuncAlternativeContext) *CaseItem {
	/*
		WHEN condition=functionArg
		THEN consequent=functionArg
	*/
	return &CaseItem{
		Condition: self.doVisitFunctionArg(ctx.FunctionArg(0).(*parser.FunctionArgContext)).(SqlToken),
		Value:     self.doVisitFunctionArg(ctx.FunctionArg(1).(*parser.FunctionArgContext)).(SqlToken),
	}
}

func (self *visitor) doVisitScalarFunctionCall(ctx *parser.ScalarFunctionCallContext) *Func {
	/*
		ID '(' functionArgs? ')'
	*/
	funcName := ctx.ID().GetText()
	_funcArgs := ctx.FunctionArgs()
	name := strings.ToLower(funcName)
	if _funcArgs != nil {
		funcArgs := self.doVisitFunctionArgs(_funcArgs.(*parser.FunctionArgsContext))
		return &Func{Name: name, Args: funcArgs}
	}
	return &Func{Name: name, Args: make([]interface{}, 0)}
}

func (self *visitor) doVisitUnaryExpressionAtom(ctx *parser.UnaryExpressionAtomContext) *Func {
	/*
		unaryOperator expressionAtom
	*/
	expressionAtom := self.visitExpressionAtom(ctx.ExpressionAtom()).(SqlToken)
	op := self.doVisitUnaryOperator(ctx.UnaryOperator().(*parser.UnaryOperatorContext))
	return &Func{Name: op, Args: []interface{}{expressionAtom}}
}

func (self *visitor) doVisitUnaryOperator(ctx *parser.UnaryOperatorContext) string {
	if ctx.NOT() != nil {
		return "signal_not"
	} else {
		return "signal_" + self.VisitChildren(ctx).(antlr.ParseTree).GetText()
	}
}

func (self *visitor) doVisitNestedExpressionAtom(ctx *parser.NestedExpressionAtomContext) []SqlToken {
	/*
		'(' expression (',' expression)* ')'
	*/
	var args []SqlToken
	items := ctx.AllExpression()
	for _, item := range items {
		args = append(args, self.visitExpression(item).(SqlToken))
	}
	return args
}

func (self *visitor) doVisitExistsExpressionAtom(ctx *parser.ExistsExpressionAtomContext) *Func {
	/*
		EXISTS '(' selectStatement ')'
	*/
	selectStatement := self.visitSelectStatement(ctx.SelectStatement())
	return &Func{Name: "exits", Args: []interface{}{selectStatement}}
}

func (self *visitor) doVisitMathExpressionAtom(ctx *parser.MathExpressionAtomContext) *Func {
	/*
		left=expressionAtom mathOperator right=expressionAtom
	*/
	left := self.visitExpressionAtom(ctx.ExpressionAtom(0)).(SqlToken)
	right := self.visitExpressionAtom(ctx.ExpressionAtom(1)).(SqlToken)
	op := self.doVisitMathOperator(ctx.MathOperator().(*parser.MathOperatorContext))
	return &Func{Name: op, Args: []interface{}{left, right}}
}

func (self *visitor) doVisitMathOperator(ctx *parser.MathOperatorContext) string {
	/*
		'+' | '-' | '--'
	*/
	return ctx.GetChildren()[0].(antlr.ParseTree).GetText()
}

func (self *visitor) doVisitPriorityMathExpressionAtom(ctx *parser.PriorityMathExpressionAtomContext) *Func {
	/*
		left=expressionAtom op=('*'|'/'|'%'| DIV | MOD) right=expressionAtom
	*/
	left := self.visitExpressionAtom(ctx.ExpressionAtom(0)).(SqlToken)
	right := self.visitExpressionAtom(ctx.ExpressionAtom(1)).(SqlToken)
	return &Func{Name: ctx.GetOp().GetText(), Args: []interface{}{left, right}}
}

func (self *visitor) doVisitLogicalOperator(ctx *parser.LogicalOperatorContext) string {
	/*
		AND | OR
	*/
	if ctx.AND() != nil {
		return "and"
	} else {
		return "or"
	}
}

func (self *visitor) visitSpecificFunctionCall(node antlr.Tree) SqlToken {
	switch ctx := node.(type) {
	case *parser.CaseFunctionCallContext:
		return self.doVisitCaseFunctionCall(ctx)
	case *parser.CaseVarFunctionCallContext:
		return self.doVisitCaseVarFunctionCall(ctx)
	default:
		logrus.Errorf("[visitSpecificFunctionCall] no such type [%T]", ctx)
		return nil
	}
}

func (self *visitor) visitFunctionCall(node antlr.Tree) SqlToken {
	switch ctx := node.(type) {
	case *parser.ScalarFunctionCallContext:
		return self.doVisitScalarFunctionCall(ctx)
	case *parser.SpecificFunctionCallContext:
		return self.visitSpecificFunctionCall(ctx.SpecificFunction())
	default:
		logrus.Errorf("[visitFunctionCall] no such type [%T]", ctx)
		return nil
	}
}

func (self *visitor) visitExpression(node antlr.Tree) interface{} {
	switch ctx := node.(type) {
	case *parser.NotExpressionContext:
		return self.doVisitNotExpression(ctx)
	case *parser.LogicalExpressionContext:
		return self.doVisitLogicalExpression(ctx)
	case *parser.IsExpressionContext:
		return self.doVisitIsExpression(ctx)
	case *parser.PredicateExpressionContext:
		return self.visitPredicate(ctx.Predicate())
	default:
		logrus.Errorf("[visitExpression] no such type [%T]", ctx)
		return nil
	}
}

func (self *visitor) visitPredicate(node antlr.Tree) interface{} {
	switch ctx := node.(type) {
	case *parser.InPredicateContext:
		return self.doVisitInPredicate(ctx)
	case *parser.IsNullPredicateContext:
		return self.doVisitIsNullPredicate(ctx)
	case *parser.BinaryComparasionPredicateContext:
		return self.doVisitBinaryComparasionPredicate(ctx)
	case *parser.SubqueryComparasionPredicateContext:
		return self.doVisitSubqueryComparasionPredicate(ctx)
	case *parser.BetweenPredicateContext:
		return self.doVisitBetweenPredicate(ctx)
	case *parser.LikePredicateContext:
		return self.doVisitLikePredicate(ctx)
	case *parser.RegexpPredicateContext:
		return self.doVisitRegexpPredicate(ctx)
	case *parser.ExpressionAtomPredicateContext:
		return self.visitExpressionAtom(ctx.ExpressionAtom())
	default:
		logrus.Errorf("[visitPredicate] no such type [%T]", ctx)
		return nil
	}
}

func (self *visitor) visitExpressionAtom(node antlr.Tree) interface{} {
	switch ctx := node.(type) {
	case *parser.ConstantExpressionAtomContext:
		return self.doVisitConstant(ctx.Constant().(*parser.ConstantContext))
	case *parser.FullColumnNameExpressionAtomContext:
		return self.doVisitFullColumnName(ctx.FullColumnName().(*parser.FullColumnNameContext))
	case *parser.FunctionCallExpressionAtomContext:
		return self.visitFunctionCall(ctx.FunctionCall())
	case *parser.UnaryExpressionAtomContext:
		return self.doVisitUnaryExpressionAtom(ctx)
	case *parser.NestedExpressionAtomContext:
		return self.doVisitNestedExpressionAtom(ctx)
	case *parser.ExistsExpressionAtomContext:
		return self.doVisitExistsExpressionAtom(ctx)
	case *parser.SubqueryExpressionAtomContext:
		return self.visitSelectStatement(ctx.SelectStatement())
	case *parser.PriorityMathExpressionAtomContext:
		return self.doVisitPriorityMathExpressionAtom(ctx)
	case *parser.MathExpressionAtomContext:
		return self.doVisitMathExpressionAtom(ctx)
	default:
		logrus.Errorf("[visitExpressionAtom] no such type [%T]", ctx)
		return nil
	}
}

func (self *visitor) visitSelectElement(node antlr.Tree) SqlToken {
	switch ctx := node.(type) {
	case *parser.SelectColumnElementContext:
		return self.doVisitSelectColumnElement(ctx)
	case *parser.SelectFunctionElementContext:
		return self.doVisitSelectFunctionElement(ctx)
	case *parser.SelectExpressionElementContext:
		return self.doVisitSelectExpressionElement(ctx)
	default:
		logrus.Errorf("[visitSelectElement] no such type [%T]", ctx)
		return nil
	}
}

func (self *visitor) visitTableSourceItem(node antlr.Tree) SqlToken {
	switch ctx := node.(type) {
	case *parser.AtomTableItemContext:
		return self.doVisitAtomTableItem(ctx)
	case *parser.SubqueryTableItemContext:
		return self.doVisitSubqueryTableItem(ctx)
	case *parser.TableSourcesItemContext:
		return self.doVisitTableSourcesItem(ctx)
	default:
		logrus.Errorf("[visitTableSourceItem] no such type [%T]", ctx)
		return nil
	}
}

func (self *visitor) visitJoinPart(node antlr.Tree) SqlToken {
	switch ctx := node.(type) {
	case *parser.InnerJoinContext:
		return self.doVisitInnerJoin(ctx)
	case *parser.OuterJoinContext:
		return self.doVisitOuterJoin(ctx)
	default:
		logrus.Errorf("[visitJoinPart] no such type [%T]", ctx)
		return nil
	}
}

func (self *visitor) visitTableSource(node antlr.Tree) *Tables {
	switch ctx := node.(type) {
	case *parser.TableSourceBaseContext:
		return self.doVisitTableSourceBase(ctx)
	case *parser.TableSourceNestedContext:
		return self.doVisitTableSourceNested(ctx)
	default:
		logrus.Errorf("[visitTableSource] no such type [%T]", ctx)
		return nil
	}
}

func (self *visitor) visitSelectStatement(node antlr.Tree) *Link {
	switch ctx := node.(type) {
	case *parser.SimpleSelectContext:
		return self.doVisitQuerySpecification(ctx.QuerySpecification().(*parser.QuerySpecificationContext))
	case *parser.ParenthesisSelectContext:
		return self.doVisitQueryExpression(ctx.QueryExpression().(*parser.QueryExpressionContext))
	default:
		logrus.Errorf("[visitSelectStatement] no such type [%T]", ctx)
		return nil
	}
}
