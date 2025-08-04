package evaluator

import "interpreter/object"

var builtins = map[string]*object.Builtin{
	"length": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				// todo add error
				return nil
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				// TODO add error
				return nil
			}
		},
	},
}
