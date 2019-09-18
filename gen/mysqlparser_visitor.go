// Code generated from /Users/zhangzhiqiang/go/src/github.com/Aiyane/golinq/sql-parser/MySqlParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // MySqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by MySqlParser.
type MySqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MySqlParser#dmlStatement.
	VisitDmlStatement(ctx *DmlStatementContext) interface{}

	// Visit a parse tree produced by MySqlParser#simpleSelect.
	VisitSimpleSelect(ctx *SimpleSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#parenthesisSelect.
	VisitParenthesisSelect(ctx *ParenthesisSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#orderByExpression.
	VisitOrderByExpression(ctx *OrderByExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableSourceBase.
	VisitTableSourceBase(ctx *TableSourceBaseContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableSourceNested.
	VisitTableSourceNested(ctx *TableSourceNestedContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableSources.
	VisitTableSources(ctx *TableSourcesContext) interface{}

	// Visit a parse tree produced by MySqlParser#atomTableItem.
	VisitAtomTableItem(ctx *AtomTableItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#subqueryTableItem.
	VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableSourcesItem.
	VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#innerJoin.
	VisitInnerJoin(ctx *InnerJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#outerJoin.
	VisitOuterJoin(ctx *OuterJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#queryExpression.
	VisitQueryExpression(ctx *QueryExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#querySpecification.
	VisitQuerySpecification(ctx *QuerySpecificationContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectElements.
	VisitSelectElements(ctx *SelectElementsContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectColumnElement.
	VisitSelectColumnElement(ctx *SelectColumnElementContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectFunctionElement.
	VisitSelectFunctionElement(ctx *SelectFunctionElementContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectExpressionElement.
	VisitSelectExpressionElement(ctx *SelectExpressionElementContext) interface{}

	// Visit a parse tree produced by MySqlParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#groupByItem.
	VisitGroupByItem(ctx *GroupByItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#fullColumnName.
	VisitFullColumnName(ctx *FullColumnNameContext) interface{}

	// Visit a parse tree produced by MySqlParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by MySqlParser#nullNotnull.
	VisitNullNotnull(ctx *NullNotnullContext) interface{}

	// Visit a parse tree produced by MySqlParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by MySqlParser#uidList.
	VisitUidList(ctx *UidListContext) interface{}

	// Visit a parse tree produced by MySqlParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by MySqlParser#specificFunctionCall.
	VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#scalarFunctionCall.
	VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#caseVarFunctionCall.
	VisitCaseVarFunctionCall(ctx *CaseVarFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#caseFunctionCall.
	VisitCaseFunctionCall(ctx *CaseFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#caseFuncAlternative.
	VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) interface{}

	// Visit a parse tree produced by MySqlParser#functionArgs.
	VisitFunctionArgs(ctx *FunctionArgsContext) interface{}

	// Visit a parse tree produced by MySqlParser#allFunctionArg.
	VisitAllFunctionArg(ctx *AllFunctionArgContext) interface{}

	// Visit a parse tree produced by MySqlParser#functionArg.
	VisitFunctionArg(ctx *FunctionArgContext) interface{}

	// Visit a parse tree produced by MySqlParser#isExpression.
	VisitIsExpression(ctx *IsExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#predicateExpression.
	VisitPredicateExpression(ctx *PredicateExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#expressionAtomPredicate.
	VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#inPredicate.
	VisitInPredicate(ctx *InPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#subqueryComparasionPredicate.
	VisitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#betweenPredicate.
	VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#binaryComparasionPredicate.
	VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#isNullPredicate.
	VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#regexpPredicate.
	VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#unaryExpressionAtom.
	VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#subqueryExpressionAtom.
	VisitSubqueryExpressionAtom(ctx *SubqueryExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#priorityMathExpressionAtom.
	VisitPriorityMathExpressionAtom(ctx *PriorityMathExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#constantExpressionAtom.
	VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#functionCallExpressionAtom.
	VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#fullColumnNameExpressionAtom.
	VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#nestedExpressionAtom.
	VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#mathExpressionAtom.
	VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#existsExpressionAtom.
	VisitExistsExpressionAtom(ctx *ExistsExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#logicalOperator.
	VisitLogicalOperator(ctx *LogicalOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#mathOperator.
	VisitMathOperator(ctx *MathOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by MySqlParser#insertStatementValue.
	VisitInsertStatementValue(ctx *InsertStatementValueContext) interface{}

	// Visit a parse tree produced by MySqlParser#updatedElement.
	VisitUpdatedElement(ctx *UpdatedElementContext) interface{}

	// Visit a parse tree produced by MySqlParser#expressionsWithDefaults.
	VisitExpressionsWithDefaults(ctx *ExpressionsWithDefaultsContext) interface{}

	// Visit a parse tree produced by MySqlParser#expressionOrDefault.
	VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{}

	// Visit a parse tree produced by MySqlParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by MySqlParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

}