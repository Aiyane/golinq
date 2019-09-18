// Code generated from /Users/zhangzhiqiang/go/src/github.com/Aiyane/golinq/sql-parser/MySqlParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // MySqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMySqlParserListener is a complete listener for a parse tree produced by MySqlParser.
type BaseMySqlParserListener struct{}

var _ MySqlParserListener = &BaseMySqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMySqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMySqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMySqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMySqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDmlStatement is called when production dmlStatement is entered.
func (s *BaseMySqlParserListener) EnterDmlStatement(ctx *DmlStatementContext) {}

// ExitDmlStatement is called when production dmlStatement is exited.
func (s *BaseMySqlParserListener) ExitDmlStatement(ctx *DmlStatementContext) {}

// EnterSimpleSelect is called when production simpleSelect is entered.
func (s *BaseMySqlParserListener) EnterSimpleSelect(ctx *SimpleSelectContext) {}

// ExitSimpleSelect is called when production simpleSelect is exited.
func (s *BaseMySqlParserListener) ExitSimpleSelect(ctx *SimpleSelectContext) {}

// EnterParenthesisSelect is called when production parenthesisSelect is entered.
func (s *BaseMySqlParserListener) EnterParenthesisSelect(ctx *ParenthesisSelectContext) {}

// ExitParenthesisSelect is called when production parenthesisSelect is exited.
func (s *BaseMySqlParserListener) ExitParenthesisSelect(ctx *ParenthesisSelectContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseMySqlParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseMySqlParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByExpression is called when production orderByExpression is entered.
func (s *BaseMySqlParserListener) EnterOrderByExpression(ctx *OrderByExpressionContext) {}

// ExitOrderByExpression is called when production orderByExpression is exited.
func (s *BaseMySqlParserListener) ExitOrderByExpression(ctx *OrderByExpressionContext) {}

// EnterTableSourceBase is called when production tableSourceBase is entered.
func (s *BaseMySqlParserListener) EnterTableSourceBase(ctx *TableSourceBaseContext) {}

// ExitTableSourceBase is called when production tableSourceBase is exited.
func (s *BaseMySqlParserListener) ExitTableSourceBase(ctx *TableSourceBaseContext) {}

// EnterTableSourceNested is called when production tableSourceNested is entered.
func (s *BaseMySqlParserListener) EnterTableSourceNested(ctx *TableSourceNestedContext) {}

// ExitTableSourceNested is called when production tableSourceNested is exited.
func (s *BaseMySqlParserListener) ExitTableSourceNested(ctx *TableSourceNestedContext) {}

// EnterTableSources is called when production tableSources is entered.
func (s *BaseMySqlParserListener) EnterTableSources(ctx *TableSourcesContext) {}

// ExitTableSources is called when production tableSources is exited.
func (s *BaseMySqlParserListener) ExitTableSources(ctx *TableSourcesContext) {}

// EnterAtomTableItem is called when production atomTableItem is entered.
func (s *BaseMySqlParserListener) EnterAtomTableItem(ctx *AtomTableItemContext) {}

// ExitAtomTableItem is called when production atomTableItem is exited.
func (s *BaseMySqlParserListener) ExitAtomTableItem(ctx *AtomTableItemContext) {}

// EnterSubqueryTableItem is called when production subqueryTableItem is entered.
func (s *BaseMySqlParserListener) EnterSubqueryTableItem(ctx *SubqueryTableItemContext) {}

// ExitSubqueryTableItem is called when production subqueryTableItem is exited.
func (s *BaseMySqlParserListener) ExitSubqueryTableItem(ctx *SubqueryTableItemContext) {}

// EnterTableSourcesItem is called when production tableSourcesItem is entered.
func (s *BaseMySqlParserListener) EnterTableSourcesItem(ctx *TableSourcesItemContext) {}

// ExitTableSourcesItem is called when production tableSourcesItem is exited.
func (s *BaseMySqlParserListener) ExitTableSourcesItem(ctx *TableSourcesItemContext) {}

// EnterInnerJoin is called when production innerJoin is entered.
func (s *BaseMySqlParserListener) EnterInnerJoin(ctx *InnerJoinContext) {}

// ExitInnerJoin is called when production innerJoin is exited.
func (s *BaseMySqlParserListener) ExitInnerJoin(ctx *InnerJoinContext) {}

// EnterOuterJoin is called when production outerJoin is entered.
func (s *BaseMySqlParserListener) EnterOuterJoin(ctx *OuterJoinContext) {}

// ExitOuterJoin is called when production outerJoin is exited.
func (s *BaseMySqlParserListener) ExitOuterJoin(ctx *OuterJoinContext) {}

// EnterQueryExpression is called when production queryExpression is entered.
func (s *BaseMySqlParserListener) EnterQueryExpression(ctx *QueryExpressionContext) {}

// ExitQueryExpression is called when production queryExpression is exited.
func (s *BaseMySqlParserListener) ExitQueryExpression(ctx *QueryExpressionContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseMySqlParserListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseMySqlParserListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterSelectElements is called when production selectElements is entered.
func (s *BaseMySqlParserListener) EnterSelectElements(ctx *SelectElementsContext) {}

// ExitSelectElements is called when production selectElements is exited.
func (s *BaseMySqlParserListener) ExitSelectElements(ctx *SelectElementsContext) {}

// EnterSelectColumnElement is called when production selectColumnElement is entered.
func (s *BaseMySqlParserListener) EnterSelectColumnElement(ctx *SelectColumnElementContext) {}

// ExitSelectColumnElement is called when production selectColumnElement is exited.
func (s *BaseMySqlParserListener) ExitSelectColumnElement(ctx *SelectColumnElementContext) {}

// EnterSelectFunctionElement is called when production selectFunctionElement is entered.
func (s *BaseMySqlParserListener) EnterSelectFunctionElement(ctx *SelectFunctionElementContext) {}

// ExitSelectFunctionElement is called when production selectFunctionElement is exited.
func (s *BaseMySqlParserListener) ExitSelectFunctionElement(ctx *SelectFunctionElementContext) {}

// EnterSelectExpressionElement is called when production selectExpressionElement is entered.
func (s *BaseMySqlParserListener) EnterSelectExpressionElement(ctx *SelectExpressionElementContext) {}

// ExitSelectExpressionElement is called when production selectExpressionElement is exited.
func (s *BaseMySqlParserListener) ExitSelectExpressionElement(ctx *SelectExpressionElementContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseMySqlParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseMySqlParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterGroupByItem is called when production groupByItem is entered.
func (s *BaseMySqlParserListener) EnterGroupByItem(ctx *GroupByItemContext) {}

// ExitGroupByItem is called when production groupByItem is exited.
func (s *BaseMySqlParserListener) ExitGroupByItem(ctx *GroupByItemContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseMySqlParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseMySqlParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterFullColumnName is called when production fullColumnName is entered.
func (s *BaseMySqlParserListener) EnterFullColumnName(ctx *FullColumnNameContext) {}

// ExitFullColumnName is called when production fullColumnName is exited.
func (s *BaseMySqlParserListener) ExitFullColumnName(ctx *FullColumnNameContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseMySqlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseMySqlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterNullNotnull is called when production nullNotnull is entered.
func (s *BaseMySqlParserListener) EnterNullNotnull(ctx *NullNotnullContext) {}

// ExitNullNotnull is called when production nullNotnull is exited.
func (s *BaseMySqlParserListener) ExitNullNotnull(ctx *NullNotnullContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseMySqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseMySqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterUidList is called when production uidList is entered.
func (s *BaseMySqlParserListener) EnterUidList(ctx *UidListContext) {}

// ExitUidList is called when production uidList is exited.
func (s *BaseMySqlParserListener) ExitUidList(ctx *UidListContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseMySqlParserListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseMySqlParserListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterSpecificFunctionCall is called when production specificFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterSpecificFunctionCall(ctx *SpecificFunctionCallContext) {}

// ExitSpecificFunctionCall is called when production specificFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitSpecificFunctionCall(ctx *SpecificFunctionCallContext) {}

// EnterScalarFunctionCall is called when production scalarFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterScalarFunctionCall(ctx *ScalarFunctionCallContext) {}

// ExitScalarFunctionCall is called when production scalarFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitScalarFunctionCall(ctx *ScalarFunctionCallContext) {}

// EnterCaseVarFunctionCall is called when production caseVarFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterCaseVarFunctionCall(ctx *CaseVarFunctionCallContext) {}

// ExitCaseVarFunctionCall is called when production caseVarFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitCaseVarFunctionCall(ctx *CaseVarFunctionCallContext) {}

// EnterCaseFunctionCall is called when production caseFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterCaseFunctionCall(ctx *CaseFunctionCallContext) {}

// ExitCaseFunctionCall is called when production caseFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitCaseFunctionCall(ctx *CaseFunctionCallContext) {}

// EnterCaseFuncAlternative is called when production caseFuncAlternative is entered.
func (s *BaseMySqlParserListener) EnterCaseFuncAlternative(ctx *CaseFuncAlternativeContext) {}

// ExitCaseFuncAlternative is called when production caseFuncAlternative is exited.
func (s *BaseMySqlParserListener) ExitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BaseMySqlParserListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BaseMySqlParserListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterAllFunctionArg is called when production allFunctionArg is entered.
func (s *BaseMySqlParserListener) EnterAllFunctionArg(ctx *AllFunctionArgContext) {}

// ExitAllFunctionArg is called when production allFunctionArg is exited.
func (s *BaseMySqlParserListener) ExitAllFunctionArg(ctx *AllFunctionArgContext) {}

// EnterFunctionArg is called when production functionArg is entered.
func (s *BaseMySqlParserListener) EnterFunctionArg(ctx *FunctionArgContext) {}

// ExitFunctionArg is called when production functionArg is exited.
func (s *BaseMySqlParserListener) ExitFunctionArg(ctx *FunctionArgContext) {}

// EnterIsExpression is called when production isExpression is entered.
func (s *BaseMySqlParserListener) EnterIsExpression(ctx *IsExpressionContext) {}

// ExitIsExpression is called when production isExpression is exited.
func (s *BaseMySqlParserListener) ExitIsExpression(ctx *IsExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseMySqlParserListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseMySqlParserListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseMySqlParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseMySqlParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterPredicateExpression is called when production predicateExpression is entered.
func (s *BaseMySqlParserListener) EnterPredicateExpression(ctx *PredicateExpressionContext) {}

// ExitPredicateExpression is called when production predicateExpression is exited.
func (s *BaseMySqlParserListener) ExitPredicateExpression(ctx *PredicateExpressionContext) {}

// EnterExpressionAtomPredicate is called when production expressionAtomPredicate is entered.
func (s *BaseMySqlParserListener) EnterExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// ExitExpressionAtomPredicate is called when production expressionAtomPredicate is exited.
func (s *BaseMySqlParserListener) ExitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// EnterInPredicate is called when production inPredicate is entered.
func (s *BaseMySqlParserListener) EnterInPredicate(ctx *InPredicateContext) {}

// ExitInPredicate is called when production inPredicate is exited.
func (s *BaseMySqlParserListener) ExitInPredicate(ctx *InPredicateContext) {}

// EnterSubqueryComparasionPredicate is called when production subqueryComparasionPredicate is entered.
func (s *BaseMySqlParserListener) EnterSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) {}

// ExitSubqueryComparasionPredicate is called when production subqueryComparasionPredicate is exited.
func (s *BaseMySqlParserListener) ExitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) {}

// EnterBetweenPredicate is called when production betweenPredicate is entered.
func (s *BaseMySqlParserListener) EnterBetweenPredicate(ctx *BetweenPredicateContext) {}

// ExitBetweenPredicate is called when production betweenPredicate is exited.
func (s *BaseMySqlParserListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {}

// EnterBinaryComparasionPredicate is called when production binaryComparasionPredicate is entered.
func (s *BaseMySqlParserListener) EnterBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) {}

// ExitBinaryComparasionPredicate is called when production binaryComparasionPredicate is exited.
func (s *BaseMySqlParserListener) ExitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) {}

// EnterIsNullPredicate is called when production isNullPredicate is entered.
func (s *BaseMySqlParserListener) EnterIsNullPredicate(ctx *IsNullPredicateContext) {}

// ExitIsNullPredicate is called when production isNullPredicate is exited.
func (s *BaseMySqlParserListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseMySqlParserListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseMySqlParserListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterRegexpPredicate is called when production regexpPredicate is entered.
func (s *BaseMySqlParserListener) EnterRegexpPredicate(ctx *RegexpPredicateContext) {}

// ExitRegexpPredicate is called when production regexpPredicate is exited.
func (s *BaseMySqlParserListener) ExitRegexpPredicate(ctx *RegexpPredicateContext) {}

// EnterUnaryExpressionAtom is called when production unaryExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// ExitUnaryExpressionAtom is called when production unaryExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// EnterSubqueryExpressionAtom is called when production subqueryExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterSubqueryExpressionAtom(ctx *SubqueryExpressionAtomContext) {}

// ExitSubqueryExpressionAtom is called when production subqueryExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitSubqueryExpressionAtom(ctx *SubqueryExpressionAtomContext) {}

// EnterPriorityMathExpressionAtom is called when production priorityMathExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterPriorityMathExpressionAtom(ctx *PriorityMathExpressionAtomContext) {}

// ExitPriorityMathExpressionAtom is called when production priorityMathExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitPriorityMathExpressionAtom(ctx *PriorityMathExpressionAtomContext) {}

// EnterConstantExpressionAtom is called when production constantExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// ExitConstantExpressionAtom is called when production constantExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// EnterFunctionCallExpressionAtom is called when production functionCallExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) {}

// ExitFunctionCallExpressionAtom is called when production functionCallExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) {}

// EnterFullColumnNameExpressionAtom is called when production fullColumnNameExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) {}

// ExitFullColumnNameExpressionAtom is called when production fullColumnNameExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) {}

// EnterNestedExpressionAtom is called when production nestedExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterNestedExpressionAtom(ctx *NestedExpressionAtomContext) {}

// ExitNestedExpressionAtom is called when production nestedExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitNestedExpressionAtom(ctx *NestedExpressionAtomContext) {}

// EnterMathExpressionAtom is called when production mathExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// ExitMathExpressionAtom is called when production mathExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// EnterExistsExpressionAtom is called when production existsExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterExistsExpressionAtom(ctx *ExistsExpressionAtomContext) {}

// ExitExistsExpressionAtom is called when production existsExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitExistsExpressionAtom(ctx *ExistsExpressionAtomContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseMySqlParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseMySqlParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseMySqlParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseMySqlParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BaseMySqlParserListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BaseMySqlParserListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterMathOperator is called when production mathOperator is entered.
func (s *BaseMySqlParserListener) EnterMathOperator(ctx *MathOperatorContext) {}

// ExitMathOperator is called when production mathOperator is exited.
func (s *BaseMySqlParserListener) ExitMathOperator(ctx *MathOperatorContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseMySqlParserListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseMySqlParserListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterInsertStatementValue is called when production insertStatementValue is entered.
func (s *BaseMySqlParserListener) EnterInsertStatementValue(ctx *InsertStatementValueContext) {}

// ExitInsertStatementValue is called when production insertStatementValue is exited.
func (s *BaseMySqlParserListener) ExitInsertStatementValue(ctx *InsertStatementValueContext) {}

// EnterUpdatedElement is called when production updatedElement is entered.
func (s *BaseMySqlParserListener) EnterUpdatedElement(ctx *UpdatedElementContext) {}

// ExitUpdatedElement is called when production updatedElement is exited.
func (s *BaseMySqlParserListener) ExitUpdatedElement(ctx *UpdatedElementContext) {}

// EnterExpressionsWithDefaults is called when production expressionsWithDefaults is entered.
func (s *BaseMySqlParserListener) EnterExpressionsWithDefaults(ctx *ExpressionsWithDefaultsContext) {}

// ExitExpressionsWithDefaults is called when production expressionsWithDefaults is exited.
func (s *BaseMySqlParserListener) ExitExpressionsWithDefaults(ctx *ExpressionsWithDefaultsContext) {}

// EnterExpressionOrDefault is called when production expressionOrDefault is entered.
func (s *BaseMySqlParserListener) EnterExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// ExitExpressionOrDefault is called when production expressionOrDefault is exited.
func (s *BaseMySqlParserListener) ExitExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseMySqlParserListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseMySqlParserListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseMySqlParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseMySqlParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}
