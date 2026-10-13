// Code generated from Daedalus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Daedalus

import "github.com/antlr4-go/antlr/v4"

// DaedalusListener is a complete listener for a parse tree produced by DaedalusParser.
type DaedalusListener interface {
	antlr.ParseTreeListener

	// EnterDaedalusFile is called when entering the daedalusFile production.
	EnterDaedalusFile(c *DaedalusFileContext)

	// EnterBlockDef is called when entering the blockDef production.
	EnterBlockDef(c *BlockDefContext)

	// EnterInlineDef is called when entering the inlineDef production.
	EnterInlineDef(c *InlineDefContext)

	// EnterFunctionDef is called when entering the functionDef production.
	EnterFunctionDef(c *FunctionDefContext)

	// EnterConstDef is called when entering the constDef production.
	EnterConstDef(c *ConstDefContext)

	// EnterClassDef is called when entering the classDef production.
	EnterClassDef(c *ClassDefContext)

	// EnterPrototypeDef is called when entering the prototypeDef production.
	EnterPrototypeDef(c *PrototypeDefContext)

	// EnterInstanceDef is called when entering the instanceDef production.
	EnterInstanceDef(c *InstanceDefContext)

	// EnterNamespaceDef is called when entering the namespaceDef production.
	EnterNamespaceDef(c *NamespaceDefContext)

	// EnterMainBlock is called when entering the mainBlock production.
	EnterMainBlock(c *MainBlockContext)

	// EnterContentBlock is called when entering the contentBlock production.
	EnterContentBlock(c *ContentBlockContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterMetaValue is called when entering the metaValue production.
	EnterMetaValue(c *MetaValueContext)

	// EnterZParserExtenderMeta is called when entering the zParserExtenderMeta production.
	EnterZParserExtenderMeta(c *ZParserExtenderMetaContext)

	// EnterZParserExtenderMetaBlock is called when entering the zParserExtenderMetaBlock production.
	EnterZParserExtenderMetaBlock(c *ZParserExtenderMetaBlockContext)

	// EnterConstArrayDef is called when entering the constArrayDef production.
	EnterConstArrayDef(c *ConstArrayDefContext)

	// EnterConstArrayAssignment is called when entering the constArrayAssignment production.
	EnterConstArrayAssignment(c *ConstArrayAssignmentContext)

	// EnterConstValueDef is called when entering the constValueDef production.
	EnterConstValueDef(c *ConstValueDefContext)

	// EnterConstValueAssignment is called when entering the constValueAssignment production.
	EnterConstValueAssignment(c *ConstValueAssignmentContext)

	// EnterVarArrayDecl is called when entering the varArrayDecl production.
	EnterVarArrayDecl(c *VarArrayDeclContext)

	// EnterVarValueDecl is called when entering the varValueDecl production.
	EnterVarValueDecl(c *VarValueDeclContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameterDecl is called when entering the parameterDecl production.
	EnterParameterDecl(c *ParameterDeclContext)

	// EnterStatementBlock is called when entering the statementBlock production.
	EnterStatementBlock(c *StatementBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterElseIfBlock is called when entering the elseIfBlock production.
	EnterElseIfBlock(c *ElseIfBlockContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterIfBlockStatement is called when entering the ifBlockStatement production.
	EnterIfBlockStatement(c *IfBlockStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterBitMoveExpression is called when entering the bitMoveExpression production.
	EnterBitMoveExpression(c *BitMoveExpressionContext)

	// EnterReferenceExpression is called when entering the referenceExpression production.
	EnterReferenceExpression(c *ReferenceExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterBracketExpression is called when entering the bracketExpression production.
	EnterBracketExpression(c *BracketExpressionContext)

	// EnterUnaryOperation is called when entering the unaryOperation production.
	EnterUnaryOperation(c *UnaryOperationContext)

	// EnterFloatLiteralExpression is called when entering the floatLiteralExpression production.
	EnterFloatLiteralExpression(c *FloatLiteralExpressionContext)

	// EnterLogAndExpression is called when entering the logAndExpression production.
	EnterLogAndExpression(c *LogAndExpressionContext)

	// EnterFuncCallExpression is called when entering the funcCallExpression production.
	EnterFuncCallExpression(c *FuncCallExpressionContext)

	// EnterStringLiteralExpression is called when entering the stringLiteralExpression production.
	EnterStringLiteralExpression(c *StringLiteralExpressionContext)

	// EnterEqExpression is called when entering the eqExpression production.
	EnterEqExpression(c *EqExpressionContext)

	// EnterAddExpression is called when entering the addExpression production.
	EnterAddExpression(c *AddExpressionContext)

	// EnterCompExpression is called when entering the compExpression production.
	EnterCompExpression(c *CompExpressionContext)

	// EnterLogOrExpression is called when entering the logOrExpression production.
	EnterLogOrExpression(c *LogOrExpressionContext)

	// EnterBinAndExpression is called when entering the binAndExpression production.
	EnterBinAndExpression(c *BinAndExpressionContext)

	// EnterIntegerLiteralExpression is called when entering the integerLiteralExpression production.
	EnterIntegerLiteralExpression(c *IntegerLiteralExpressionContext)

	// EnterBinOrExpression is called when entering the binOrExpression production.
	EnterBinOrExpression(c *BinOrExpressionContext)

	// EnterMultExpression is called when entering the multExpression production.
	EnterMultExpression(c *MultExpressionContext)

	// EnterNullLiteralExpression is called when entering the nullLiteralExpression production.
	EnterNullLiteralExpression(c *NullLiteralExpressionContext)

	// EnterArrayIndex is called when entering the arrayIndex production.
	EnterArrayIndex(c *ArrayIndexContext)

	// EnterArraySize is called when entering the arraySize production.
	EnterArraySize(c *ArraySizeContext)

	// EnterTypeReference is called when entering the typeReference production.
	EnterTypeReference(c *TypeReferenceContext)

	// EnterNameNode is called when entering the nameNode production.
	EnterNameNode(c *NameNodeContext)

	// EnterParentReference is called when entering the parentReference production.
	EnterParentReference(c *ParentReferenceContext)

	// ExitDaedalusFile is called when exiting the daedalusFile production.
	ExitDaedalusFile(c *DaedalusFileContext)

	// ExitBlockDef is called when exiting the blockDef production.
	ExitBlockDef(c *BlockDefContext)

	// ExitInlineDef is called when exiting the inlineDef production.
	ExitInlineDef(c *InlineDefContext)

	// ExitFunctionDef is called when exiting the functionDef production.
	ExitFunctionDef(c *FunctionDefContext)

	// ExitConstDef is called when exiting the constDef production.
	ExitConstDef(c *ConstDefContext)

	// ExitClassDef is called when exiting the classDef production.
	ExitClassDef(c *ClassDefContext)

	// ExitPrototypeDef is called when exiting the prototypeDef production.
	ExitPrototypeDef(c *PrototypeDefContext)

	// ExitInstanceDef is called when exiting the instanceDef production.
	ExitInstanceDef(c *InstanceDefContext)

	// ExitNamespaceDef is called when exiting the namespaceDef production.
	ExitNamespaceDef(c *NamespaceDefContext)

	// ExitMainBlock is called when exiting the mainBlock production.
	ExitMainBlock(c *MainBlockContext)

	// ExitContentBlock is called when exiting the contentBlock production.
	ExitContentBlock(c *ContentBlockContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitMetaValue is called when exiting the metaValue production.
	ExitMetaValue(c *MetaValueContext)

	// ExitZParserExtenderMeta is called when exiting the zParserExtenderMeta production.
	ExitZParserExtenderMeta(c *ZParserExtenderMetaContext)

	// ExitZParserExtenderMetaBlock is called when exiting the zParserExtenderMetaBlock production.
	ExitZParserExtenderMetaBlock(c *ZParserExtenderMetaBlockContext)

	// ExitConstArrayDef is called when exiting the constArrayDef production.
	ExitConstArrayDef(c *ConstArrayDefContext)

	// ExitConstArrayAssignment is called when exiting the constArrayAssignment production.
	ExitConstArrayAssignment(c *ConstArrayAssignmentContext)

	// ExitConstValueDef is called when exiting the constValueDef production.
	ExitConstValueDef(c *ConstValueDefContext)

	// ExitConstValueAssignment is called when exiting the constValueAssignment production.
	ExitConstValueAssignment(c *ConstValueAssignmentContext)

	// ExitVarArrayDecl is called when exiting the varArrayDecl production.
	ExitVarArrayDecl(c *VarArrayDeclContext)

	// ExitVarValueDecl is called when exiting the varValueDecl production.
	ExitVarValueDecl(c *VarValueDeclContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameterDecl is called when exiting the parameterDecl production.
	ExitParameterDecl(c *ParameterDeclContext)

	// ExitStatementBlock is called when exiting the statementBlock production.
	ExitStatementBlock(c *StatementBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitElseIfBlock is called when exiting the elseIfBlock production.
	ExitElseIfBlock(c *ElseIfBlockContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitIfBlockStatement is called when exiting the ifBlockStatement production.
	ExitIfBlockStatement(c *IfBlockStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitBitMoveExpression is called when exiting the bitMoveExpression production.
	ExitBitMoveExpression(c *BitMoveExpressionContext)

	// ExitReferenceExpression is called when exiting the referenceExpression production.
	ExitReferenceExpression(c *ReferenceExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitBracketExpression is called when exiting the bracketExpression production.
	ExitBracketExpression(c *BracketExpressionContext)

	// ExitUnaryOperation is called when exiting the unaryOperation production.
	ExitUnaryOperation(c *UnaryOperationContext)

	// ExitFloatLiteralExpression is called when exiting the floatLiteralExpression production.
	ExitFloatLiteralExpression(c *FloatLiteralExpressionContext)

	// ExitLogAndExpression is called when exiting the logAndExpression production.
	ExitLogAndExpression(c *LogAndExpressionContext)

	// ExitFuncCallExpression is called when exiting the funcCallExpression production.
	ExitFuncCallExpression(c *FuncCallExpressionContext)

	// ExitStringLiteralExpression is called when exiting the stringLiteralExpression production.
	ExitStringLiteralExpression(c *StringLiteralExpressionContext)

	// ExitEqExpression is called when exiting the eqExpression production.
	ExitEqExpression(c *EqExpressionContext)

	// ExitAddExpression is called when exiting the addExpression production.
	ExitAddExpression(c *AddExpressionContext)

	// ExitCompExpression is called when exiting the compExpression production.
	ExitCompExpression(c *CompExpressionContext)

	// ExitLogOrExpression is called when exiting the logOrExpression production.
	ExitLogOrExpression(c *LogOrExpressionContext)

	// ExitBinAndExpression is called when exiting the binAndExpression production.
	ExitBinAndExpression(c *BinAndExpressionContext)

	// ExitIntegerLiteralExpression is called when exiting the integerLiteralExpression production.
	ExitIntegerLiteralExpression(c *IntegerLiteralExpressionContext)

	// ExitBinOrExpression is called when exiting the binOrExpression production.
	ExitBinOrExpression(c *BinOrExpressionContext)

	// ExitMultExpression is called when exiting the multExpression production.
	ExitMultExpression(c *MultExpressionContext)

	// ExitNullLiteralExpression is called when exiting the nullLiteralExpression production.
	ExitNullLiteralExpression(c *NullLiteralExpressionContext)

	// ExitArrayIndex is called when exiting the arrayIndex production.
	ExitArrayIndex(c *ArrayIndexContext)

	// ExitArraySize is called when exiting the arraySize production.
	ExitArraySize(c *ArraySizeContext)

	// ExitTypeReference is called when exiting the typeReference production.
	ExitTypeReference(c *TypeReferenceContext)

	// ExitNameNode is called when exiting the nameNode production.
	ExitNameNode(c *NameNodeContext)

	// ExitParentReference is called when exiting the parentReference production.
	ExitParentReference(c *ParentReferenceContext)
}
