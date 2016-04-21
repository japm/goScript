/*
The MIT License (MIT)
Copyright (c) 2016 Juan Pascual
*/
package goScript

import (
	"fmt"
	"reflect"
)

func isEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, c := range s {
		if c != ' ' && c != '\t' {
			return false
		}
	}
	return true
}

func evalBinary(left interface{}, right interface{}, tp typeDesc, oper operation) (interface{}, error) {
	if tp.IsNil() {
		return nil, nil
	}
	if !tp.IsNumeric() {
		switch left.(type) {
		case string:
			return oper.OperStrInterf(left.(string), right)
		case bool:
			return oper.OperBoolInterf(left.(bool), right)
		case nil:
			return oper.OperNilLeft(right)
		}
	} else if tp.Signed {
		if tp.Float() {
			if tp.Size == 32 {
				l, err := castFloat32(left)
				if err != nil {
					return nil, err
				}
				r, err := castFloat32(right)
				if err != nil {
					return nil, err
				}
				return oper.OperF32F32(l, r)
			}
			l, err := castFloat64(left)
			if err != nil {
				return nil, err
			}
			r, err := castFloat64(right)
			if err != nil {
				return nil, err
			}
			return oper.OperF64F64(l, r)

		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return oper.OperI64I64(l, r)

		} else if tp.Size == 32 {
			l, err := castInt32(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt32(right)
			if err != nil {
				return nil, err
			}
			return oper.OperI32I32(l, r)
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return oper.OperI16I16(l, r)
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return oper.OperI8I8(l, r)
		}
	} else {
		if tp.Size == 64 {
			l, err := castUint64(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint64(right)
			if err != nil {
				return nil, err
			}
			return oper.OperUI64UI64(l, r)
		} else if tp.Size == 32 {
			l, err := castUint32(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint32(right)
			if err != nil {
				return nil, err
			}
			return oper.OperUI32UI32(l, r)
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return oper.OperUI16UI16(l, r)
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return oper.OperUI8UI8(l, r)
		}
	}
	return nil, fmt.Errorf("Unimplemented op for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}
