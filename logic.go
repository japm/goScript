package evalGo

import (
	"fmt"
	"go/ast"
	"reflect"
)

func evalBinaryExprLAND(expr *ast.BinaryExpr, context interface{}) (interface{}, error) {
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

func evalBinaryExprLOR(expr *ast.BinaryExpr, context interface{}) (interface{}, error) {
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

	tp, e := valType(value)

	if e != nil {
		return nil, e
	}

	if tp.IsNil {
		return nil, nil
	}
	if !tp.IsNumeric {
		switch value.(type) {
		case string:
			valb, err := castBool(value.(string))
			if err == nil {
				return !valb, nil
			}
			vali, err := castInt(value.(string))
			if err == nil {
				return vali == 0, nil
			}
			valf, err := castFloat64(value.(string))
			if err != nil {
				return nil, err
			}
			return valf == float64(0), nil

		case bool:
			return !value.(bool), nil

		case nil:
			return nil, nil
		}
	} else if tp.Signed {
		if tp.Float {
			if tp.Size == 32 {
				l, err := castFloat32(value)
				if err != nil {
					return nil, err
				}

				return l == 0, nil
			} else {
				l, err := castFloat64(value)
				if err != nil {
					return nil, err
				}
				return l == 0, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(value)
			if err != nil {
				return nil, err
			}

			return l == 0, nil
		} else if tp.Size == 32 {
			l, err := castInt(value)
			if err != nil {
				return nil, err
			}

			return l == 0, nil
		} else if tp.Size == 16 {
			l, err := castInt16(value)
			if err != nil {
				return nil, err
			}

			return l == 0, nil
		} else if tp.Size == 8 {
			l, err := castInt8(value)
			if err != nil {
				return nil, err
			}

			return l == 0, nil
		}
	} else {
		if tp.Size == 64 {
			l, err := castUint64(value)
			if err != nil {
				return nil, err
			}

			return l == 0, nil
		} else if tp.Size == 32 {
			l, err := castUint(value)
			if err != nil {
				return nil, err
			}

			return l == 0, nil
		} else if tp.Size == 16 {
			l, err := castUint16(value)
			if err != nil {
				return nil, err
			}

			return l == 0, nil
		} else if tp.Size == 8 {
			l, err := castUint8(value)
			if err != nil {
				return nil, err
			}
			return l == 0, nil
		}
	}
	return nil, fmt.Errorf("Unimplemented not for type  %s ", reflect.TypeOf(value))

}
