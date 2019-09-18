// Code generated from /Users/zhangzhiqiang/go/src/github.com/Aiyane/golinq/sql-parser/MySqlParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // MySqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MySqlParserListener is a complete listener for a parse tree produced by MySqlParser.
type MySqlParserListener interface {
	antlr.ParseTreeListener

	// EnterDmlStatement is called when entering the dmlStatement production.
	EnterDmlStatement(c *DmlStatementContext)

	// EnterSimpleSelect is called when entering the simpleSelect production.
	EnterSimpleSelect(c *SimpleSelectContext)

	// EnterParenthesisSelect is called when entering the parenthesisSelect production.
	EnterParenthesisSelect(c *ParenthesisSelectContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByExpression is called when entering the orderByExpression production.
	EnterOrderByExpression(c *OrderByExpressionContext)

	// EnterTableSourceBase is called when entering the tableSourceBase production.
	EnterTableSourceBase(c *TableSourceBaseContext)

	// EnterTableSourceNested is called when entering the tableSourceNested production.
	EnterTableSourceNested(c *TableSourceNestedContext)

	// EnterTableSources is called when entering the tableSources production.
	EnterTableSources(c *TableSourcesContext)

	// EnterAtomTableItem is called when entering the atomTableItem production.
	EnterAtomTableItem(c *AtomTableItemContext)

	// EnterSubqueryTableItem is called when entering the subqueryTableItem production.
	EnterSubqueryTableItem(c *SubqueryTableItemContext)

	// EnterTableSourcesItem is called when entering the tableSourcesItem production.
	EnterTableSourcesItem(c *TableSourcesItemContext)

	// EnterInnerJoin is called when entering the innerJoin production.
	EnterInnerJoin(c *InnerJoinContext)

	// EnterOuterJoin is called when entering the outerJoin production.
	EnterOuterJoin(c *OuterJoinContext)

	// EnterQueryExpression is called when entering the queryExpression production.
	EnterQueryExpression(c *QueryExpressionContext)

	// EnterQuerySpecification is called when entering the querySpecification production.
	EnterQuerySpecification(c *QuerySpecificationContext)

	// EnterSelectElements is called when entering the selectElements production.
	EnterSelectElements(c *SelectElementsContext)

	// EnterSelectColumnElement is called when entering the selectColumnElement production.
	EnterSelectColumnElement(c *SelectColumnElementContext)

	// EnterSelectFunctionElement is called when entering the selectFunctionElement production.
	EnterSelectFunctionElement(c *SelectFunctionElementContext)

	// EnterSelectExpressionElement is called when entering the selectExpressionElement production.
	EnterSelectExpressionElement(c *SelectExpressionElementContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterGroupByItem is called when entering the groupByItem production.
	EnterGroupByItem(c *GroupByItemContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterFullColumnName is called when entering the fullColumnName production.
	EnterFullColumnName(c *FullColumnNameContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterNullNotnull is called when entering the nullNotnull production.
	EnterNullNotnull(c *NullNotnullContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterUidList is called when entering the uidList production.
	EnterUidList(c *UidListContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterSpecificFunctionCall is called when entering the specificFunctionCall production.
	EnterSpecificFunctionCall(c *SpecificFunctionCallContext)

	// EnterScalarFunctionCall is called when entering the scalarFunctionCall production.
	EnterScalarFunctionCall(c *ScalarFunctionCallContext)

	// EnterCaseVarFunctionCall is called when entering the caseVarFunctionCall production.
	EnterCaseVarFunctionCall(c *CaseVarFunctionCallContext)

	// EnterCaseFunctionCall is called when entering the caseFunctionCall production.
	EnterCaseFunctionCall(c *CaseFunctionCallContext)

	// EnterCaseFuncAlternative is called when entering the caseFuncAlternative production.
	EnterCaseFuncAlternative(c *CaseFuncAlternativeContext)

	// EnterFunctionArgs is called when entering the functionArgs production.
	EnterFunctionArgs(c *FunctionArgsContext)

	// EnterAllFunctionArg is called when entering the allFunctionArg production.
	EnterAllFunctionArg(c *AllFunctionArgContext)

	// EnterFunctionArg is called when entering the functionArg production.
	EnterFunctionArg(c *FunctionArgContext)

	// EnterIsExpression is called when entering the isExpression production.
	EnterIsExpression(c *IsExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterPredicateExpression is called when entering the predicateExpression production.
	EnterPredicateExpression(c *PredicateExpressionContext)

	// EnterExpressionAtomPredicate is called when entering the expressionAtomPredicate production.
	EnterExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// EnterInPredicate is called when entering the inPredicate production.
	EnterInPredicate(c *InPredicateContext)

	// EnterSubqueryComparasionPredicate is called when entering the subqueryComparasionPredicate production.
	EnterSubqueryComparasionPredicate(c *SubqueryComparasionPredicateContext)

	// EnterBetweenPredicate is called when entering the betweenPredicate production.
	EnterBetweenPredicate(c *BetweenPredicateContext)

	// EnterBinaryComparasionPredicate is called when entering the binaryComparasionPredicate production.
	EnterBinaryComparasionPredicate(c *BinaryComparasionPredicateContext)

	// EnterIsNullPredicate is called when entering the isNullPredicate production.
	EnterIsNullPredicate(c *IsNullPredicateContext)

	// EnterLikePredicate is called when entering the likePredicate production.
	EnterLikePredicate(c *LikePredicateContext)

	// EnterRegexpPredicate is called when entering the regexpPredicate production.
	EnterRegexpPredicate(c *RegexpPredicateContext)

	// EnterUnaryExpressionAtom is called when entering the unaryExpressionAtom production.
	EnterUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// EnterSubqueryExpressionAtom is called when entering the subqueryExpressionAtom production.
	EnterSubqueryExpressionAtom(c *SubqueryExpressionAtomContext)

	// EnterPriorityMathExpressionAtom is called when entering the priorityMathExpressionAtom production.
	EnterPriorityMathExpressionAtom(c *PriorityMathExpressionAtomContext)

	// EnterConstantExpressionAtom is called when entering the constantExpressionAtom production.
	EnterConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// EnterFunctionCallExpressionAtom is called when entering the functionCallExpressionAtom production.
	EnterFunctionCallExpressionAtom(c *FunctionCallExpressionAtomContext)

	// EnterFullColumnNameExpressionAtom is called when entering the fullColumnNameExpressionAtom production.
	EnterFullColumnNameExpressionAtom(c *FullColumnNameExpressionAtomContext)

	// EnterNestedExpressionAtom is called when entering the nestedExpressionAtom production.
	EnterNestedExpressionAtom(c *NestedExpressionAtomContext)

	// EnterMathExpressionAtom is called when entering the mathExpressionAtom production.
	EnterMathExpressionAtom(c *MathExpressionAtomContext)

	// EnterExistsExpressionAtom is called when entering the existsExpressionAtom production.
	EnterExistsExpressionAtom(c *ExistsExpressionAtomContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterLogicalOperator is called when entering the logicalOperator production.
	EnterLogicalOperator(c *LogicalOperatorContext)

	// EnterMathOperator is called when entering the mathOperator production.
	EnterMathOperator(c *MathOperatorContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterInsertStatementValue is called when entering the insertStatementValue production.
	EnterInsertStatementValue(c *InsertStatementValueContext)

	// EnterUpdatedElement is called when entering the updatedElement production.
	EnterUpdatedElement(c *UpdatedElementContext)

	// EnterExpressionsWithDefaults is called when entering the expressionsWithDefaults production.
	EnterExpressionsWithDefaults(c *ExpressionsWithDefaultsContext)

	// EnterExpressionOrDefault is called when entering the expressionOrDefault production.
	EnterExpressionOrDefault(c *ExpressionOrDefaultContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// ExitDmlStatement is called when exiting the dmlStatement production.
	ExitDmlStatement(c *DmlStatementContext)

	// ExitSimpleSelect is called when exiting the simpleSelect production.
	ExitSimpleSelect(c *SimpleSelectContext)

	// ExitParenthesisSelect is called when exiting the parenthesisSelect production.
	ExitParenthesisSelect(c *ParenthesisSelectContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByExpression is called when exiting the orderByExpression production.
	ExitOrderByExpression(c *OrderByExpressionContext)

	// ExitTableSourceBase is called when exiting the tableSourceBase production.
	ExitTableSourceBase(c *TableSourceBaseContext)

	// ExitTableSourceNested is called when exiting the tableSourceNested production.
	ExitTableSourceNested(c *TableSourceNestedContext)

	// ExitTableSources is called when exiting the tableSources production.
	ExitTableSources(c *TableSourcesContext)

	// ExitAtomTableItem is called when exiting the atomTableItem production.
	ExitAtomTableItem(c *AtomTableItemContext)

	// ExitSubqueryTableItem is called when exiting the subqueryTableItem production.
	ExitSubqueryTableItem(c *SubqueryTableItemContext)

	// ExitTableSourcesItem is called when exiting the tableSourcesItem production.
	ExitTableSourcesItem(c *TableSourcesItemContext)

	// ExitInnerJoin is called when exiting the innerJoin production.
	ExitInnerJoin(c *InnerJoinContext)

	// ExitOuterJoin is called when exiting the outerJoin production.
	ExitOuterJoin(c *OuterJoinContext)

	// ExitQueryExpression is called when exiting the queryExpression production.
	ExitQueryExpression(c *QueryExpressionContext)

	// ExitQuerySpecification is called when exiting the querySpecification production.
	ExitQuerySpecification(c *QuerySpecificationContext)

	// ExitSelectElements is called when exiting the selectElements production.
	ExitSelectElements(c *SelectElementsContext)

	// ExitSelectColumnElement is called when exiting the selectColumnElement production.
	ExitSelectColumnElement(c *SelectColumnElementContext)

	// ExitSelectFunctionElement is called when exiting the selectFunctionElement production.
	ExitSelectFunctionElement(c *SelectFunctionElementContext)

	// ExitSelectExpressionElement is called when exiting the selectExpressionElement production.
	ExitSelectExpressionElement(c *SelectExpressionElementContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitGroupByItem is called when exiting the groupByItem production.
	ExitGroupByItem(c *GroupByItemContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitFullColumnName is called when exiting the fullColumnName production.
	ExitFullColumnName(c *FullColumnNameContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitNullNotnull is called when exiting the nullNotnull production.
	ExitNullNotnull(c *NullNotnullContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitUidList is called when exiting the uidList production.
	ExitUidList(c *UidListContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitSpecificFunctionCall is called when exiting the specificFunctionCall production.
	ExitSpecificFunctionCall(c *SpecificFunctionCallContext)

	// ExitScalarFunctionCall is called when exiting the scalarFunctionCall production.
	ExitScalarFunctionCall(c *ScalarFunctionCallContext)

	// ExitCaseVarFunctionCall is called when exiting the caseVarFunctionCall production.
	ExitCaseVarFunctionCall(c *CaseVarFunctionCallContext)

	// ExitCaseFunctionCall is called when exiting the caseFunctionCall production.
	ExitCaseFunctionCall(c *CaseFunctionCallContext)

	// ExitCaseFuncAlternative is called when exiting the caseFuncAlternative production.
	ExitCaseFuncAlternative(c *CaseFuncAlternativeContext)

	// ExitFunctionArgs is called when exiting the functionArgs production.
	ExitFunctionArgs(c *FunctionArgsContext)

	// ExitAllFunctionArg is called when exiting the allFunctionArg production.
	ExitAllFunctionArg(c *AllFunctionArgContext)

	// ExitFunctionArg is called when exiting the functionArg production.
	ExitFunctionArg(c *FunctionArgContext)

	// ExitIsExpression is called when exiting the isExpression production.
	ExitIsExpression(c *IsExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitPredicateExpression is called when exiting the predicateExpression production.
	ExitPredicateExpression(c *PredicateExpressionContext)

	// ExitExpressionAtomPredicate is called when exiting the expressionAtomPredicate production.
	ExitExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// ExitInPredicate is called when exiting the inPredicate production.
	ExitInPredicate(c *InPredicateContext)

	// ExitSubqueryComparasionPredicate is called when exiting the subqueryComparasionPredicate production.
	ExitSubqueryComparasionPredicate(c *SubqueryComparasionPredicateContext)

	// ExitBetweenPredicate is called when exiting the betweenPredicate production.
	ExitBetweenPredicate(c *BetweenPredicateContext)

	// ExitBinaryComparasionPredicate is called when exiting the binaryComparasionPredicate production.
	ExitBinaryComparasionPredicate(c *BinaryComparasionPredicateContext)

	// ExitIsNullPredicate is called when exiting the isNullPredicate production.
	ExitIsNullPredicate(c *IsNullPredicateContext)

	// ExitLikePredicate is called when exiting the likePredicate production.
	ExitLikePredicate(c *LikePredicateContext)

	// ExitRegexpPredicate is called when exiting the regexpPredicate production.
	ExitRegexpPredicate(c *RegexpPredicateContext)

	// ExitUnaryExpressionAtom is called when exiting the unaryExpressionAtom production.
	ExitUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// ExitSubqueryExpressionAtom is called when exiting the subqueryExpressionAtom production.
	ExitSubqueryExpressionAtom(c *SubqueryExpressionAtomContext)

	// ExitPriorityMathExpressionAtom is called when exiting the priorityMathExpressionAtom production.
	ExitPriorityMathExpressionAtom(c *PriorityMathExpressionAtomContext)

	// ExitConstantExpressionAtom is called when exiting the constantExpressionAtom production.
	ExitConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// ExitFunctionCallExpressionAtom is called when exiting the functionCallExpressionAtom production.
	ExitFunctionCallExpressionAtom(c *FunctionCallExpressionAtomContext)

	// ExitFullColumnNameExpressionAtom is called when exiting the fullColumnNameExpressionAtom production.
	ExitFullColumnNameExpressionAtom(c *FullColumnNameExpressionAtomContext)

	// ExitNestedExpressionAtom is called when exiting the nestedExpressionAtom production.
	ExitNestedExpressionAtom(c *NestedExpressionAtomContext)

	// ExitMathExpressionAtom is called when exiting the mathExpressionAtom production.
	ExitMathExpressionAtom(c *MathExpressionAtomContext)

	// ExitExistsExpressionAtom is called when exiting the existsExpressionAtom production.
	ExitExistsExpressionAtom(c *ExistsExpressionAtomContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitLogicalOperator is called when exiting the logicalOperator production.
	ExitLogicalOperator(c *LogicalOperatorContext)

	// ExitMathOperator is called when exiting the mathOperator production.
	ExitMathOperator(c *MathOperatorContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitInsertStatementValue is called when exiting the insertStatementValue production.
	ExitInsertStatementValue(c *InsertStatementValueContext)

	// ExitUpdatedElement is called when exiting the updatedElement production.
	ExitUpdatedElement(c *UpdatedElementContext)

	// ExitExpressionsWithDefaults is called when exiting the expressionsWithDefaults production.
	ExitExpressionsWithDefaults(c *ExpressionsWithDefaultsContext)

	// ExitExpressionOrDefault is called when exiting the expressionOrDefault production.
	ExitExpressionOrDefault(c *ExpressionOrDefaultContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)
}
