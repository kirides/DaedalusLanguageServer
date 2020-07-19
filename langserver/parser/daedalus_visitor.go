// Code generated from Daedalus.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Daedalus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DaedalusParser.
type DaedalusVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DaedalusParser#symbolSummary.
	VisitSymbolSummary(ctx *SymbolSummaryContext) interface{}

	// Visit a parse tree produced by DaedalusParser#daedalusFile.
	VisitDaedalusFile(ctx *DaedalusFileContext) interface{}

	// Visit a parse tree produced by DaedalusParser#blockDef.
	VisitBlockDef(ctx *BlockDefContext) interface{}

	// Visit a parse tree produced by DaedalusParser#inlineDef.
	VisitInlineDef(ctx *InlineDefContext) interface{}

	// Visit a parse tree produced by DaedalusParser#functionDef.
	VisitFunctionDef(ctx *FunctionDefContext) interface{}

	// Visit a parse tree produced by DaedalusParser#constDef.
	VisitConstDef(ctx *ConstDefContext) interface{}

	// Visit a parse tree produced by DaedalusParser#classDef.
	VisitClassDef(ctx *ClassDefContext) interface{}

	// Visit a parse tree produced by DaedalusParser#prototypeDef.
	VisitPrototypeDef(ctx *PrototypeDefContext) interface{}

	// Visit a parse tree produced by DaedalusParser#instanceDef.
	VisitInstanceDef(ctx *InstanceDefContext) interface{}

	// Visit a parse tree produced by DaedalusParser#instanceDecl.
	VisitInstanceDecl(ctx *InstanceDeclContext) interface{}

	// Visit a parse tree produced by DaedalusParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by DaedalusParser#constArrayDef.
	VisitConstArrayDef(ctx *ConstArrayDefContext) interface{}

	// Visit a parse tree produced by DaedalusParser#constArrayAssignment.
	VisitConstArrayAssignment(ctx *ConstArrayAssignmentContext) interface{}

	// Visit a parse tree produced by DaedalusParser#constValueDef.
	VisitConstValueDef(ctx *ConstValueDefContext) interface{}

	// Visit a parse tree produced by DaedalusParser#constValueAssignment.
	VisitConstValueAssignment(ctx *ConstValueAssignmentContext) interface{}

	// Visit a parse tree produced by DaedalusParser#varArrayDecl.
	VisitVarArrayDecl(ctx *VarArrayDeclContext) interface{}

	// Visit a parse tree produced by DaedalusParser#varValueDecl.
	VisitVarValueDecl(ctx *VarValueDeclContext) interface{}

	// Visit a parse tree produced by DaedalusParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by DaedalusParser#parameterDecl.
	VisitParameterDecl(ctx *ParameterDeclContext) interface{}

	// Visit a parse tree produced by DaedalusParser#statementBlock.
	VisitStatementBlock(ctx *StatementBlockContext) interface{}

	// Visit a parse tree produced by DaedalusParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by DaedalusParser#funcCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by DaedalusParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by DaedalusParser#ifCondition.
	VisitIfCondition(ctx *IfConditionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#elseBlock.
	VisitElseBlock(ctx *ElseBlockContext) interface{}

	// Visit a parse tree produced by DaedalusParser#elseIfBlock.
	VisitElseIfBlock(ctx *ElseIfBlockContext) interface{}

	// Visit a parse tree produced by DaedalusParser#ifBlock.
	VisitIfBlock(ctx *IfBlockContext) interface{}

	// Visit a parse tree produced by DaedalusParser#ifBlockStatement.
	VisitIfBlockStatement(ctx *IfBlockStatementContext) interface{}

	// Visit a parse tree produced by DaedalusParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by DaedalusParser#funcArgExpression.
	VisitFuncArgExpression(ctx *FuncArgExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#expressionBlock.
	VisitExpressionBlock(ctx *ExpressionBlockContext) interface{}

	// Visit a parse tree produced by DaedalusParser#bitMoveExpression.
	VisitBitMoveExpression(ctx *BitMoveExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#oneArgExpression.
	VisitOneArgExpression(ctx *OneArgExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#eqExpression.
	VisitEqExpression(ctx *EqExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#valExpression.
	VisitValExpression(ctx *ValExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#addExpression.
	VisitAddExpression(ctx *AddExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#compExpression.
	VisitCompExpression(ctx *CompExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#logOrExpression.
	VisitLogOrExpression(ctx *LogOrExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#binAndExpression.
	VisitBinAndExpression(ctx *BinAndExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#binOrExpression.
	VisitBinOrExpression(ctx *BinOrExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#multExpression.
	VisitMultExpression(ctx *MultExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#bracketExpression.
	VisitBracketExpression(ctx *BracketExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#logAndExpression.
	VisitLogAndExpression(ctx *LogAndExpressionContext) interface{}

	// Visit a parse tree produced by DaedalusParser#arrayIndex.
	VisitArrayIndex(ctx *ArrayIndexContext) interface{}

	// Visit a parse tree produced by DaedalusParser#arraySize.
	VisitArraySize(ctx *ArraySizeContext) interface{}

	// Visit a parse tree produced by DaedalusParser#integerLiteralValue.
	VisitIntegerLiteralValue(ctx *IntegerLiteralValueContext) interface{}

	// Visit a parse tree produced by DaedalusParser#floatLiteralValue.
	VisitFloatLiteralValue(ctx *FloatLiteralValueContext) interface{}

	// Visit a parse tree produced by DaedalusParser#stringLiteralValue.
	VisitStringLiteralValue(ctx *StringLiteralValueContext) interface{}

	// Visit a parse tree produced by DaedalusParser#nullLiteralValue.
	VisitNullLiteralValue(ctx *NullLiteralValueContext) interface{}

	// Visit a parse tree produced by DaedalusParser#noFuncLiteralValue.
	VisitNoFuncLiteralValue(ctx *NoFuncLiteralValueContext) interface{}

	// Visit a parse tree produced by DaedalusParser#funcCallValue.
	VisitFuncCallValue(ctx *FuncCallValueContext) interface{}

	// Visit a parse tree produced by DaedalusParser#referenceValue.
	VisitReferenceValue(ctx *ReferenceValueContext) interface{}

	// Visit a parse tree produced by DaedalusParser#anyIdentifierValue.
	VisitAnyIdentifierValue(ctx *AnyIdentifierValueContext) interface{}

	// Visit a parse tree produced by DaedalusParser#referenceAtom.
	VisitReferenceAtom(ctx *ReferenceAtomContext) interface{}

	// Visit a parse tree produced by DaedalusParser#reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by DaedalusParser#typeReference.
	VisitTypeReference(ctx *TypeReferenceContext) interface{}

	// Visit a parse tree produced by DaedalusParser#anyIdentifier.
	VisitAnyIdentifier(ctx *AnyIdentifierContext) interface{}

	// Visit a parse tree produced by DaedalusParser#nameNode.
	VisitNameNode(ctx *NameNodeContext) interface{}

	// Visit a parse tree produced by DaedalusParser#parentReference.
	VisitParentReference(ctx *ParentReferenceContext) interface{}

	// Visit a parse tree produced by DaedalusParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#addOperator.
	VisitAddOperator(ctx *AddOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#bitMoveOperator.
	VisitBitMoveOperator(ctx *BitMoveOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#compOperator.
	VisitCompOperator(ctx *CompOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#eqOperator.
	VisitEqOperator(ctx *EqOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#oneArgOperator.
	VisitOneArgOperator(ctx *OneArgOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#multOperator.
	VisitMultOperator(ctx *MultOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#binAndOperator.
	VisitBinAndOperator(ctx *BinAndOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#binOrOperator.
	VisitBinOrOperator(ctx *BinOrOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#logAndOperator.
	VisitLogAndOperator(ctx *LogAndOperatorContext) interface{}

	// Visit a parse tree produced by DaedalusParser#logOrOperator.
	VisitLogOrOperator(ctx *LogOrOperatorContext) interface{}
}
