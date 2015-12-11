package evalGo

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
				return l + r, nil
			} else {
				l, err := castFloat64(left)
				if err != nil {
					return nil, err
				}
				r, err := castFloat64(right)
				if err != nil {
					return nil, err
				}
				return l + r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l + r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l + r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l + r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l + r, nil
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
			return l + r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l + r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l + r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l + r, nil
		}
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
				return l - r, nil
			} else {
				l, err := castFloat64(left)
				if err != nil {
					return nil, err
				}
				r, err := castFloat64(right)
				if err != nil {
					return nil, err
				}
				return l - r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l - r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l - r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l - r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l - r, nil
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
			return l - r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l - r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l - r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l - r, nil
		}
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
				return l * r, nil
			} else {
				l, err := castFloat64(left)
				if err != nil {
					return nil, err
				}
				r, err := castFloat64(right)
				if err != nil {
					return nil, err
				}
				return l * r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l * r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l * r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l * r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l * r, nil
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
			return l * r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l * r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l * r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l * r, nil
		}
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
				return l / r, nil
			} else {
				l, err := castFloat64(left)
				if err != nil {
					return nil, err
				}
				r, err := castFloat64(right)
				if err != nil {
					return nil, err
				}
				return l / r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l / r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l / r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l / r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l / r, nil
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
			return l / r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l / r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l / r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l / r, nil
		}
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
	} else if tp.Signed {
		if tp.Float() {
			if tp.Size == 32 {
				l, err := castInt(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt(right)
				if err != nil {
					return nil, err
				}
				return l % r, nil
			} else {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt64(right)
				if err != nil {
					return nil, err
				}
				return l % r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l % r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l % r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l % r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l % r, nil
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
			return l % r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l % r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l % r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l % r, nil
		}
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
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l & r, nil
		case nil:
			return nil, nil
		}
	} else if tp.Signed {
		if tp.Float() {
			if tp.Size == 32 {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt64(right)
				if err != nil {
					return nil, err
				}
				return l & r, nil
			} else {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt64(right)
				if err != nil {
					return nil, err
				}
				return l & r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l & r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l & r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l & r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l & r, nil
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
			return l & r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l & r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l & r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l & r, nil
		}
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
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l | r, nil
		case nil:
			return nil, nil
		}
	} else if tp.Signed {
		if tp.Float() {
			if tp.Size == 32 {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt64(right)
				if err != nil {
					return nil, err
				}
				return l | r, nil
			} else {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt64(right)
				if err != nil {
					return nil, err
				}
				return l | r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l | r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l | r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l | r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l | r, nil
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
			return l | r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l | r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l | r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l | r, nil
		}
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
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint64(right)
			if err != nil {
				return nil, err
			}
			return l << r, nil
		case nil:
			return nil, nil
		}
	} else if tp.Signed {
		if tp.Float() {
			if tp.Size == 32 {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castUint64(right)
				if err != nil {
					return nil, err
				}
				return l << r, nil
			} else {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castUint64(right)
				if err != nil {
					return nil, err
				}
				return l << r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint64(right)
			if err != nil {
				return nil, err
			}
			return l << r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l << r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l << r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l << r, nil
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
			return l << r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l << r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l << r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l << r, nil
		}
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
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint64(right)
			if err != nil {
				return nil, err
			}
			return l >> r, nil
		case nil:
			return nil, nil
		}
	} else if tp.Signed {
		if tp.Float() {
			if tp.Size == 32 {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castUint64(right)
				if err != nil {
					return nil, err
				}
				return l >> r, nil
			} else {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castUint64(right)
				if err != nil {
					return nil, err
				}
				return l >> r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint64(right)
			if err != nil {
				return nil, err
			}
			return l >> r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l >> r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l >> r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l >> r, nil
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
			return l >> r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l >> r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l >> r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l >> r, nil
		}
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
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l ^ r, nil
		case nil:
			return nil, nil
		}
	} else if tp.Signed {
		if tp.Float() {
			if tp.Size == 32 {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt64(right)
				if err != nil {
					return nil, err
				}
				return l ^ r, nil
			} else {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt64(right)
				if err != nil {
					return nil, err
				}
				return l ^ r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l ^ r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l ^ r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l ^ r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l ^ r, nil
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
			return l ^ r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l ^ r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l ^ r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l ^ r, nil
		}
	}

	return nil, fmt.Errorf("Unimplemented ^ for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprAND_NOT(left interface{}, right interface{}) (interface{}, error) {
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
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l &^ r, nil
		case nil:
			return nil, nil
		}
	} else if tp.Signed {
		if tp.Float() {
			if tp.Size == 32 {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt64(right)
				if err != nil {
					return nil, err
				}
				return l &^ r, nil
			} else {
				l, err := castInt64(left)
				if err != nil {
					return nil, err
				}
				r, err := castInt64(right)
				if err != nil {
					return nil, err
				}
				return l &^ r, nil
			}
		} else if tp.Size == 64 {
			l, err := castInt64(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt64(right)
			if err != nil {
				return nil, err
			}
			return l &^ r, nil
		} else if tp.Size == 32 {
			l, err := castInt(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt(right)
			if err != nil {
				return nil, err
			}
			return l &^ r, nil
		} else if tp.Size == 16 {
			l, err := castInt16(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt16(right)
			if err != nil {
				return nil, err
			}
			return l &^ r, nil
		} else if tp.Size == 8 {
			l, err := castInt8(left)
			if err != nil {
				return nil, err
			}
			r, err := castInt8(right)
			if err != nil {
				return nil, err
			}
			return l &^ r, nil
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
			return l &^ r, nil
		} else if tp.Size == 32 {
			l, err := castUint(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint(right)
			if err != nil {
				return nil, err
			}
			return l &^ r, nil
		} else if tp.Size == 16 {
			l, err := castUint16(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint16(right)
			if err != nil {
				return nil, err
			}
			return l &^ r, nil
		} else if tp.Size == 8 {
			l, err := castUint8(left)
			if err != nil {
				return nil, err
			}
			r, err := castUint8(right)
			if err != nil {
				return nil, err
			}
			return l &^ r, nil
		}
	}

	return nil, fmt.Errorf("Unimplemented &^ for types %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}
