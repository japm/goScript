package evalGo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

type callSyte struct {
	callee interface{}
	fnName string
}

//Eval expression within a context
func Eval(expr string, context map[string]interface{}) (val interface{}, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Eval paniked. %s", r)
		}
	}()

	exp, err := parser.ParseExpr(expr)

	if err != nil {
		return nil, err
	}

	val, err = eval(exp, context)
	return eval(exp, context)
}

func eval(expr ast.Node, context map[string]interface{}) (interface{}, error) {
	switch expr.(type) {
	case *ast.ArrayType:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.AssignStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.BadDecl:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.BadExpr:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.BadStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.BasicLit:
		return evalBasicLit(expr.(*ast.BasicLit), context)
	case *ast.BinaryExpr:
		return evalBinaryExpr(expr.(*ast.BinaryExpr), context)
	case *ast.BlockStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.BranchStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.CallExpr:
		return evalCallExpr(expr.(*ast.CallExpr), context)
	case *ast.CaseClause:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.ChanType:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.CommClause:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.Comment:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.CommentGroup:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.CompositeLit:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.DeclStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.DeferStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.Ellipsis:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.EmptyStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.Field:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.FieldList:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.ForStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.FuncDecl:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.FuncLit:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.FuncType:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.GenDecl:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.GoStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.Ident:
		return evalIdent(expr.(*ast.Ident), context)
	case *ast.IfStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.ImportSpec:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.IncDecStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.IndexExpr:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.InterfaceType:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.KeyValueExpr:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.LabeledStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.MapType:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.Package:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.ParenExpr:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.RangeStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.ReturnStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.SelectStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.SelectorExpr:
		return evalSelectorExpr(expr.(*ast.SelectorExpr), context)
	case *ast.SendStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.SliceExpr:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.StarExpr:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.StructType:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.SwitchStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.TypeAssertExpr:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.TypeSpec:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.TypeSwitchStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.UnaryExpr:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.ValueSpec:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	default:
		return nil, fmt.Errorf("Default %s not suported", reflect.TypeOf(expr))
	}
}

func evalSelectorExpr(expr *ast.SelectorExpr, context map[string]interface{}) (interface{}, error) {

	callee, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}

	fnName := expr.Sel.Name

	return &callSyte{callee, fnName}, nil
}

func evalCallExpr(expr *ast.CallExpr, context map[string]interface{}) (interface{}, error) {

	args := make([]reflect.Value, len(expr.Args))

	for key, value := range expr.Args {
		val, err := eval(value, context)
		if err != nil {
			return nil, err
		}
		args[key] = reflect.ValueOf(val)
	}

	val, err := eval(expr.Fun, context)
	if err != nil {
		return nil, err
	}
	switch val.(type) {
	case *callSyte:
		break
	default:
		return nil, fmt.Errorf("Waiting callsite found %s", reflect.TypeOf(val))
	}

	callsite := val.(*callSyte)
	_, ok := reflect.TypeOf(callsite.callee).MethodByName(callsite.fnName)
	if !ok {
		return nil, fmt.Errorf("Method %s not found", callsite.fnName)
	}

	caleeVal := reflect.ValueOf(callsite.callee)
	method := caleeVal.MethodByName(callsite.fnName)

	retVal := method.Call(args)

	return retVal[0].Interface(), nil
}

func evalBinaryExpr(expr *ast.BinaryExpr, context map[string]interface{}) (interface{}, error) {
	left, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}

	right, err := eval(expr.Y, context)
	if err != nil {
		return nil, err
	}

	return evalBinaryExprOp(expr, left, right)
}

func evalIdent(expr *ast.Ident, context map[string]interface{}) (interface{}, error) {

	lname := len(expr.Name)
	if lname == 4 && expr.Name == "true" {
		return true, nil
	} else if lname == 5 && expr.Name == "false" {
		return false, nil
	}
	if context == nil {
		return nil, fmt.Errorf("Context is null, no ident possible for %s", expr.Name)
	}
	val, ok := context[expr.Name]
	if !ok {
		return nil, fmt.Errorf("Symbol %s not found", expr.Name)
	}
	return val, nil
}

func evalBasicLit(expr *ast.BasicLit, context map[string]interface{}) (interface{}, error) {
	switch expr.Kind {
	case token.INT:
		return strconv.ParseInt(expr.Value, 10, 64)
	case token.FLOAT:
		return strconv.ParseFloat(expr.Value, 10)
	case token.IMAG:
		return nil, fmt.Errorf(" token.IMAG not suported")
	case token.CHAR:
		return expr.Value, nil
	case token.STRING:
		return expr.Value[1 : len(expr.Value)-1], nil
	default:
		return nil, fmt.Errorf("token type not suported %d %s", expr.Kind, expr.Value)
	}
}

func evalBinaryExprOp(expr *ast.BinaryExpr, left interface{}, right interface{}) (interface{}, error) {
	switch expr.Op {
	case token.ADD:
		return evalBinaryExprADD(left, right)
	case token.SUB:
		return evalBinaryExprSUB(left, right)
	case token.MUL:
		return evalBinaryExprMUL(left, right)
	case token.QUO:
		return evalBinaryExprQUO(left, right)
	case token.REM:
		return evalBinaryExprREM(left, right)
	case token.EQL: // ==
		return evalBinaryExprEQL(left, right)
	case token.LSS: // <
		return evalBinaryExprLSS(left, right)
	case token.GTR: // >
		return evalBinaryExprGTR(left, right)
	case token.NEQ: // !=
		return evalBinaryExprNEQ(left, right)
	case token.GEQ: // >=
		return evalBinaryExprGEQ(left, right)
	case token.LEQ: // <=
		return evalBinaryExprLEQ(left, right)
	default:
		return nil, fmt.Errorf("evalBinaryExprOp not implemented for %d", expr.Op)
	}

}

func evalBinaryExprADD(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) + right.(int64), nil
		case int:
			return left.(int64) + int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) + right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) + vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) + valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) + int64(1), nil
			}
			return left, nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) + right.(int64), nil
		case int:
			return left.(int) + right.(int), nil
		case float64:
			return float64(left.(int)) + right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) + vali, nil
			}

			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) + valf, nil
		case bool:
			if right.(bool) {
				return left.(int) + 1, nil
			}
			return left, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) + float64(right.(int64)), nil
		case int:
			return left.(float64) + float64(right.(int)), nil
		case float64:
			return left.(float64) + right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) + valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) + float64(1), nil
			}
			return left, nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) + right.(string), nil
		case int64:
			return left.(string) + strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) + strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) + strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) + strconv.FormatBool(right.(bool)), nil
		}
	}

	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprSUB(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) - right.(int64), nil
		case int:
			return left.(int64) - int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) - right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) - vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) - valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) - int64(1), nil
			}
			return left, nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) - right.(int64), nil
		case int:
			return left.(int) - right.(int), nil
		case float64:
			return float64(left.(int)) - right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) - vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) - valf, nil
		case bool:
			if right.(bool) {
				return left.(int) - 1, nil
			}
			return left, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) - float64(right.(int64)), nil
		case int:
			return left.(float64) - float64(right.(int)), nil
		case float64:
			return left.(float64) - right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) - valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) - float64(1), nil
			}
			return left, nil
		}
	case string:
		switch right.(type) {
		case string:
			return strings.Replace(left.(string), right.(string), "", -1), nil
		case int64:
			return strings.Replace(left.(string), strconv.FormatInt(right.(int64), 10), "", -1), nil
		case int:
			return strings.Replace(left.(string), strconv.Itoa(right.(int)), "", -1), nil
		case float64:
			return strings.Replace(left.(string), strconv.FormatFloat(right.(float64), 'f', -1, 64), "", -1), nil
		case bool:
			return strings.Replace(left.(string), strconv.FormatBool(right.(bool)), "", -1), nil
		}
	}
	return nil, fmt.Errorf("Unimplemented sub for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprMUL(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) * right.(int64), nil
		case int:
			return left.(int64) * int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) * right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) * vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) * valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) * right.(int64), nil
		case int:
			return left.(int) * right.(int), nil
		case float64:
			return float64(left.(int)) * right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) * vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) * valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) * float64(right.(int64)), nil
		case int:
			return left.(float64) * float64(right.(int)), nil
		case float64:
			return left.(float64) * right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(float64) * float64(vali), nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) * valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err != nil {
				valf, err := strconv.ParseFloat(right.(string), 10)
				if err != nil {
					return nil, err
				}
				vali = int64(valf)
			}
			return strings.Repeat(left.(string), int(vali)), nil
		case int64:
			return strings.Repeat(left.(string), int(right.(int64))), nil
		case int:
			return strings.Repeat(left.(string), right.(int)), nil
		case float64:
			return strings.Repeat(left.(string), int(right.(float64))), nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return "", nil
		}
	}
	return nil, fmt.Errorf("Unimplemented mul for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprQUO(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) / right.(int64), nil
		case int:
			return left.(int64) / int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) / right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) / vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) / valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return nil, fmt.Errorf("Divide by false no allowed")
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) / right.(int64), nil
		case int:
			return left.(int) / right.(int), nil
		case float64:
			return float64(left.(int)) / right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) / vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) / valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return nil, fmt.Errorf("Divide by false no allowed")
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) / float64(right.(int64)), nil
		case int:
			return left.(float64) / float64(right.(int)), nil
		case float64:
			return left.(float64) / right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(float64) / float64(vali), nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) / valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return nil, fmt.Errorf("Divide by false no allowed")
		}
	}
	return nil, fmt.Errorf("Unimplemented quo for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprREM(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) % right.(int64), nil
		case int:
			return left.(int64) % int64(right.(int)), nil
		case float64:
			return left.(int64) % int64(right.(float64)), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) % vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(int64) % int64(valf), nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return nil, fmt.Errorf("Mod by false no allowed")
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) % right.(int64), nil
		case int:
			return left.(int) % right.(int), nil
		case float64:
			return int64(left.(int)) % int64(right.(float64)), nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return nil, fmt.Errorf("Mod by false no allowed")
		}
	case float64:
		switch right.(type) {
		case int64:
			return int64(left.(float64)) % right.(int64), nil
		case int:
			return int64(left.(float64)) % int64(right.(int)), nil
		case float64:
			return int64(left.(float64)) % int64(right.(float64)), nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return nil, fmt.Errorf("Mod by false no allowed")
		}
	}
	return nil, fmt.Errorf("Unimplemented rem for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprEQL(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) == right.(int64), nil
		case int:
			return left.(int64) == int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) == right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) == vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) == valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) == int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) == right.(int64), nil
		case int:
			return left.(int) == right.(int), nil
		case float64:
			return float64(left.(int)) == right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) == vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) == valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) == 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) == float64(right.(int64)), nil
		case int:
			return left.(float64) == float64(right.(int)), nil
		case float64:
			return left.(float64) == right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) == valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) == float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) == right.(string), nil
		case int64:
			return left.(string) == strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) == strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) == strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) == strconv.FormatBool(right.(bool)), nil
		}
	case bool:
		switch right.(type) {
		case string:
			valb, err := strconv.ParseBool(right.(string))
			if err == nil {
				return left.(bool) == valb, nil
			}
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(bool) == (vali != int64(0)), nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(bool) == (valf != float64(0)), nil
		case int64:
			return left.(bool) == (right.(int64) != int64(0)), nil
		case int:
			return left.(bool) == (right.(int) != int(0)), nil
		case float64:
			return left.(bool) == (right.(float64) != float64(0)), nil
		case bool:
			return left.(bool) == right.(bool), nil
		}
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprLSS(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) < right.(int64), nil
		case int:
			return left.(int64) < int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) < right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) < vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) < valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) < int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) < right.(int64), nil
		case int:
			return left.(int) < right.(int), nil
		case float64:
			return float64(left.(int)) < right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) < vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) < valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) < 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) < float64(right.(int64)), nil
		case int:
			return left.(float64) < float64(right.(int)), nil
		case float64:
			return left.(float64) < right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) < valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) < float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) < right.(string), nil
		case int64:
			return left.(string) < strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) < strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) < strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) < strconv.FormatBool(right.(bool)), nil
		}
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprGTR(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) > right.(int64), nil
		case int:
			return left.(int64) > int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) > right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) > vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) > valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) > int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) > right.(int64), nil
		case int:
			return left.(int) > right.(int), nil
		case float64:
			return float64(left.(int)) > right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) > vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) > valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) > 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) > float64(right.(int64)), nil
		case int:
			return left.(float64) > float64(right.(int)), nil
		case float64:
			return left.(float64) > right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) > valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) > float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) > right.(string), nil
		case int64:
			return left.(string) > strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) > strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) > strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) > strconv.FormatBool(right.(bool)), nil
		}
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprGEQ(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) >= right.(int64), nil
		case int:
			return left.(int64) >= int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) >= right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) >= vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) >= valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) >= int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) >= right.(int64), nil
		case int:
			return left.(int) >= right.(int), nil
		case float64:
			return float64(left.(int)) >= right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) >= vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) >= valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) >= 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) >= float64(right.(int64)), nil
		case int:
			return left.(float64) >= float64(right.(int)), nil
		case float64:
			return left.(float64) >= right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) >= valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) >= float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) >= right.(string), nil
		case int64:
			return left.(string) >= strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) >= strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) >= strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) >= strconv.FormatBool(right.(bool)), nil
		}
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}
func evalBinaryExprLEQ(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) <= right.(int64), nil
		case int:
			return left.(int64) <= int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) <= right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) <= vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) <= valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) <= int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) <= right.(int64), nil
		case int:
			return left.(int) <= right.(int), nil
		case float64:
			return float64(left.(int)) <= right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) <= vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) <= valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) <= 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) <= float64(right.(int64)), nil
		case int:
			return left.(float64) <= float64(right.(int)), nil
		case float64:
			return left.(float64) <= right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) <= valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) <= float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) <= right.(string), nil
		case int64:
			return left.(string) <= strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) <= strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) <= strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) <= strconv.FormatBool(right.(bool)), nil
		}
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprNEQ(left interface{}, right interface{}) (interface{}, error) {
	val, err := evalBinaryExprEQL(left, right)
	if err != nil {
		return nil, err
	}
	return !val.(bool), nil
}
