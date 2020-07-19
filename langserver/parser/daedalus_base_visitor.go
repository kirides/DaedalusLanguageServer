// Code generated from Daedalus.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Daedalus

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDaedalusVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDaedalusVisitor) VisitSymbolSummary(ctx *SymbolSummaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitDaedalusFile(ctx *DaedalusFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitBlockDef(ctx *BlockDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitInlineDef(ctx *InlineDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitFunctionDef(ctx *FunctionDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitConstDef(ctx *ConstDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitClassDef(ctx *ClassDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitPrototypeDef(ctx *PrototypeDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitInstanceDef(ctx *InstanceDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitInstanceDecl(ctx *InstanceDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitConstArrayDef(ctx *ConstArrayDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitConstArrayAssignment(ctx *ConstArrayAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitConstValueDef(ctx *ConstValueDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitConstValueAssignment(ctx *ConstValueAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitVarArrayDecl(ctx *VarArrayDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitVarValueDecl(ctx *VarValueDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitParameterDecl(ctx *ParameterDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitStatementBlock(ctx *StatementBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitIfCondition(ctx *IfConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitElseBlock(ctx *ElseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitElseIfBlock(ctx *ElseIfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitIfBlock(ctx *IfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitIfBlockStatement(ctx *IfBlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitFuncArgExpression(ctx *FuncArgExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitExpressionBlock(ctx *ExpressionBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitBitMoveExpression(ctx *BitMoveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitOneArgExpression(ctx *OneArgExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitEqExpression(ctx *EqExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitValExpression(ctx *ValExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitAddExpression(ctx *AddExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitCompExpression(ctx *CompExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitLogOrExpression(ctx *LogOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitBinAndExpression(ctx *BinAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitBinOrExpression(ctx *BinOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitMultExpression(ctx *MultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitBracketExpression(ctx *BracketExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitLogAndExpression(ctx *LogAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitArrayIndex(ctx *ArrayIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitArraySize(ctx *ArraySizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitIntegerLiteralValue(ctx *IntegerLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitFloatLiteralValue(ctx *FloatLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitStringLiteralValue(ctx *StringLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitNullLiteralValue(ctx *NullLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitNoFuncLiteralValue(ctx *NoFuncLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitFuncCallValue(ctx *FuncCallValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitReferenceValue(ctx *ReferenceValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitAnyIdentifierValue(ctx *AnyIdentifierValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitReferenceAtom(ctx *ReferenceAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitReference(ctx *ReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitTypeReference(ctx *TypeReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitAnyIdentifier(ctx *AnyIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitNameNode(ctx *NameNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitParentReference(ctx *ParentReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitAddOperator(ctx *AddOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitBitMoveOperator(ctx *BitMoveOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitCompOperator(ctx *CompOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitEqOperator(ctx *EqOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitOneArgOperator(ctx *OneArgOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitMultOperator(ctx *MultOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitBinAndOperator(ctx *BinAndOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitBinOrOperator(ctx *BinOrOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitLogAndOperator(ctx *LogAndOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDaedalusVisitor) VisitLogOrOperator(ctx *LogOrOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}
