package langserver

/*
	Regex to generate this:
	find: (func \()(\*)(.+?)(\) )(.+?)(\(\w+)(.+? {\n\t)(?:panic\("unimplemented"\))(\n})
	replace: $1l $2$3$4$5$6$7l.left.$5$6)\n\tl.right.$5$6)$8

	Example: https://regex101.com/r/2x7yoy/1
*/

import (
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"

	"github.com/antlr4-go/antlr/v4"
)

func combineListeners(listeners ...parser.DaedalusListener) parser.DaedalusListener {
	if len(listeners) == 1 {
		return listeners[0]
	}

	var current parser.DaedalusListener
	for i := 0; i < len(listeners); i++ {
		if listeners[i] == nil {
			continue
		}
		if current == nil {
			current = listeners[i]
		} else {
			current = &CombinedDaedalusListener{left: current, right: listeners[i]}
		}
	}
	return current
}

type CombinedDaedalusListener struct {
	// parser.BaseDaedalusListener

	left, right parser.DaedalusListener
}

// EnterZParserExtenderMeta implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterZParserExtenderMeta(c *parser.ZParserExtenderMetaContext) {
	l.left.EnterZParserExtenderMeta(c)
	l.right.EnterZParserExtenderMeta(c)
}

// ExitZParserExtenderMeta implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitZParserExtenderMeta(c *parser.ZParserExtenderMetaContext) {
	l.left.ExitZParserExtenderMeta(c)
	l.right.ExitZParserExtenderMeta(c)
}

// EnterContentBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterContentBlock(c *parser.ContentBlockContext) {
	l.left.EnterContentBlock(c)
	l.right.EnterContentBlock(c)
}

// ExitContentBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitContentBlock(c *parser.ContentBlockContext) {
	l.left.ExitContentBlock(c)
	l.right.ExitContentBlock(c)
}

// EnterMetaValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterMetaValue(c *parser.MetaValueContext) {
	l.left.EnterMetaValue(c)
	l.right.EnterMetaValue(c)
}

// ExitMetaValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitMetaValue(c *parser.MetaValueContext) {
	l.left.ExitMetaValue(c)
	l.right.ExitMetaValue(c)
}

// EnterZParserExtenderMetaBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterZParserExtenderMetaBlock(c *parser.ZParserExtenderMetaBlockContext) {
	l.left.EnterZParserExtenderMetaBlock(c)
	l.right.EnterZParserExtenderMetaBlock(c)
}

// ExitZParserExtenderMetaBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitZParserExtenderMetaBlock(c *parser.ZParserExtenderMetaBlockContext) {
	l.left.ExitZParserExtenderMetaBlock(c)
	l.right.ExitZParserExtenderMetaBlock(c)
}

// EnterMainBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterMainBlock(c *parser.MainBlockContext) {
	l.left.EnterMainBlock(c)
	l.right.EnterMainBlock(c)
}

// ExitMainBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitMainBlock(c *parser.MainBlockContext) {
	l.left.ExitMainBlock(c)
	l.right.ExitMainBlock(c)
}

// EnterNamespaceDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterNamespaceDef(c *parser.NamespaceDefContext) {
	l.left.EnterNamespaceDef(c)
	l.right.EnterNamespaceDef(c)
}

// ExitNamespaceDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitNamespaceDef(c *parser.NamespaceDefContext) {
	l.left.ExitNamespaceDef(c)
	l.right.ExitNamespaceDef(c)
}

// EnterUnaryOperation implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterUnaryOperation(c *parser.UnaryOperationContext) {
	l.left.EnterUnaryOperation(c)
	l.right.EnterUnaryOperation(c)
}

// EnterUnaryOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterUnaryOperator(c *parser.UnaryOperatorContext) {
	l.left.EnterUnaryOperator(c)
	l.right.EnterUnaryOperator(c)
}

// ExitUnaryOperation implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitUnaryOperation(c *parser.UnaryOperationContext) {
	l.left.ExitUnaryOperation(c)
	l.right.ExitUnaryOperation(c)
}

// ExitUnaryOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitUnaryOperator(c *parser.UnaryOperatorContext) {
	l.left.ExitUnaryOperator(c)
	l.right.ExitUnaryOperator(c)
}

// EnterEveryRule implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	l.left.EnterEveryRule(ctx)
	l.right.EnterEveryRule(ctx)
}

// ExitEveryRule implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	l.left.ExitEveryRule(ctx)
	l.right.ExitEveryRule(ctx)
}

// VisitErrorNode implements parser.DaedalusListener
func (l *CombinedDaedalusListener) VisitErrorNode(node antlr.ErrorNode) {
	l.left.VisitErrorNode(node)
	l.right.VisitErrorNode(node)
}

// VisitTerminal implements parser.DaedalusListener
func (l *CombinedDaedalusListener) VisitTerminal(node antlr.TerminalNode) {
	l.left.VisitTerminal(node)
	l.right.VisitTerminal(node)
}

// EnterAddExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterAddExpression(c *parser.AddExpressionContext) {
	l.left.EnterAddExpression(c)
	l.right.EnterAddExpression(c)
}

// EnterAddOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterAddOperator(c *parser.AddOperatorContext) {
	l.left.EnterAddOperator(c)
	l.right.EnterAddOperator(c)
}

// EnterAnyIdentifier implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterAnyIdentifier(c *parser.AnyIdentifierContext) {
	l.left.EnterAnyIdentifier(c)
	l.right.EnterAnyIdentifier(c)
}

// EnterArrayIndex implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterArrayIndex(c *parser.ArrayIndexContext) {
	l.left.EnterArrayIndex(c)
	l.right.EnterArrayIndex(c)
}

// EnterArraySize implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterArraySize(c *parser.ArraySizeContext) {
	l.left.EnterArraySize(c)
	l.right.EnterArraySize(c)
}

// EnterAssignment implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterAssignment(c *parser.AssignmentContext) {
	l.left.EnterAssignment(c)
	l.right.EnterAssignment(c)
}

// EnterAssignmentOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterAssignmentOperator(c *parser.AssignmentOperatorContext) {
	l.left.EnterAssignmentOperator(c)
	l.right.EnterAssignmentOperator(c)
}

// EnterBinAndExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterBinAndExpression(c *parser.BinAndExpressionContext) {
	l.left.EnterBinAndExpression(c)
	l.right.EnterBinAndExpression(c)
}

// EnterBinAndOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterBinAndOperator(c *parser.BinAndOperatorContext) {
	l.left.EnterBinAndOperator(c)
	l.right.EnterBinAndOperator(c)
}

// EnterBinOrExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterBinOrExpression(c *parser.BinOrExpressionContext) {
	l.left.EnterBinOrExpression(c)
	l.right.EnterBinOrExpression(c)
}

// EnterBinOrOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterBinOrOperator(c *parser.BinOrOperatorContext) {
	l.left.EnterBinOrOperator(c)
	l.right.EnterBinOrOperator(c)
}

// EnterBitMoveExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterBitMoveExpression(c *parser.BitMoveExpressionContext) {
	l.left.EnterBitMoveExpression(c)
	l.right.EnterBitMoveExpression(c)
}

// EnterBitMoveOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterBitMoveOperator(c *parser.BitMoveOperatorContext) {
	l.left.EnterBitMoveOperator(c)
	l.right.EnterBitMoveOperator(c)
}

// EnterBlockDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterBlockDef(c *parser.BlockDefContext) {
	l.left.EnterBlockDef(c)
	l.right.EnterBlockDef(c)
}

// EnterBracketExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterBracketExpression(c *parser.BracketExpressionContext) {
	l.left.EnterBracketExpression(c)
	l.right.EnterBracketExpression(c)
}

// EnterClassDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterClassDef(c *parser.ClassDefContext) {
	l.left.EnterClassDef(c)
	l.right.EnterClassDef(c)
}

// EnterCompExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterCompExpression(c *parser.CompExpressionContext) {
	l.left.EnterCompExpression(c)
	l.right.EnterCompExpression(c)
}

// EnterCompOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterCompOperator(c *parser.CompOperatorContext) {
	l.left.EnterCompOperator(c)
	l.right.EnterCompOperator(c)
}

// EnterConstArrayAssignment implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterConstArrayAssignment(c *parser.ConstArrayAssignmentContext) {
	l.left.EnterConstArrayAssignment(c)
	l.right.EnterConstArrayAssignment(c)
}

// EnterConstArrayDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterConstArrayDef(c *parser.ConstArrayDefContext) {
	l.left.EnterConstArrayDef(c)
	l.right.EnterConstArrayDef(c)
}

// EnterConstDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterConstDef(c *parser.ConstDefContext) {
	l.left.EnterConstDef(c)
	l.right.EnterConstDef(c)
}

// EnterConstValueAssignment implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterConstValueAssignment(c *parser.ConstValueAssignmentContext) {
	l.left.EnterConstValueAssignment(c)
	l.right.EnterConstValueAssignment(c)
}

// EnterConstValueDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterConstValueDef(c *parser.ConstValueDefContext) {
	l.left.EnterConstValueDef(c)
	l.right.EnterConstValueDef(c)
}

// EnterDaedalusFile implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterDaedalusFile(c *parser.DaedalusFileContext) {
	l.left.EnterDaedalusFile(c)
	l.right.EnterDaedalusFile(c)
}

// EnterElseBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterElseBlock(c *parser.ElseBlockContext) {
	l.left.EnterElseBlock(c)
	l.right.EnterElseBlock(c)
}

// EnterElseIfBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterElseIfBlock(c *parser.ElseIfBlockContext) {
	l.left.EnterElseIfBlock(c)
	l.right.EnterElseIfBlock(c)
}

// EnterEqExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterEqExpression(c *parser.EqExpressionContext) {
	l.left.EnterEqExpression(c)
	l.right.EnterEqExpression(c)
}

// EnterEqOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterEqOperator(c *parser.EqOperatorContext) {
	l.left.EnterEqOperator(c)
	l.right.EnterEqOperator(c)
}

// EnterExpressionBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterExpressionBlock(c *parser.ExpressionBlockContext) {
	l.left.EnterExpressionBlock(c)
	l.right.EnterExpressionBlock(c)
}

// EnterFloatLiteralValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterFloatLiteralValue(c *parser.FloatLiteralValueContext) {
	l.left.EnterFloatLiteralValue(c)
	l.right.EnterFloatLiteralValue(c)
}

// EnterFuncArgExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterFuncArgExpression(c *parser.FuncArgExpressionContext) {
	l.left.EnterFuncArgExpression(c)
	l.right.EnterFuncArgExpression(c)
}

// EnterFuncCall implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterFuncCall(c *parser.FuncCallContext) {
	l.left.EnterFuncCall(c)
	l.right.EnterFuncCall(c)
}

// EnterFuncCallValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterFuncCallValue(c *parser.FuncCallValueContext) {
	l.left.EnterFuncCallValue(c)
	l.right.EnterFuncCallValue(c)
}

// EnterFunctionDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterFunctionDef(c *parser.FunctionDefContext) {
	l.left.EnterFunctionDef(c)
	l.right.EnterFunctionDef(c)
}

// EnterIfBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterIfBlock(c *parser.IfBlockContext) {
	l.left.EnterIfBlock(c)
	l.right.EnterIfBlock(c)
}

// EnterIfBlockStatement implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterIfBlockStatement(c *parser.IfBlockStatementContext) {
	l.left.EnterIfBlockStatement(c)
	l.right.EnterIfBlockStatement(c)
}

// EnterIfCondition implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterIfCondition(c *parser.IfConditionContext) {
	l.left.EnterIfCondition(c)
	l.right.EnterIfCondition(c)
}

// EnterInlineDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterInlineDef(c *parser.InlineDefContext) {
	l.left.EnterInlineDef(c)
	l.right.EnterInlineDef(c)
}

// EnterInstanceDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterInstanceDecl(c *parser.InstanceDeclContext) {
	l.left.EnterInstanceDecl(c)
	l.right.EnterInstanceDecl(c)
}

// EnterInstanceDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterInstanceDef(c *parser.InstanceDefContext) {
	l.left.EnterInstanceDef(c)
	l.right.EnterInstanceDef(c)
}

// EnterIntegerLiteralValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterIntegerLiteralValue(c *parser.IntegerLiteralValueContext) {
	l.left.EnterIntegerLiteralValue(c)
	l.right.EnterIntegerLiteralValue(c)
}

// EnterLogAndExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterLogAndExpression(c *parser.LogAndExpressionContext) {
	l.left.EnterLogAndExpression(c)
	l.right.EnterLogAndExpression(c)
}

// EnterLogAndOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterLogAndOperator(c *parser.LogAndOperatorContext) {
	l.left.EnterLogAndOperator(c)
	l.right.EnterLogAndOperator(c)
}

// EnterLogOrExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterLogOrExpression(c *parser.LogOrExpressionContext) {
	l.left.EnterLogOrExpression(c)
	l.right.EnterLogOrExpression(c)
}

// EnterLogOrOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterLogOrOperator(c *parser.LogOrOperatorContext) {
	l.left.EnterLogOrOperator(c)
	l.right.EnterLogOrOperator(c)
}

// EnterMultExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterMultExpression(c *parser.MultExpressionContext) {
	l.left.EnterMultExpression(c)
	l.right.EnterMultExpression(c)
}

// EnterMultOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterMultOperator(c *parser.MultOperatorContext) {
	l.left.EnterMultOperator(c)
	l.right.EnterMultOperator(c)
}

// EnterNameNode implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterNameNode(c *parser.NameNodeContext) {
	l.left.EnterNameNode(c)
	l.right.EnterNameNode(c)
}

// EnterNullLiteralValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterNullLiteralValue(c *parser.NullLiteralValueContext) {
	l.left.EnterNullLiteralValue(c)
	l.right.EnterNullLiteralValue(c)
}

// EnterParameterDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterParameterDecl(c *parser.ParameterDeclContext) {
	l.left.EnterParameterDecl(c)
	l.right.EnterParameterDecl(c)
}

// EnterParameterList implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterParameterList(c *parser.ParameterListContext) {
	l.left.EnterParameterList(c)
	l.right.EnterParameterList(c)
}

// EnterParentReference implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterParentReference(c *parser.ParentReferenceContext) {
	l.left.EnterParentReference(c)
	l.right.EnterParentReference(c)
}

// EnterPrototypeDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterPrototypeDef(c *parser.PrototypeDefContext) {
	l.left.EnterPrototypeDef(c)
	l.right.EnterPrototypeDef(c)
}

// EnterReference implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterReference(c *parser.ReferenceContext) {
	l.left.EnterReference(c)
	l.right.EnterReference(c)
}

// EnterReferenceAtom implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterReferenceAtom(c *parser.ReferenceAtomContext) {
	l.left.EnterReferenceAtom(c)
	l.right.EnterReferenceAtom(c)
}

// EnterReferenceValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterReferenceValue(c *parser.ReferenceValueContext) {
	l.left.EnterReferenceValue(c)
	l.right.EnterReferenceValue(c)
}

// EnterReturnStatement implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterReturnStatement(c *parser.ReturnStatementContext) {
	l.left.EnterReturnStatement(c)
	l.right.EnterReturnStatement(c)
}

// EnterStatement implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterStatement(c *parser.StatementContext) {
	l.left.EnterStatement(c)
	l.right.EnterStatement(c)
}

// EnterStatementBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterStatementBlock(c *parser.StatementBlockContext) {
	l.left.EnterStatementBlock(c)
	l.right.EnterStatementBlock(c)
}

// EnterStringLiteralValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterStringLiteralValue(c *parser.StringLiteralValueContext) {
	l.left.EnterStringLiteralValue(c)
	l.right.EnterStringLiteralValue(c)
}

// EnterTypeReference implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterTypeReference(c *parser.TypeReferenceContext) {
	l.left.EnterTypeReference(c)
	l.right.EnterTypeReference(c)
}

// EnterValExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterValExpression(c *parser.ValExpressionContext) {
	l.left.EnterValExpression(c)
	l.right.EnterValExpression(c)
}

// EnterVarArrayDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterVarArrayDecl(c *parser.VarArrayDeclContext) {
	l.left.EnterVarArrayDecl(c)
	l.right.EnterVarArrayDecl(c)
}

// EnterVarDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterVarDecl(c *parser.VarDeclContext) {
	l.left.EnterVarDecl(c)
	l.right.EnterVarDecl(c)
}

// EnterVarValueDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) EnterVarValueDecl(c *parser.VarValueDeclContext) {
	l.left.EnterVarValueDecl(c)
	l.right.EnterVarValueDecl(c)
}

// ExitAddExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitAddExpression(c *parser.AddExpressionContext) {
	l.left.ExitAddExpression(c)
	l.right.ExitAddExpression(c)
}

// ExitAddOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitAddOperator(c *parser.AddOperatorContext) {
	l.left.ExitAddOperator(c)
	l.right.ExitAddOperator(c)
}

// ExitAnyIdentifier implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitAnyIdentifier(c *parser.AnyIdentifierContext) {
	l.left.ExitAnyIdentifier(c)
	l.right.ExitAnyIdentifier(c)
}

// ExitArrayIndex implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitArrayIndex(c *parser.ArrayIndexContext) {
	l.left.ExitArrayIndex(c)
	l.right.ExitArrayIndex(c)
}

// ExitArraySize implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitArraySize(c *parser.ArraySizeContext) {
	l.left.ExitArraySize(c)
	l.right.ExitArraySize(c)
}

// ExitAssignment implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitAssignment(c *parser.AssignmentContext) {
	l.left.ExitAssignment(c)
	l.right.ExitAssignment(c)
}

// ExitAssignmentOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitAssignmentOperator(c *parser.AssignmentOperatorContext) {
	l.left.ExitAssignmentOperator(c)
	l.right.ExitAssignmentOperator(c)
}

// ExitBinAndExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitBinAndExpression(c *parser.BinAndExpressionContext) {
	l.left.ExitBinAndExpression(c)
	l.right.ExitBinAndExpression(c)
}

// ExitBinAndOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitBinAndOperator(c *parser.BinAndOperatorContext) {
	l.left.ExitBinAndOperator(c)
	l.right.ExitBinAndOperator(c)
}

// ExitBinOrExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitBinOrExpression(c *parser.BinOrExpressionContext) {
	l.left.ExitBinOrExpression(c)
	l.right.ExitBinOrExpression(c)
}

// ExitBinOrOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitBinOrOperator(c *parser.BinOrOperatorContext) {
	l.left.ExitBinOrOperator(c)
	l.right.ExitBinOrOperator(c)
}

// ExitBitMoveExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitBitMoveExpression(c *parser.BitMoveExpressionContext) {
	l.left.ExitBitMoveExpression(c)
	l.right.ExitBitMoveExpression(c)
}

// ExitBitMoveOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitBitMoveOperator(c *parser.BitMoveOperatorContext) {
	l.left.ExitBitMoveOperator(c)
	l.right.ExitBitMoveOperator(c)
}

// ExitBlockDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitBlockDef(c *parser.BlockDefContext) {
	l.left.ExitBlockDef(c)
	l.right.ExitBlockDef(c)
}

// ExitBracketExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitBracketExpression(c *parser.BracketExpressionContext) {
	l.left.ExitBracketExpression(c)
	l.right.ExitBracketExpression(c)
}

// ExitClassDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitClassDef(c *parser.ClassDefContext) {
	l.left.ExitClassDef(c)
	l.right.ExitClassDef(c)
}

// ExitCompExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitCompExpression(c *parser.CompExpressionContext) {
	l.left.ExitCompExpression(c)
	l.right.ExitCompExpression(c)
}

// ExitCompOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitCompOperator(c *parser.CompOperatorContext) {
	l.left.ExitCompOperator(c)
	l.right.ExitCompOperator(c)
}

// ExitConstArrayAssignment implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitConstArrayAssignment(c *parser.ConstArrayAssignmentContext) {
	l.left.ExitConstArrayAssignment(c)
	l.right.ExitConstArrayAssignment(c)
}

// ExitConstArrayDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitConstArrayDef(c *parser.ConstArrayDefContext) {
	l.left.ExitConstArrayDef(c)
	l.right.ExitConstArrayDef(c)
}

// ExitConstDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitConstDef(c *parser.ConstDefContext) {
	l.left.ExitConstDef(c)
	l.right.ExitConstDef(c)
}

// ExitConstValueAssignment implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitConstValueAssignment(c *parser.ConstValueAssignmentContext) {
	l.left.ExitConstValueAssignment(c)
	l.right.ExitConstValueAssignment(c)
}

// ExitConstValueDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitConstValueDef(c *parser.ConstValueDefContext) {
	l.left.ExitConstValueDef(c)
	l.right.ExitConstValueDef(c)
}

// ExitDaedalusFile implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitDaedalusFile(c *parser.DaedalusFileContext) {
	l.left.ExitDaedalusFile(c)
	l.right.ExitDaedalusFile(c)
}

// ExitElseBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitElseBlock(c *parser.ElseBlockContext) {
	l.left.ExitElseBlock(c)
	l.right.ExitElseBlock(c)
}

// ExitElseIfBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitElseIfBlock(c *parser.ElseIfBlockContext) {
	l.left.ExitElseIfBlock(c)
	l.right.ExitElseIfBlock(c)
}

// ExitEqExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitEqExpression(c *parser.EqExpressionContext) {
	l.left.ExitEqExpression(c)
	l.right.ExitEqExpression(c)
}

// ExitEqOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitEqOperator(c *parser.EqOperatorContext) {
	l.left.ExitEqOperator(c)
	l.right.ExitEqOperator(c)
}

// ExitExpressionBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitExpressionBlock(c *parser.ExpressionBlockContext) {
	l.left.ExitExpressionBlock(c)
	l.right.ExitExpressionBlock(c)
}

// ExitFloatLiteralValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitFloatLiteralValue(c *parser.FloatLiteralValueContext) {
	l.left.ExitFloatLiteralValue(c)
	l.right.ExitFloatLiteralValue(c)
}

// ExitFuncArgExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitFuncArgExpression(c *parser.FuncArgExpressionContext) {
	l.left.ExitFuncArgExpression(c)
	l.right.ExitFuncArgExpression(c)
}

// ExitFuncCall implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitFuncCall(c *parser.FuncCallContext) {
	l.left.ExitFuncCall(c)
	l.right.ExitFuncCall(c)
}

// ExitFuncCallValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitFuncCallValue(c *parser.FuncCallValueContext) {
	l.left.ExitFuncCallValue(c)
	l.right.ExitFuncCallValue(c)
}

// ExitFunctionDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitFunctionDef(c *parser.FunctionDefContext) {
	l.left.ExitFunctionDef(c)
	l.right.ExitFunctionDef(c)
}

// ExitIfBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitIfBlock(c *parser.IfBlockContext) {
	l.left.ExitIfBlock(c)
	l.right.ExitIfBlock(c)
}

// ExitIfBlockStatement implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitIfBlockStatement(c *parser.IfBlockStatementContext) {
	l.left.ExitIfBlockStatement(c)
	l.right.ExitIfBlockStatement(c)
}

// ExitIfCondition implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitIfCondition(c *parser.IfConditionContext) {
	l.left.ExitIfCondition(c)
	l.right.ExitIfCondition(c)
}

// ExitInlineDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitInlineDef(c *parser.InlineDefContext) {
	l.left.ExitInlineDef(c)
	l.right.ExitInlineDef(c)
}

// ExitInstanceDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitInstanceDecl(c *parser.InstanceDeclContext) {
	l.left.ExitInstanceDecl(c)
	l.right.ExitInstanceDecl(c)
}

// ExitInstanceDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitInstanceDef(c *parser.InstanceDefContext) {
	l.left.ExitInstanceDef(c)
	l.right.ExitInstanceDef(c)
}

// ExitIntegerLiteralValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitIntegerLiteralValue(c *parser.IntegerLiteralValueContext) {
	l.left.ExitIntegerLiteralValue(c)
	l.right.ExitIntegerLiteralValue(c)
}

// ExitLogAndExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitLogAndExpression(c *parser.LogAndExpressionContext) {
	l.left.ExitLogAndExpression(c)
	l.right.ExitLogAndExpression(c)
}

// ExitLogAndOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitLogAndOperator(c *parser.LogAndOperatorContext) {
	l.left.ExitLogAndOperator(c)
	l.right.ExitLogAndOperator(c)
}

// ExitLogOrExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitLogOrExpression(c *parser.LogOrExpressionContext) {
	l.left.ExitLogOrExpression(c)
	l.right.ExitLogOrExpression(c)
}

// ExitLogOrOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitLogOrOperator(c *parser.LogOrOperatorContext) {
	l.left.ExitLogOrOperator(c)
	l.right.ExitLogOrOperator(c)
}

// ExitMultExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitMultExpression(c *parser.MultExpressionContext) {
	l.left.ExitMultExpression(c)
	l.right.ExitMultExpression(c)
}

// ExitMultOperator implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitMultOperator(c *parser.MultOperatorContext) {
	l.left.ExitMultOperator(c)
	l.right.ExitMultOperator(c)
}

// ExitNameNode implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitNameNode(c *parser.NameNodeContext) {
	l.left.ExitNameNode(c)
	l.right.ExitNameNode(c)
}

// ExitNullLiteralValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitNullLiteralValue(c *parser.NullLiteralValueContext) {
	l.left.ExitNullLiteralValue(c)
	l.right.ExitNullLiteralValue(c)
}

// ExitParameterDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitParameterDecl(c *parser.ParameterDeclContext) {
	l.left.ExitParameterDecl(c)
	l.right.ExitParameterDecl(c)
}

// ExitParameterList implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitParameterList(c *parser.ParameterListContext) {
	l.left.ExitParameterList(c)
	l.right.ExitParameterList(c)
}

// ExitParentReference implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitParentReference(c *parser.ParentReferenceContext) {
	l.left.ExitParentReference(c)
	l.right.ExitParentReference(c)
}

// ExitPrototypeDef implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitPrototypeDef(c *parser.PrototypeDefContext) {
	l.left.ExitPrototypeDef(c)
	l.right.ExitPrototypeDef(c)
}

// ExitReference implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitReference(c *parser.ReferenceContext) {
	l.left.ExitReference(c)
	l.right.ExitReference(c)
}

// ExitReferenceAtom implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitReferenceAtom(c *parser.ReferenceAtomContext) {
	l.left.ExitReferenceAtom(c)
	l.right.ExitReferenceAtom(c)
}

// ExitReferenceValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitReferenceValue(c *parser.ReferenceValueContext) {
	l.left.ExitReferenceValue(c)
	l.right.ExitReferenceValue(c)
}

// ExitReturnStatement implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitReturnStatement(c *parser.ReturnStatementContext) {
	l.left.ExitReturnStatement(c)
	l.right.ExitReturnStatement(c)
}

// ExitStatement implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitStatement(c *parser.StatementContext) {
	l.left.ExitStatement(c)
	l.right.ExitStatement(c)
}

// ExitStatementBlock implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitStatementBlock(c *parser.StatementBlockContext) {
	l.left.ExitStatementBlock(c)
	l.right.ExitStatementBlock(c)
}

// ExitStringLiteralValue implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitStringLiteralValue(c *parser.StringLiteralValueContext) {
	l.left.ExitStringLiteralValue(c)
	l.right.ExitStringLiteralValue(c)
}

// ExitTypeReference implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitTypeReference(c *parser.TypeReferenceContext) {
	l.left.ExitTypeReference(c)
	l.right.ExitTypeReference(c)
}

// ExitValExpression implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitValExpression(c *parser.ValExpressionContext) {
	l.left.ExitValExpression(c)
	l.right.ExitValExpression(c)
}

// ExitVarArrayDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitVarArrayDecl(c *parser.VarArrayDeclContext) {
	l.left.ExitVarArrayDecl(c)
	l.right.ExitVarArrayDecl(c)
}

// ExitVarDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitVarDecl(c *parser.VarDeclContext) {
	l.left.ExitVarDecl(c)
	l.right.ExitVarDecl(c)
}

// ExitVarValueDecl implements parser.DaedalusListener
func (l *CombinedDaedalusListener) ExitVarValueDecl(c *parser.VarValueDeclContext) {
	l.left.ExitVarValueDecl(c)
	l.right.ExitVarValueDecl(c)
}

var _ parser.DaedalusListener = (*CombinedDaedalusListener)(nil)
