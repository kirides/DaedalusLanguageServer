// Code generated from Daedalus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Daedalus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DaedalusListener is a complete listener for a parse tree produced by DaedalusParser.
type DaedalusListener interface {
	antlr.ParseTreeListener

	// EnterSymbolSummary is called when entering the symbolSummary production.
	EnterSymbolSummary(c *SymbolSummaryContext)

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

	// EnterInstanceDecl is called when entering the instanceDecl production.
	EnterInstanceDecl(c *InstanceDeclContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

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

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterIfCondition is called when entering the ifCondition production.
	EnterIfCondition(c *IfConditionContext)

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

	// EnterFuncArgExpression is called when entering the funcArgExpression production.
	EnterFuncArgExpression(c *FuncArgExpressionContext)

	// EnterExpressionBlock is called when entering the expressionBlock production.
	EnterExpressionBlock(c *ExpressionBlockContext)

	// EnterBitMoveExpression is called when entering the bitMoveExpression production.
	EnterBitMoveExpression(c *BitMoveExpressionContext)

	// EnterOneArgExpression is called when entering the oneArgExpression production.
	EnterOneArgExpression(c *OneArgExpressionContext)

	// EnterEqExpression is called when entering the eqExpression production.
	EnterEqExpression(c *EqExpressionContext)

	// EnterValExpression is called when entering the valExpression production.
	EnterValExpression(c *ValExpressionContext)

	// EnterAddExpression is called when entering the addExpression production.
	EnterAddExpression(c *AddExpressionContext)

	// EnterCompExpression is called when entering the compExpression production.
	EnterCompExpression(c *CompExpressionContext)

	// EnterLogOrExpression is called when entering the logOrExpression production.
	EnterLogOrExpression(c *LogOrExpressionContext)

	// EnterBinAndExpression is called when entering the binAndExpression production.
	EnterBinAndExpression(c *BinAndExpressionContext)

	// EnterBinOrExpression is called when entering the binOrExpression production.
	EnterBinOrExpression(c *BinOrExpressionContext)

	// EnterMultExpression is called when entering the multExpression production.
	EnterMultExpression(c *MultExpressionContext)

	// EnterBracketExpression is called when entering the bracketExpression production.
	EnterBracketExpression(c *BracketExpressionContext)

	// EnterLogAndExpression is called when entering the logAndExpression production.
	EnterLogAndExpression(c *LogAndExpressionContext)

	// EnterArrayIndex is called when entering the arrayIndex production.
	EnterArrayIndex(c *ArrayIndexContext)

	// EnterArraySize is called when entering the arraySize production.
	EnterArraySize(c *ArraySizeContext)

	// EnterIntegerLiteralValue is called when entering the integerLiteralValue production.
	EnterIntegerLiteralValue(c *IntegerLiteralValueContext)

	// EnterFloatLiteralValue is called when entering the floatLiteralValue production.
	EnterFloatLiteralValue(c *FloatLiteralValueContext)

	// EnterStringLiteralValue is called when entering the stringLiteralValue production.
	EnterStringLiteralValue(c *StringLiteralValueContext)

	// EnterNullLiteralValue is called when entering the nullLiteralValue production.
	EnterNullLiteralValue(c *NullLiteralValueContext)

	// EnterNoFuncLiteralValue is called when entering the noFuncLiteralValue production.
	EnterNoFuncLiteralValue(c *NoFuncLiteralValueContext)

	// EnterFuncCallValue is called when entering the funcCallValue production.
	EnterFuncCallValue(c *FuncCallValueContext)

	// EnterReferenceValue is called when entering the referenceValue production.
	EnterReferenceValue(c *ReferenceValueContext)

	// EnterAnyIdentifierValue is called when entering the anyIdentifierValue production.
	EnterAnyIdentifierValue(c *AnyIdentifierValueContext)

	// EnterReferenceAtom is called when entering the referenceAtom production.
	EnterReferenceAtom(c *ReferenceAtomContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterTypeReference is called when entering the typeReference production.
	EnterTypeReference(c *TypeReferenceContext)

	// EnterAnyIdentifier is called when entering the anyIdentifier production.
	EnterAnyIdentifier(c *AnyIdentifierContext)

	// EnterNameNode is called when entering the nameNode production.
	EnterNameNode(c *NameNodeContext)

	// EnterParentReference is called when entering the parentReference production.
	EnterParentReference(c *ParentReferenceContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterAddOperator is called when entering the addOperator production.
	EnterAddOperator(c *AddOperatorContext)

	// EnterBitMoveOperator is called when entering the bitMoveOperator production.
	EnterBitMoveOperator(c *BitMoveOperatorContext)

	// EnterCompOperator is called when entering the compOperator production.
	EnterCompOperator(c *CompOperatorContext)

	// EnterEqOperator is called when entering the eqOperator production.
	EnterEqOperator(c *EqOperatorContext)

	// EnterOneArgOperator is called when entering the oneArgOperator production.
	EnterOneArgOperator(c *OneArgOperatorContext)

	// EnterMultOperator is called when entering the multOperator production.
	EnterMultOperator(c *MultOperatorContext)

	// EnterBinAndOperator is called when entering the binAndOperator production.
	EnterBinAndOperator(c *BinAndOperatorContext)

	// EnterBinOrOperator is called when entering the binOrOperator production.
	EnterBinOrOperator(c *BinOrOperatorContext)

	// EnterLogAndOperator is called when entering the logAndOperator production.
	EnterLogAndOperator(c *LogAndOperatorContext)

	// EnterLogOrOperator is called when entering the logOrOperator production.
	EnterLogOrOperator(c *LogOrOperatorContext)

	// ExitSymbolSummary is called when exiting the symbolSummary production.
	ExitSymbolSummary(c *SymbolSummaryContext)

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

	// ExitInstanceDecl is called when exiting the instanceDecl production.
	ExitInstanceDecl(c *InstanceDeclContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

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

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitIfCondition is called when exiting the ifCondition production.
	ExitIfCondition(c *IfConditionContext)

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

	// ExitFuncArgExpression is called when exiting the funcArgExpression production.
	ExitFuncArgExpression(c *FuncArgExpressionContext)

	// ExitExpressionBlock is called when exiting the expressionBlock production.
	ExitExpressionBlock(c *ExpressionBlockContext)

	// ExitBitMoveExpression is called when exiting the bitMoveExpression production.
	ExitBitMoveExpression(c *BitMoveExpressionContext)

	// ExitOneArgExpression is called when exiting the oneArgExpression production.
	ExitOneArgExpression(c *OneArgExpressionContext)

	// ExitEqExpression is called when exiting the eqExpression production.
	ExitEqExpression(c *EqExpressionContext)

	// ExitValExpression is called when exiting the valExpression production.
	ExitValExpression(c *ValExpressionContext)

	// ExitAddExpression is called when exiting the addExpression production.
	ExitAddExpression(c *AddExpressionContext)

	// ExitCompExpression is called when exiting the compExpression production.
	ExitCompExpression(c *CompExpressionContext)

	// ExitLogOrExpression is called when exiting the logOrExpression production.
	ExitLogOrExpression(c *LogOrExpressionContext)

	// ExitBinAndExpression is called when exiting the binAndExpression production.
	ExitBinAndExpression(c *BinAndExpressionContext)

	// ExitBinOrExpression is called when exiting the binOrExpression production.
	ExitBinOrExpression(c *BinOrExpressionContext)

	// ExitMultExpression is called when exiting the multExpression production.
	ExitMultExpression(c *MultExpressionContext)

	// ExitBracketExpression is called when exiting the bracketExpression production.
	ExitBracketExpression(c *BracketExpressionContext)

	// ExitLogAndExpression is called when exiting the logAndExpression production.
	ExitLogAndExpression(c *LogAndExpressionContext)

	// ExitArrayIndex is called when exiting the arrayIndex production.
	ExitArrayIndex(c *ArrayIndexContext)

	// ExitArraySize is called when exiting the arraySize production.
	ExitArraySize(c *ArraySizeContext)

	// ExitIntegerLiteralValue is called when exiting the integerLiteralValue production.
	ExitIntegerLiteralValue(c *IntegerLiteralValueContext)

	// ExitFloatLiteralValue is called when exiting the floatLiteralValue production.
	ExitFloatLiteralValue(c *FloatLiteralValueContext)

	// ExitStringLiteralValue is called when exiting the stringLiteralValue production.
	ExitStringLiteralValue(c *StringLiteralValueContext)

	// ExitNullLiteralValue is called when exiting the nullLiteralValue production.
	ExitNullLiteralValue(c *NullLiteralValueContext)

	// ExitNoFuncLiteralValue is called when exiting the noFuncLiteralValue production.
	ExitNoFuncLiteralValue(c *NoFuncLiteralValueContext)

	// ExitFuncCallValue is called when exiting the funcCallValue production.
	ExitFuncCallValue(c *FuncCallValueContext)

	// ExitReferenceValue is called when exiting the referenceValue production.
	ExitReferenceValue(c *ReferenceValueContext)

	// ExitAnyIdentifierValue is called when exiting the anyIdentifierValue production.
	ExitAnyIdentifierValue(c *AnyIdentifierValueContext)

	// ExitReferenceAtom is called when exiting the referenceAtom production.
	ExitReferenceAtom(c *ReferenceAtomContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitTypeReference is called when exiting the typeReference production.
	ExitTypeReference(c *TypeReferenceContext)

	// ExitAnyIdentifier is called when exiting the anyIdentifier production.
	ExitAnyIdentifier(c *AnyIdentifierContext)

	// ExitNameNode is called when exiting the nameNode production.
	ExitNameNode(c *NameNodeContext)

	// ExitParentReference is called when exiting the parentReference production.
	ExitParentReference(c *ParentReferenceContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitAddOperator is called when exiting the addOperator production.
	ExitAddOperator(c *AddOperatorContext)

	// ExitBitMoveOperator is called when exiting the bitMoveOperator production.
	ExitBitMoveOperator(c *BitMoveOperatorContext)

	// ExitCompOperator is called when exiting the compOperator production.
	ExitCompOperator(c *CompOperatorContext)

	// ExitEqOperator is called when exiting the eqOperator production.
	ExitEqOperator(c *EqOperatorContext)

	// ExitOneArgOperator is called when exiting the oneArgOperator production.
	ExitOneArgOperator(c *OneArgOperatorContext)

	// ExitMultOperator is called when exiting the multOperator production.
	ExitMultOperator(c *MultOperatorContext)

	// ExitBinAndOperator is called when exiting the binAndOperator production.
	ExitBinAndOperator(c *BinAndOperatorContext)

	// ExitBinOrOperator is called when exiting the binOrOperator production.
	ExitBinOrOperator(c *BinOrOperatorContext)

	// ExitLogAndOperator is called when exiting the logAndOperator production.
	ExitLogAndOperator(c *LogAndOperatorContext)

	// ExitLogOrOperator is called when exiting the logOrOperator production.
	ExitLogOrOperator(c *LogOrOperatorContext)
}
