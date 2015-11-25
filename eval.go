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
	callee     interface{}
	fnName     string
	calleeVal  reflect.Value
	calleeType reflect.Type
}

type Expr struct {
	expr ast.Expr
}

func (e *Expr) Prepare(expr string) error {
	exp, err := parser.ParseExpr(expr)
	if err != nil {
		return err
	}
	e.expr = exp
	return err
}

func (e *Expr) Eval(context map[string]interface{}) (val interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Eval paniked. %s", r)
		}
	}()
	val, err = eval(e.expr, context)
	return val, err
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

func evalInt(expr ast.Node, context map[string]interface{}) (int, error) {
	val, err := eval(expr, context)
	if err != nil {
		return 0, err
	}
	return castInt(val)
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
		return evalIndexExpr(expr.(*ast.IndexExpr), context)
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
		return evalParenExpr(expr.(*ast.ParenExpr), context)
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
		return evalSliceExpr(expr.(*ast.SliceExpr), context)
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
		return evalUnaryExpr(expr.(*ast.UnaryExpr), context)
	case *ast.ValueSpec:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	default:
		return nil, fmt.Errorf("Default %s not suported", reflect.TypeOf(expr))
	}
}

func evalIndexExpr(expr *ast.IndexExpr, context map[string]interface{}) (interface{}, error) {

	val, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}
	sl := reflect.ValueOf(val)

	if sl.Kind() != reflect.Map {
		return nil, fmt.Errorf("Expected map %d, found %d ", reflect.Array, sl.Kind())
	}

	idx, err := eval(expr.Index, context)
	if err != nil {
		return nil, err
	}

	retVal := sl.MapIndex(reflect.ValueOf(idx))

	if !retVal.IsValid() {
		return nil, nil
	}
	if retVal.IsNil() {
		return nil, nil
	}
	return sl.MapIndex(reflect.ValueOf(idx)).Interface(), nil
}

func evalSliceExpr(expr *ast.SliceExpr, context map[string]interface{}) (interface{}, error) {

	val, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}
	sl := reflect.ValueOf(val)

	if sl.Kind() != reflect.Array &&
		sl.Kind() != reflect.Slice {
		return nil, fmt.Errorf("Expected array %d, found %d ", reflect.Array, sl.Kind())
	}

	var i, j, k int
	if expr.Low == nil {
		i = 0
	} else {
		i, err = evalInt(expr.Low, context)
		if err != nil {
			return nil, err
		}
	}

	if expr.High == nil {
		j = sl.Len()
	} else {
		j, err = evalInt(expr.High, context)
		if err != nil {
			return nil, err
		}
	}

	if expr.Slice3 {
		k, err = evalInt(expr.Max, context)
		if err != nil {
			return nil, err
		}
		return sl.Slice3(i, j, k).Interface(), nil
	}
	return sl.Slice(i, j).Interface(), nil
}

func evalParenExpr(expr *ast.ParenExpr, context map[string]interface{}) (interface{}, error) {
	return eval(expr.X, context)
}

func evalUnaryExpr(expr *ast.UnaryExpr, context map[string]interface{}) (interface{}, error) {
	val, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}
	switch expr.Op {
	case token.NOT:
		return evalUnaryExprNOT(val)
	case token.SUB:
		return evalUnaryExprSUB(val)
	case token.AND:
		return evalUnaryExprAND(val)
	default:
		return nil, fmt.Errorf("Unary operation %d not implemented", expr.Op)
	}
}

func evalSelectorExpr(expr *ast.SelectorExpr, context map[string]interface{}) (interface{}, error) {
	callee, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}

	calleeVal := reflect.ValueOf(callee)
	calleeType := reflect.TypeOf(callee)

	_, ok := calleeType.FieldByName(expr.Sel.Name)
	if ok {
		return calleeVal.FieldByName(expr.Sel.Name).Interface(), nil
	}
	return callSyte{callee, expr.Sel.Name, calleeVal, calleeType}, nil
}

func evalCallExpr(expr *ast.CallExpr, context map[string]interface{}) (interface{}, error) {

	val, err := eval(expr.Fun, context)
	if err != nil {
		return nil, err
	}
	switch val.(type) {
	case callSyte:
		break
	default:
		return nil, fmt.Errorf("Waiting callsite found %s", reflect.TypeOf(val))
	}

	callsite := val.(callSyte)
	tpM, ok := callsite.calleeType.MethodByName(callsite.fnName)
	if !ok {
		return nil, fmt.Errorf("Method %s not found", callsite.fnName)
	}

	numArgs := tpM.Type.NumIn()
	argsOffset := 0
	if callsite.callee != nil {
		numArgs = numArgs - 1
		argsOffset = 1
	}
	if len(expr.Args) != numArgs {
		return nil, fmt.Errorf("Method alguments count mismatch. Expected %d get %d", numArgs, len(expr.Args))
	}

	caleeVal := callsite.calleeVal
	method := caleeVal.MethodByName(callsite.fnName)

	args := make([]reflect.Value, len(expr.Args))
	for key, value := range expr.Args {
		tpArg := tpM.Type.In(key + argsOffset)

		val, err := eval(value, context)
		if err != nil {
			return nil, err
		}

		rVal := reflect.ValueOf(val)
		if !rVal.Type().ConvertibleTo(tpArg) {
			return nil, fmt.Errorf("Method alguments type mismatch. Expected %d get %d", tpArg, rVal.Type())
		}

		args[key] = rVal
	}

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
	if lname == 3 && expr.Name == "nil" {
		return nil, nil
	} else if lname == 4 && expr.Name == "true" {
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
	case token.LAND: // &&
		return evalBinaryExprAND(left, right)
	case token.LOR: // ||
		return evalBinaryExprOR(left, right)
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
		case nil:
			return left.(int64), nil
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
		case nil:
			return left.(int), nil
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
		case nil:
			return left.(float64), nil
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
		case nil:
			return left.(string), nil
		}
	case nil:
		return nil, nil
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
		case nil:
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
		case nil:
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
		case nil:
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
		case nil:
			return left, nil
		}
	case nil:
		return nil, nil
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
		case nil:
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
		case nil:
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
		case nil:
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
		case nil:
			return "", nil
		}
	case nil:
		return nil, nil
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
		case nil:
			return nil, fmt.Errorf("Divide by <nil> no allowed")
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
		case nil:
			return nil, fmt.Errorf("Divide by <nil> no allowed")
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
		case nil:
			return nil, fmt.Errorf("Divide by <nil> no allowed")
		}
	case nil:
		return nil, nil
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
		case nil:
			return nil, fmt.Errorf("Mod by <nil> no allowed")
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
		case nil:
			return nil, fmt.Errorf("Mod by <nil> no allowed")
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
		case nil:
			return nil, fmt.Errorf("Mod by <nil> no allowed")
		}
	case nil:
		return nil, nil
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
		case nil:
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
		case nil:
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
		case nil:
			return left.(float64) == 0, nil
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
		case nil:
			return left.(string) == "", nil
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
		case nil:
			return !left.(bool), nil
		}
	case nil:
		return nil == right, nil
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprLSS(left interface{}, right interface{}) (interface{}, error) {
	if right == nil {
		return nil, nil
	}
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
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprGTR(left interface{}, right interface{}) (interface{}, error) {
	if right == nil {
		return nil, nil
	}
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
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprGEQ(left interface{}, right interface{}) (interface{}, error) {
	if right == nil {
		return nil, nil
	}
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
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}
func evalBinaryExprLEQ(left interface{}, right interface{}) (interface{}, error) {
	if right == nil {
		return nil, nil
	}
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
	case nil:
		return nil, nil
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

func evalUnaryExprNOT(value interface{}) (interface{}, error) {
	switch value.(type) {
	case int64:
		return value.(int64) == int64(0), nil
	case int:
		return value.(int) == 0, nil
	case float64:
		return value.(float64) == float64(0), nil
	case string:
		valb, err := strconv.ParseBool(value.(string))
		if err == nil {
			return !valb, nil
		}
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return vali == 0, nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return nil, err
		}
		return valf == float64(0), nil
	case bool:
		return !value.(bool), nil
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented not for type %s", reflect.TypeOf(value))
}

func evalUnaryExprSUB(value interface{}) (interface{}, error) {
	switch value.(type) {
	case int64:
		return -value.(int64), nil
	case int:
		return -value.(int), nil
	case float64:
		return -value.(float64), nil
	case bool:
		return !value.(bool), nil
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented not for type %s", reflect.TypeOf(value))
}

func evalUnaryExprAND(value interface{}) (interface{}, error) {

	val := reflect.ValueOf(value)

	if !val.CanAddr() {
		return nil, fmt.Errorf("Value is not addressable %s", val)
	}
	if val.IsNil() {
		return nil, fmt.Errorf("Value is nill, not addressable %s", val)
	}

	return val.Addr(), nil
}

func evalBinaryExprAND(left interface{}, right interface{}) (interface{}, error) {
	lbool, err := castBool(left)
	if err != nil {
		return nil, err
	}
	if !lbool {
		return false, nil
	}
	rbool, err := castBool(right)

	if err != nil {
		return nil, err
	}
	return rbool, nil
}

func evalBinaryExprOR(left interface{}, right interface{}) (interface{}, error) {
	lbool, err := castBool(left)
	if err != nil {
		return nil, err
	}
	if lbool {
		return true, nil
	}
	rbool, err := castBool(right)
	if err != nil {
		return nil, err
	}
	return rbool, nil
}

func castBool(value interface{}) (bool, error) {
	switch value.(type) {
	case int64:
		return value.(int64) != int64(0), nil
	case int:
		return value.(int) != 0, nil
	case float64:
		return value.(float64) != float64(0), nil
	case string:
		valb, err := strconv.ParseBool(value.(string))
		if err == nil {
			return valb, nil
		}
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return vali != 0, nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return false, err
		}
		return valf != float64(0), nil
	case bool:
		return value.(bool), nil
	case nil:
		return false, nil
	}
	return false, fmt.Errorf("Unimplemented cast to bool for type %s", reflect.TypeOf(value))
}

func castInt(value interface{}) (int, error) {
	switch value.(type) {
	case int64:
		return int(value.(int64)), nil
	case int:
		return value.(int), nil
	case float64:
		return int(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return int(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return int(valf), nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to bool for type %s", reflect.TypeOf(value))
}
