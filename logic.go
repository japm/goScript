package evalGo

import (
	"fmt"
	"go/ast"
	"reflect"
	"strconv"
)

func evalBinaryExprAND(expr *ast.BinaryExpr, context interface{}) (interface{}, error) {
	left, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}

	lbool, err := castBool(left)
	if err != nil {
		return nil, err
	}
	if !lbool {
		return false, nil
	}
	right, err := eval(expr.Y, context)
	if err != nil {
		return nil, err
	}

	rbool, err := castBool(right)
	if err != nil {
		return nil, err
	}
	return rbool, nil
}

func evalBinaryExprOR(expr *ast.BinaryExpr, context interface{}) (interface{}, error) {
	left, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}
	lbool, err := castBool(left)
	if err != nil {
		return nil, err
	}
	if lbool {
		return true, nil
	}

	right, err := eval(expr.Y, context)
	if err != nil {
		return nil, err
	}
	rbool, err := castBool(right)
	if err != nil {
		return nil, err
	}
	return rbool, nil
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
