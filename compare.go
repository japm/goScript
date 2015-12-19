package evalGo

import (
	"fmt"
	"reflect"
)

func evalBinaryExprLEQ(left interface{}, right interface{}) (interface{}, error) {
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
			return left.(string) <= val, nil

		case bool:

		case nil:
			return nil, nil
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
				return l <= r, nil
			}
			l, err := castFloat64(left)
			if err != nil {
				return nil, err
			}
			r, err := castFloat64(right)
			if err != nil {
				return nil, err
			}
			return l <= r, nil

		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l <= r, nil
		} else if tp.Size == 32 {
			l, err := castInt32(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt32(right)
			if err != nil {
				return nil, err
			}
			return l <= r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l <= r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l <= r, nil
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
			return l <= r, nil
		} else if tp.Size == 32 {
			l, err := castUint32(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint32(right)
			if err != nil {
				return nil, err
			}
			return l <= r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l <= r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l <= r, nil
		}
	}
	return nil, fmt.Errorf("Unimplemented <= for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprNEQ(left interface{}, right interface{}) (interface{}, error) {
	val, err := evalBinaryExprEQL(left, right)
	if err != nil {
		return nil, err
	}
	return !val.(bool), nil
}

func evalBinaryExprEQL(left interface{}, right interface{}) (interface{}, error) {
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
				return l == r, nil
			}
			l, err := castFloat64(left)
			if err != nil {
				return nil, err
			}
			r, err := castFloat64(right)
			if err != nil {
				return nil, err
			}
			return l == r, nil

		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l == r, nil
		} else if tp.Size == 32 {
			l, err := castInt32(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt32(right)
			if err != nil {
				return nil, err
			}
			return l == r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l == r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l == r, nil
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
			return l == r, nil
		} else if tp.Size == 32 {
			l, err := castUint32(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint32(right)
			if err != nil {
				return nil, err
			}
			return l == r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l == r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l == r, nil
		}
	}
	return nil, fmt.Errorf("Unimplemented == for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprLSS(left interface{}, right interface{}) (interface{}, error) {
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
			return left.(string) < val, nil

		case bool:

		case nil:

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
				return l < r, nil
			}
			l, err := castFloat64(left)
			if err != nil {
				return nil, err
			}
			r, err := castFloat64(right)
			if err != nil {
				return nil, err
			}
			return l < r, nil

		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l < r, nil
		} else if tp.Size == 32 {
			l, err := castInt32(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt32(right)
			if err != nil {
				return nil, err
			}
			return l < r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l < r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l < r, nil
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
			return l < r, nil
		} else if tp.Size == 32 {
			l, err := castUint32(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint32(right)
			if err != nil {
				return nil, err
			}
			return l < r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l < r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l < r, nil
		}
	}
	return nil, fmt.Errorf("Unimplemented < for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprGTR(left interface{}, right interface{}) (interface{}, error) {
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
			return left.(string) > val, nil

		case bool:

		case nil:

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
				return l > r, nil
			}
			l, err := castFloat64(left)
			if err != nil {
				return nil, err
			}
			r, err := castFloat64(right)
			if err != nil {
				return nil, err
			}
			return l > r, nil

		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l > r, nil
		} else if tp.Size == 32 {
			l, err := castInt32(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt32(right)
			if err != nil {
				return nil, err
			}
			return l > r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l > r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l > r, nil
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
			return l > r, nil
		} else if tp.Size == 32 {
			l, err := castUint32(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint32(right)
			if err != nil {
				return nil, err
			}
			return l > r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l > r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l > r, nil
		}
	}
	return nil, fmt.Errorf("Unimplemented > for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprGEQ(left interface{}, right interface{}) (interface{}, error) {
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
			return left.(string) >= val, nil

		case bool:

		case nil:

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
				return l >= r, nil
			}
			l, err := castFloat64(left)
			if err != nil {
				return nil, err
			}
			r, err := castFloat64(right)
			if err != nil {
				return nil, err
			}
			return l >= r, nil
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l >= r, nil
		} else if tp.Size == 32 {
			l, err := castInt32(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt32(right)
			if err != nil {
				return nil, err
			}
			return l >= r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l >= r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l >= r, nil
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
			return l >= r, nil
		} else if tp.Size == 32 {
			l, err := castUint32(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint32(right)
			if err != nil {
				return nil, err
			}
			return l >= r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l >= r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l >= r, nil
		}
	}
	return nil, fmt.Errorf("Unimplemented >= for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}
