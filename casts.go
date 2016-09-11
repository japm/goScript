//Package goScript, javascript Eval() for go
//The MIT License (MIT)
//Copyright (c) 2016 Juan Pascual

package goScript

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	leftType = iota
	rightType
)

const (
	tpNone = iota
	tpInt
	tpFloat
	tpString
	tpBool
	tpPointer
	tpNil
)

type typeDesc struct {
	Type         int
	Size         int
	Signed       bool
	PlatformSize bool
}

func (t typeDesc) IsNil() bool {
	return t.Type == tpNil
}

func (t typeDesc) IsNumeric() bool {
	return t.Type == tpInt || t.Type == tpFloat
}

func (t typeDesc) Float() bool {
	return t.Type == tpFloat
}

func (t typeDesc) Bool() bool {
	return t.Type == tpBool
}

func binaryOperType(left interface{}, right interface{}) (typeDesc, error) {
	lType := valType(left)

	rType := valType(right)

	//Left type governs type
	if !lType.IsNumeric() {
		if lType.IsNil() {
			return rType, nil
		}
		return lType, nil
	}

	//Left numeric, take the more general type
	lType.Signed = lType.Signed || rType.Signed
	if lType.Size < rType.Size {
		lType.Size = rType.Size
	}

	if rType.Type == tpFloat {
		lType.Type = tpFloat
	}

	lType.PlatformSize = lType.PlatformSize && rType.PlatformSize

	return lType, nil
}

func binaryOperTypeL(left interface{}, right interface{}) (typeDesc, error) {

	lType := valType(left)

	rType := valType(right)

	//Bool have precedence here
	if lType.Bool() {
		if !rType.IsNumeric() {
			return rType, nil
		}
		return lType, nil
	}
	if rType.Bool() {
		if !lType.IsNumeric() {
			return lType, nil
		}
		return rType, nil
	}

	if !lType.IsNumeric() {
		if lType.IsNil() {
			return rType, nil
		}
		return lType, nil
	}
	tp := typeDesc{}

	tp.Signed = lType.Signed || rType.Signed
	if lType.Size > rType.Size {
		tp.Size = lType.Size
	} else {
		tp.Size = rType.Size
	}
	if lType.Type == tpFloat || rType.Type == tpFloat {
		tp.Type = tpFloat
	} else {
		tp.Type = tpInt
	}
	tp.PlatformSize = lType.PlatformSize && rType.PlatformSize

	return tp, nil
}

func valType(value interface{}) typeDesc {

	switch value.(type) {
	case uint8:
		return typeDesc{tpInt, 8, false, false}
	case uint16:
		return typeDesc{tpInt, 16, false, false}
	case uint32:
		return typeDesc{tpInt, 32, false, false}
	case uint64:
		return typeDesc{tpInt, 64, false, false}
	case uint:
		return typeDesc{tpInt, strconv.IntSize, false, true}
	case int8:
		return typeDesc{tpInt, 8, true, false}
	case int16:
		return typeDesc{tpInt, 16, true, false}
	case int32:
		return typeDesc{tpInt, 32, true, false}
	case int64:
		return typeDesc{tpInt, 64, true, false}
	case int:
		return typeDesc{tpInt, strconv.IntSize, true, true}
	case float32:
		return typeDesc{tpFloat, 32, true, false}
	case float64:
		return typeDesc{tpFloat, 64, true, false}
	case string:
		return typeDesc{tpString, 0, false, false}
	case bool:
		return typeDesc{tpBool, 0, false, false}
	case nil:
		return typeDesc{tpNil, 0, false, false}
	default:
		return typeDesc{tpPointer, 0, false, false}
	}
	//return typeDesc{tpNone, 0, false, false}, fmt.Errorf("Unimplemented type size", reflect.TypeOf(value))
}

func castUint8(value interface{}) (uint8, error) {
	switch value.(type) {
	case uint8:
		return value.(uint8), nil
	case uint16:
		return uint8(value.(uint16)), nil
	case uint32:
		return uint8(value.(uint32)), nil
	case uint64:
		return uint8(value.(uint64)), nil
	case uint:
		return uint8(value.(uint)), nil
	case int8:
		return uint8(value.(int8)), nil
	case int16:
		return uint8(value.(int16)), nil
	case int32:
		return uint8(value.(int32)), nil
	case int64:
		return uint8(value.(int64)), nil
	case int:
		return uint8(value.(int)), nil
	case float32:
		return uint8(value.(float32)), nil
	case float64:
		return uint8(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return uint8(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return uint8(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to uint8 for type %s", reflect.TypeOf(value))
}

func castUint16(value interface{}) (uint16, error) {
	switch value.(type) {
	case uint8:
		return uint16(value.(uint8)), nil
	case uint16:
		return value.(uint16), nil
	case uint32:
		return uint16(value.(uint32)), nil
	case uint64:
		return uint16(value.(uint64)), nil
	case uint:
		return uint16(value.(uint)), nil
	case int8:
		return uint16(value.(int8)), nil
	case int16:
		return uint16(value.(int16)), nil
	case int32:
		return uint16(value.(int32)), nil
	case int64:
		return uint16(value.(int64)), nil
	case int:
		return uint16(value.(int)), nil
	case float32:
		return uint16(value.(float32)), nil
	case float64:
		return uint16(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return uint16(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return uint16(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to uint16 for type %s", reflect.TypeOf(value))
}

func castUint32(value interface{}) (uint32, error) {
	switch value.(type) {
	case uint8:
		return uint32(value.(uint8)), nil
	case uint16:
		return uint32(value.(uint16)), nil
	case uint32:
		return value.(uint32), nil
	case uint64:
		return uint32(value.(uint64)), nil
	case uint:
		return uint32(value.(uint)), nil
	case int8:
		return uint32(value.(int8)), nil
	case int16:
		return uint32(value.(int16)), nil
	case int32:
		return uint32(value.(int32)), nil
	case int64:
		return uint32(value.(int64)), nil
	case int:
		return uint32(value.(int)), nil
	case float32:
		return uint32(value.(float32)), nil
	case float64:
		return uint32(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return uint32(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return uint32(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to uint32 for type %s", reflect.TypeOf(value))
}

func castUint(value interface{}) (uint, error) {
	switch value.(type) {
	case uint8:
		return uint(value.(uint8)), nil
	case uint16:
		return uint(value.(uint16)), nil
	case uint32:
		return uint(value.(uint32)), nil
	case uint64:
		return uint(value.(uint64)), nil
	case uint:
		return value.(uint), nil
	case int8:
		return uint(value.(int8)), nil
	case int16:
		return uint(value.(int16)), nil
	case int32:
		return uint(value.(int32)), nil
	case int64:
		return uint(value.(int64)), nil
	case int:
		return uint(value.(int)), nil
	case float32:
		return uint(value.(float32)), nil
	case float64:
		return uint(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return uint(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return uint(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to uint for type %s", reflect.TypeOf(value))
}

func castUint64(value interface{}) (uint64, error) {
	switch value.(type) {
	case uint8:
		return uint64(value.(uint8)), nil
	case uint16:
		return uint64(value.(uint16)), nil
	case uint32:
		return uint64(value.(uint32)), nil
	case uint64:
		return value.(uint64), nil
	case uint:
		return uint64(value.(uint)), nil
	case int8:
		return uint64(value.(int8)), nil
	case int16:
		return uint64(value.(int16)), nil
	case int:
		return uint64(value.(int)), nil
	case int64:
		return uint64(value.(int64)), nil
	case float32:
		return uint64(value.(float32)), nil
	case float64:
		return uint64(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return uint64(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return uint64(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to uint64 for type %s", reflect.TypeOf(value))
}

func castInt8(value interface{}) (int8, error) {
	switch value.(type) {
	case uint8:
		return int8(value.(uint8)), nil
	case uint16:
		return int8(value.(uint16)), nil
	case uint32:
		return int8(value.(uint32)), nil
	case uint64:
		return int8(value.(uint64)), nil
	case uint:
		return int8(value.(uint)), nil
	case int8:
		return value.(int8), nil
	case int16:
		return int8(value.(int16)), nil
	case int32:
		return int8(value.(int32)), nil
	case int64:
		return int8(value.(int64)), nil
	case int:
		return int8(value.(int)), nil
	case float32:
		return int8(value.(float32)), nil
	case float64:
		return int8(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return int8(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return int8(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to int8 for type %s", reflect.TypeOf(value))
}

func castInt16(value interface{}) (int16, error) {
	switch value.(type) {
	case uint8:
		return int16(value.(uint8)), nil
	case uint16:
		return int16(value.(uint16)), nil
	case uint32:
		return int16(value.(uint32)), nil
	case uint64:
		return int16(value.(uint64)), nil
	case uint:
		return int16(value.(uint)), nil
	case int8:
		return int16(value.(int8)), nil
	case int16:
		return value.(int16), nil
	case int32:
		return int16(value.(int32)), nil
	case int64:
		return int16(value.(int64)), nil
	case int:
		return int16(value.(int)), nil
	case float32:
		return int16(value.(float32)), nil
	case float64:
		return int16(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return int16(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return int16(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to int16 for type %s", reflect.TypeOf(value))
}

func castInt(value interface{}) (int, error) {
	switch value.(type) {
	case uint8:
		return int(value.(uint8)), nil
	case uint16:
		return int(value.(uint16)), nil
	case uint32:
		return int(value.(uint32)), nil
	case uint64:
		return int(value.(uint64)), nil
	case uint:
		return int(value.(uint)), nil
	case int8:
		return int(value.(int8)), nil
	case int16:
		return int(value.(int16)), nil
	case int32:
		return int(value.(int32)), nil
	case int64:
		return int(value.(int64)), nil
	case int:
		return value.(int), nil
	case float32:
		return int(value.(float32)), nil
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
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to int for type %s", reflect.TypeOf(value))
}

func castIntFromBool(value bool) (int, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func castInt32(value interface{}) (int32, error) {
	switch value.(type) {
	case uint8:
		return int32(value.(uint8)), nil
	case uint16:
		return int32(value.(uint16)), nil
	case uint32:
		return int32(value.(uint32)), nil
	case uint64:
		return int32(value.(uint64)), nil
	case uint:
		return int32(value.(uint)), nil
	case int8:
		return int32(value.(int8)), nil
	case int16:
		return int32(value.(int16)), nil
	case int32:
		return value.(int32), nil
	case int64:
		return int32(value.(int64)), nil
	case int:
		return int32(value.(int)), nil
	case float32:
		return int32(value.(float32)), nil
	case float64:
		return int32(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return int32(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return int32(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to uint32 for type %s", reflect.TypeOf(value))
}

func castInt64(value interface{}) (int64, error) {
	var ret int64
	var err error
	switch value.(type) {
	case uint8:
		ret, err = int64(value.(uint8)), nil
	case uint16:
		ret, err = int64(value.(uint16)), nil
	case uint32:
		ret, err = int64(value.(uint32)), nil
	case uint64:
		ret, err = int64(value.(uint64)), nil
	case uint:
		ret, err = int64(value.(uint)), nil
	case int8:
		ret, err = int64(value.(int8)), nil
	case int16:
		ret, err = int64(value.(int16)), nil
	case int32:
		ret, err = int64(value.(int32)), nil
	case int64:
		ret, err = value.(int64), nil
	case int:
		ret, err = int64(value.(int)), nil
	case float32:
		ret, err = int64(value.(float32)), nil
	case float64:
		ret, err = int64(value.(float64)), nil
	case string:
		vali, e := strconv.ParseInt(value.(string), 10, 64)
		if e == nil {
			ret, err = int64(vali), nil
			break
		}
		valf, e := strconv.ParseFloat(value.(string), 10)
		if e != nil {
			ret, err = 0, e
			break
		}
		ret, err = int64(valf), nil
	case bool:
		if value.(bool) {
			ret, err = 1, nil
		} else {
			ret, err = 0, nil
		}
	case nil:
		ret, err = 0, nil
	default:
		ret, err = 0, fmt.Errorf("Unimplemented cast to int64 for type %s", reflect.TypeOf(value))
	}
	return ret, err
}

func castFloat32(value interface{}) (float32, error) {
	switch value.(type) {
	case uint8:
		return float32(value.(uint8)), nil
	case uint16:
		return float32(value.(uint16)), nil
	case uint32:
		return float32(value.(uint32)), nil
	case uint64:
		return float32(value.(uint64)), nil
	case uint:
		return float32(value.(uint)), nil
	case int8:
		return float32(value.(int8)), nil
	case int16:
		return float32(value.(int16)), nil
	case int32:
		return float32(value.(int32)), nil
	case int64:
		return float32(value.(int64)), nil
	case int:
		return float32(value.(int)), nil
	case float32:
		return value.(float32), nil
	case float64:
		return float32(value.(float64)), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return float32(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return float32(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to float32 for type %s", reflect.TypeOf(value))
}

func castFloat64(value interface{}) (float64, error) {
	switch value.(type) {
	case uint8:
		return float64(value.(uint8)), nil
	case uint16:
		return float64(value.(uint16)), nil
	case uint32:
		return float64(value.(uint32)), nil
	case uint64:
		return float64(value.(uint64)), nil
	case uint:
		return float64(value.(uint)), nil
	case int8:
		return float64(value.(int8)), nil
	case int16:
		return float64(value.(int16)), nil
	case int32:
		return float64(value.(int32)), nil
	case int64:
		return float64(value.(int64)), nil
	case int:
		return float64(value.(int)), nil
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return value.(float64), nil
	case string:
		vali, err := strconv.ParseInt(value.(string), 10, 64)
		if err == nil {
			return float64(vali), nil
		}
		valf, err := strconv.ParseFloat(value.(string), 10)
		if err != nil {
			return 0, err
		}
		return float64(valf), nil
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	}
	return 0, fmt.Errorf("Unimplemented cast to int64 for type %s", reflect.TypeOf(value))
}

func castBool(value interface{}) (bool, error) {
	switch value.(type) {
	case uint8:
		return value.(uint8) != uint8(0), nil
	case uint16:
		return value.(uint16) != uint16(0), nil
	case uint32:
		return value.(uint32) != uint32(0), nil
	case uint64:
		return value.(uint64) != uint64(0), nil
	case uint:
		return value.(uint) != uint(0), nil
	case int8:
		return value.(int8) != int8(0), nil
	case int16:
		return value.(int16) != int16(0), nil
	case int32:
		return value.(int32) != 0, nil
	case int64:
		return value.(int64) != int64(0), nil
	case int:
		return value.(int) != 0, nil
	case float32:
		return value.(float32) != float32(0), nil
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

func castString(value interface{}) (string, error) {
	switch value.(type) {
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10), nil
	case uint64:
		return strconv.FormatUint(value.(uint64), 10), nil
	case uint:
		return strconv.FormatUint(uint64(value.(uint)), 10), nil
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10), nil
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10), nil
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10), nil
	case int64:
		return strconv.FormatInt(value.(int64), 10), nil
	case int:
		return strconv.FormatInt(int64(value.(int)), 10), nil
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64), nil
	case string:
		return value.(string), nil
	case bool:
		return strconv.FormatBool(value.(bool)), nil
	case nil:
		return "", nil
	}
	return "", fmt.Errorf("Unimplemented cast to string for type %s", reflect.TypeOf(value))
}
