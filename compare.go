/*
Package goScript
The MIT License (MIT)
Copyright (c) 2016 Juan Pascual
*/
package goScript

/*
import (
	"fmt"
	"reflect"
)

func evalBinaryExprNEQ(left interface{}, right interface{}) (value, error) {
	val, err := evalBinaryExprEQL(left, right)
	if err != nil {
		return nil, err
	}
	return !val.(bool), nil
}


func evalBinaryExprEQL(left interface{}, right interface{}) (value, error) {
	tp, e := binaryOperTypeL(left, right)
	if e != nil {
		return nil, e
	}
	if tp.IsNil() {
		return nil, nil
	}
	if tp.Bool() {

		l, err := castBool(left)
		if err != nil {
			return nil, err
		}

		r, err := castBool(right)
		if err != nil {
			return nil, err
		}

		return l == r, nil

	} else if !tp.IsNumeric() {
		switch left.(type) {
		case string:
			val, err := castString(right)
			if err != nil {
				return nil, err
			}
			return left.(string) == val, nil
		case nil:
			switch right.(type) {
			case nil:
				return true, nil
			}
			return nil, nil
		}
	} else {
		op := opEql{}
		return evalBinary(left, right, tp, op)
	}
	return nil, fmt.Errorf("Unimplemented == for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}
*/
