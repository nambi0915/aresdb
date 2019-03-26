// Code generated from query/sql/SqlBase.g4 by ANTLR 4.7.1. DO NOT EDIT.

package antlrgen // SqlBase
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSqlBaseVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSqlBaseVisitor) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSingleExpression(ctx *SingleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitStatementDefault(ctx *StatementDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitQueryNoWith(ctx *QueryNoWithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitQueryTermDefault(ctx *QueryTermDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitGroupBy(ctx *GroupByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitGroupingExpressions(ctx *GroupingExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSelectSingle(ctx *SelectSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSelectAll(ctx *SelectAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitRelationDefault(ctx *RelationDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitJoinType(ctx *JoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitJoinCriteria(ctx *JoinCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSampledRelation(ctx *SampledRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSampleType(ctx *SampleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitAliasedRelation(ctx *AliasedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSubqueryRelation(ctx *SubqueryRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitUnnest(ctx *UnnestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitLateral(ctx *LateralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitBooleanDefault(ctx *BooleanDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitPredicated(ctx *PredicatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitQuantifiedComparison(ctx *QuantifiedComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitBetween(ctx *BetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitInList(ctx *InListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitInSubquery(ctx *InSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitNullPredicate(ctx *NullPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitDistinctFrom(ctx *DistinctFromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitAtTimeZone(ctx *AtTimeZoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitDereference(ctx *DereferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitTypeConstructor(ctx *TypeConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSpecialDateTimeFunction(ctx *SpecialDateTimeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSubstring(ctx *SubstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitCast(ctx *CastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitNormalize(ctx *NormalizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSimpleCase(ctx *SimpleCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSubscript(ctx *SubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitCurrentUser(ctx *CurrentUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitExtract(ctx *ExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitArrayConstructor(ctx *ArrayConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitExists(ctx *ExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSearchedCase(ctx *SearchedCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitGroupingOperation(ctx *GroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitBasicStringLiteral(ctx *BasicStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitUnicodeStringLiteral(ctx *UnicodeStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitTimeZoneInterval(ctx *TimeZoneIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitTimeZoneString(ctx *TimeZoneStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitComparisonQuantifier(ctx *ComparisonQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitIntervalField(ctx *IntervalFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitNormalForm(ctx *NormalFormContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSqltype(ctx *SqltypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitBaseType(ctx *BaseTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitWhenClause(ctx *WhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitFilter(ctx *FilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitOver(ctx *OverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitWindowFrame(ctx *WindowFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitUnboundedFrame(ctx *UnboundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitCurrentRowBound(ctx *CurrentRowBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitBoundedFrame(ctx *BoundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitExplainFormat(ctx *ExplainFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitExplainType(ctx *ExplainTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitIsolationLevel(ctx *IsolationLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitReadUncommitted(ctx *ReadUncommittedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitReadCommitted(ctx *ReadCommittedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitRepeatableRead(ctx *RepeatableReadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitSerializable(ctx *SerializableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitPositionalArgument(ctx *PositionalArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitNamedArgument(ctx *NamedArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitPrivilege(ctx *PrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitQuotedIdentifier(ctx *QuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitDoubleLiteral(ctx *DoubleLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlBaseVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}