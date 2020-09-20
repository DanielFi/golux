package main

import (
	"fmt"

	"github.com/DanielFi/golux/internal/ast"
	"github.com/DanielFi/golux/internal/runtime"
)

func program1() {
	// Evaluate the expression:
	// (1 + 2) * 3

	var node ast.BinaryOperationExpression
	node = ast.BinaryOperationExpression{
		Operator: ast.LessEquals,
		LHS: ast.BinaryOperationExpression{
			Operator: ast.Plus,
			LHS:      ast.IntegerLiteral(1),
			RHS:      ast.IntegerLiteral(2),
		},
		RHS: ast.IntegerLiteral(3),
	}
	fmt.Println(node.Evaluate(runtime.NewInterpreter()))
}

func program2() {
	// Evaluate the program:
	// var x = 1
	// x + 2

	var node ast.Expression
	node = ast.BlockExpression{
		Expressions: []ast.Expression{
			ast.VariableDeclaration{
				LHS: ast.Identifier("x"),
				RHS: ast.IntegerLiteral(1),
			},
			ast.BinaryOperationExpression{
				Operator: ast.Plus,
				LHS:      ast.Identifier("x"),
				RHS:      ast.IntegerLiteral(2),
			},
		},
	}
	fmt.Println(node.Evaluate(runtime.NewInterpreter()))
}

func program3() {
	// Evaluate the program:
	// fun foo(x) {
	//	  x + 2
	// }
	// foo(1)

	var node ast.Expression
	node = ast.BlockExpression{
		Expressions: []ast.Expression{
			ast.FunctionDeclaration{
				Name: ast.Identifier("foo"),
				Arguments: []ast.Identifier{
					ast.Identifier("x"),
				},
				Body: ast.BinaryOperationExpression{
					Operator: ast.Plus,
					LHS:      ast.Identifier("x"),
					RHS:      ast.IntegerLiteral(2),
				},
			},
			ast.CallExpression{
				Function: ast.Identifier("foo"),
				Parameters: []ast.Expression{
					ast.IntegerLiteral(1),
				},
			},
		},
	}
	fmt.Println(node.Evaluate(runtime.NewInterpreter()))
}

func program4() {
	// Evaluate the program:
	// fun foo(x) {
	//	  fun bar() {
	//       x + 1
	//    }
	// }
	// var foobar = foo(2)
	// foobar()

	var node ast.Expression
	node = ast.BlockExpression{
		Expressions: []ast.Expression{
			ast.FunctionDeclaration{
				Name: ast.Identifier("foo"),
				Arguments: []ast.Identifier{
					ast.Identifier("x"),
				},
				Body: ast.BlockExpression{
					Expressions: []ast.Expression{
						ast.FunctionDeclaration{
							Name:      ast.Identifier("bar"),
							Arguments: []ast.Identifier{},
							Body: ast.BinaryOperationExpression{
								Operator: ast.Plus,
								LHS:      ast.Identifier("x"),
								RHS:      ast.IntegerLiteral(1),
							},
						},
					},
				},
			},
			ast.VariableDeclaration{
				LHS: ast.Identifier("foobar"),
				RHS: ast.CallExpression{
					Function:   ast.Identifier("foo"),
					Parameters: []ast.Expression{ast.IntegerLiteral(2)},
				},
			},
			ast.CallExpression{
				Function:   ast.Identifier("foobar"),
				Parameters: []ast.Expression{},
			},
		},
	}
	fmt.Println(node.Evaluate(runtime.NewInterpreter()))
}

func main() {
	program1()
	program2()
	program3()
	program4()
}
