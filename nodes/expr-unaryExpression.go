package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type UnaryExpressionNode struct {
	ExpressionNode

	Operator lexer.Token
	Operand  ExpressionNode
}

// implement node type from interface
func (UnaryExpressionNode) NodeType() NodeType { return UnaryExpression }

// node print function
func (node UnaryExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ UnaryExpressionNode")
	fmt.Printf("%s  └ Operator: %s\n", indent, node.Operator.Kind)
	fmt.Println(indent + "  └ Operand: ")
	node.Operand.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateUnaryExpressionNode(op lexer.Token, expr ExpressionNode) UnaryExpressionNode {
	return UnaryExpressionNode{
		Operator: op,
		Operand:  expr,
	}
}
