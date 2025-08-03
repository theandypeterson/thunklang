package evaluator

import (
	"interpreter/ast"
	"interpreter/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node.Operator, node.Right, env)
	case *ast.InfixExpression:
		return evalInfixExpression(node.Left, node.Operator, node.Right, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.FunctionLiteral:
		return evalFunctionExpression(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.CallExpression:
		return evalCallExpression(node, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement, env)
	}
	return result
}

// return errororrrr
func evalPrefixExpression(operator string, right ast.Expression, env *object.Environment) object.Object {
	value := Eval(right, env)
	switch operator {
	case "-":
		if value.Type() == object.INTEGER_OBJ {
			integer := value.(*object.Integer)
			return &object.Integer{Value: -integer.Value}
		}
	}
	return nil
}

func evalInfixExpression(left ast.Expression, operator string, right ast.Expression, env *object.Environment) object.Object {
	leftVal := Eval(left, env)
	rightVal := Eval(right, env)
	switch operator {
	case "-":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			return &object.Integer{Value: leftSide.Value - rightSide.Value}
		}
	case "+":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			return &object.Integer{Value: leftSide.Value + rightSide.Value}
		}
	case "*":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			return &object.Integer{Value: leftSide.Value * rightSide.Value}
		}
	case "/":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			return &object.Integer{Value: leftSide.Value / rightSide.Value}
		}
	case ">":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			val := int64(0)
			if leftSide.Value > rightSide.Value {
				val = 1
			}
			return &object.Integer{Value: val}
		}
	case ">=":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			val := int64(0)
			if leftSide.Value >= rightSide.Value {
				val = 1
			}
			return &object.Integer{Value: val}
		}
	case "<":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			val := int64(0)
			if leftSide.Value < rightSide.Value {
				val = 1
			}
			return &object.Integer{Value: val}
		}
	case "<=":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			val := int64(0)
			if leftSide.Value <= rightSide.Value {
				val = 1
			}
			return &object.Integer{Value: val}
		}
	case "==":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			val := int64(0)
			if leftSide.Value == rightSide.Value {
				val = 1
			}
			return &object.Integer{Value: val}
		}
	case "!=":
		leftSide, rightSide := areInts(leftVal, rightVal)
		if leftSide != nil {
			val := int64(0)
			if leftSide.Value != rightSide.Value {
				val = 1
			}
			return &object.Integer{Value: val}
		}
	}
	return nil
}

func areInts(left object.Object, right object.Object) (*object.Integer, *object.Integer) {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return left.(*object.Integer), right.(*object.Integer)
	}
	return nil, nil
}

func evalIfExpression(node *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(node.Condition, env)

	if isTruthy(condition) {
		return Eval(node.Consequence, env)
	} else if node.Alternative != nil {
		return Eval(node.Alternative, env)
	} else {
		return nil
	}
}

func isTruthy(obj object.Object) bool {
	switch obj.Type() {
	case object.INTEGER_OBJ:
		return obj.(*object.Integer).Value > 0
	default:
		return false
	}
}

func evalFunctionExpression(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	functionName := node.Name
	function := &object.Function{Parameters: node.Parameters, Name: node.Name, Body: node.Body, Env: env}

	if functionName != nil {
		env.Set(functionName.Value, function)
	}

	return function
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return nil
	}
	return val
}

func evalCallExpression(node *ast.CallExpression, env *object.Environment) object.Object {
	function := Eval(node.Function, env)
	args := []object.Object{}
	for _, arg := range node.Arguments {
		args = append(args, Eval(arg, env))
	}
	return applyFunction(function, args)
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return nil
	}

	extendedEnv := extendFunctionEnv(function, args)
	evaluated := Eval(function.Body, extendedEnv)
	return evaluated
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}
