// Code generated from /Users/zhangzhiqiang/go/src/github.com/Aiyane/golinq/sql-parser/MySqlParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // MySqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMySqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMySqlParserVisitor) VisitDmlStatement(ctx *DmlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleSelect(ctx *SimpleSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitParenthesisSelect(ctx *ParenthesisSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOrderByExpression(ctx *OrderByExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourceBase(ctx *TableSourceBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourceNested(ctx *TableSourceNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSources(ctx *TableSourcesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAtomTableItem(ctx *AtomTableItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInnerJoin(ctx *InnerJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOuterJoin(ctx *OuterJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQueryExpression(ctx *QueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectElements(ctx *SelectElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectColumnElement(ctx *SelectColumnElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectFunctionElement(ctx *SelectFunctionElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectExpressionElement(ctx *SelectExpressionElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGroupByItem(ctx *GroupByItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullColumnName(ctx *FullColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNullNotnull(ctx *NullNotnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUidList(ctx *UidListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseVarFunctionCall(ctx *CaseVarFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseFunctionCall(ctx *CaseFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAllFunctionArg(ctx *AllFunctionArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionArg(ctx *FunctionArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIsExpression(ctx *IsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPredicateExpression(ctx *PredicateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInPredicate(ctx *InPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryExpressionAtom(ctx *SubqueryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPriorityMathExpressionAtom(ctx *PriorityMathExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExistsExpressionAtom(ctx *ExistsExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLogicalOperator(ctx *LogicalOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMathOperator(ctx *MathOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInsertStatementValue(ctx *InsertStatementValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUpdatedElement(ctx *UpdatedElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressionsWithDefaults(ctx *ExpressionsWithDefaultsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}
