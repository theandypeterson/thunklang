package evaluator

import (
	"interpreter/ast"
	"interpreter/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node.Operator, node.Right)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}

// return errororrrr
func evalPrefixExpression(operator string, right ast.Expression) object.Object {
	value := Eval(right)
	switch operator {
	case "-":
		if value.Type() == object.INTEGER_OBJ {
			integer := value.(*object.Integer)
			return &object.Integer{Value: -integer.Value}
		}
	}
	return nil
}
