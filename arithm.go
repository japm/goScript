/*
The MIT License (MIT)
Copyright (c) 2016 Juan Pascual
*/
package goScript

import (
	"fmt"
	"reflect"
	"strings"
)

func evalBinaryExprADD(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:
			val, err := castString(right)
			if err != nil {
				return nil, err
			}
			return left.(string) + val, nil

		case bool:
			vall, err := castInt(left)
			if err != nil {
				return nil, err
			}
			valr, err := castInt(right)
			if err != nil {
				return nil, err
			}

			return vall + valr, nil

		case nil:
			return right, nil
		}
	} else {
		op := opAdd{}
		return evalBinaryNumeric(left, right, tp, op)
	}
	return nil, fmt.Errorf("Unimplemented add for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprSUB(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:
			val, err := castString(right)
			if err != nil {
				return nil, err
			}
			return strings.Replace(left.(string), val, "", -1), nil

		case bool:
			vall, err := castInt(left)
			if err != nil {
				return nil, err
			}
			valr, err := castInt(right)
			if err != nil {
				return nil, err
			}

			return vall - valr, nil
		case nil:
			return evalUnaryExprSUB(right)
		}
	} else {
		op := opSub{}
		return evalBinaryNumeric(left, right, tp, op)
	}
	return nil, fmt.Errorf("Unimplemented sub for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprMUL(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:
			val, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return strings.Repeat(left.(string), val), nil

		case bool:
			vall, err := castInt(left)
			if err != nil {
				return nil, err
			}
			valr, err := castInt(right)
			if err != nil {
				return nil, err
			}

			return vall * valr, nil
		case nil:
			return nil, nil
		}
	} else {
		op := opMul{}
		return evalBinaryNumeric(left, right, tp, op)
	}
	return nil, fmt.Errorf("Unimplemented mul for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprQUO(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:

		case bool:
			vall, err := castInt(left)
			if err != nil {
				return nil, err
			}
			valr, err := castInt(right)
			if err != nil {
				return nil, err
			}

			return vall / valr, nil
		case nil:
			return nil, nil
		}
	} else {
		op := opQuo{}
		return evalBinaryNumeric(left, right, tp, op)
	}
	return nil, fmt.Errorf("Unimplemented / for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprREM(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:

		case bool:
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l % r, nil
		case nil:
			return nil, nil
		}
	} else {
		op := opRem{}
		return evalBinaryNumeric(left, right, tp, op)
	}
	return nil, fmt.Errorf("Unimplemented %% for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprAND(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:

		case bool:
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l & r, nil
		case nil:
			return nil, nil
		}
	} else {
		op := opAnd{}
		return evalBinaryNumeric(left, right, tp, op)
	}
	return nil, fmt.Errorf("Unimplemented & for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprOR(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:

		case bool:
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l | r, nil
		case nil:
			return nil, nil
		}
	} else {
		op := opOr{}
		return evalBinaryNumeric(left, right, tp, op)
	}

	return nil, fmt.Errorf("Unimplemented | for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprSHL(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:

		case bool:
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l << r, nil
		case nil:
			return nil, nil
		}
	} else {
		op := opShl{}
		return evalBinaryNumeric(left, right, tp, op)
	}
	return nil, fmt.Errorf("Unimplemented << for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprSHR(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:

		case bool:
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l >> r, nil
		case nil:
			return nil, nil
		}
	} else {
		op := opShr{}
		return evalBinaryNumeric(left, right, tp, op)
	}

	return nil, fmt.Errorf("Unimplemented >> for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprXOR(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:

		case bool:
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l ^ r, nil
		case nil:
			return nil, nil
		}
	} else {
		op := opXor{}
		return evalBinaryNumeric(left, right, tp, op)
	}

	return nil, fmt.Errorf("Unimplemented ^ for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprANDNOT(left interface{}, right interface{}) (interface{}, error) {
	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:

		case bool:
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l &^ r, nil
		case nil:
			return nil, nil
		}
	} else {
		op := opAndNot{}
		return evalBinaryNumeric(left, right, tp, op)
	}
	return nil, fmt.Errorf("Unimplemented &^ for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}
