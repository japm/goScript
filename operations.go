/*
The MIT License (MIT)
Copyright (c) 2016 Juan Pascual
*/
package goScript

type operation interface {
	OperF32F32(l float32, r float32) (interface{}, error)
	OperF64F64(l float64, r float64) (interface{}, error)
	OperI64I64(l int64, r int64) (interface{}, error)
	OperI32I32(l int32, r int32) (interface{}, error)
	OperI16I16(l int16, r int16) (interface{}, error)
	OperI8I8(l int8, r int8) (interface{}, error)

	OperUI64UI64(l uint64, r uint64) (interface{}, error)
	OperUI32UI32(l uint32, r uint32) (interface{}, error)
	OperUI16UI16(l uint16, r uint16) (interface{}, error)
	OperUI8UI8(l uint8, r uint8) (interface{}, error)
}

type opAdd struct {
}

func (op opAdd) OperF32F32(l float32, r float32) (interface{}, error) {
	return l + r, nil
}
func (op opAdd) OperF64F64(l float64, r float64) (interface{}, error) {
	return l + r, nil
}
func (op opAdd) OperI64I64(l int64, r int64) (interface{}, error) {
	return l + r, nil
}
func (op opAdd) OperI32I32(l int32, r int32) (interface{}, error) {
	return l + r, nil
}
func (op opAdd) OperI16I16(l int16, r int16) (interface{}, error) {
	return l + r, nil
}
func (op opAdd) OperI8I8(l int8, r int8) (interface{}, error) {
	return l + r, nil
}
func (op opAdd) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l + r, nil
}
func (op opAdd) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l + r, nil
}
func (op opAdd) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l + r, nil
}
func (op opAdd) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l + r, nil
}

type opSub struct {
}

func (op opSub) OperF32F32(l float32, r float32) (interface{}, error) {
	return l - r, nil
}
func (op opSub) OperF64F64(l float64, r float64) (interface{}, error) {
	return l - r, nil
}
func (op opSub) OperI64I64(l int64, r int64) (interface{}, error) {
	return l - r, nil
}
func (op opSub) OperI32I32(l int32, r int32) (interface{}, error) {
	return l - r, nil
}
func (op opSub) OperI16I16(l int16, r int16) (interface{}, error) {
	return l - r, nil
}
func (op opSub) OperI8I8(l int8, r int8) (interface{}, error) {
	return l - r, nil
}
func (op opSub) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l - r, nil
}
func (op opSub) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l - r, nil
}
func (op opSub) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l - r, nil
}
func (op opSub) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l - r, nil
}

type opMul struct {
}

func (op opMul) OperF32F32(l float32, r float32) (interface{}, error) {
	return l * r, nil
}
func (op opMul) OperF64F64(l float64, r float64) (interface{}, error) {
	return l * r, nil
}
func (op opMul) OperI64I64(l int64, r int64) (interface{}, error) {
	return l * r, nil
}
func (op opMul) OperI32I32(l int32, r int32) (interface{}, error) {
	return l * r, nil
}
func (op opMul) OperI16I16(l int16, r int16) (interface{}, error) {
	return l * r, nil
}
func (op opMul) OperI8I8(l int8, r int8) (interface{}, error) {
	return l * r, nil
}
func (op opMul) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l * r, nil
}
func (op opMul) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l * r, nil
}
func (op opMul) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l * r, nil
}
func (op opMul) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l * r, nil
}

type opQuo struct {
}

func (op opQuo) OperF32F32(l float32, r float32) (interface{}, error) {
	return l / r, nil
}
func (op opQuo) OperF64F64(l float64, r float64) (interface{}, error) {
	return l / r, nil
}
func (op opQuo) OperI64I64(l int64, r int64) (interface{}, error) {
	return l / r, nil
}
func (op opQuo) OperI32I32(l int32, r int32) (interface{}, error) {
	return l / r, nil
}
func (op opQuo) OperI16I16(l int16, r int16) (interface{}, error) {
	return l / r, nil
}
func (op opQuo) OperI8I8(l int8, r int8) (interface{}, error) {
	return l / r, nil
}
func (op opQuo) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l / r, nil
}
func (op opQuo) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l / r, nil
}
func (op opQuo) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l / r, nil
}
func (op opQuo) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l / r, nil
}

type opRem struct {
}

func (op opRem) OperF32F32(l float32, r float32) (interface{}, error) {
	li, err := castInt(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt(r)
	if err != nil {
		return nil, err
	}
	return li % ri, nil
}
func (op opRem) OperF64F64(l float64, r float64) (interface{}, error) {
	li, err := castInt64(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt64(r)
	if err != nil {
		return nil, err
	}
	return li % ri, nil
}

func (op opRem) OperI64I64(l int64, r int64) (interface{}, error) {
	return l % r, nil
}
func (op opRem) OperI32I32(l int32, r int32) (interface{}, error) {
	return l % r, nil
}
func (op opRem) OperI16I16(l int16, r int16) (interface{}, error) {
	return l % r, nil
}
func (op opRem) OperI8I8(l int8, r int8) (interface{}, error) {
	return l % r, nil
}
func (op opRem) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l % r, nil
}
func (op opRem) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l % r, nil
}
func (op opRem) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l % r, nil
}
func (op opRem) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l % r, nil
}

type opAnd struct {
}

func (op opAnd) OperF32F32(l float32, r float32) (interface{}, error) {
	li, err := castInt(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt(r)
	if err != nil {
		return nil, err
	}
	return li & ri, nil
}
func (op opAnd) OperF64F64(l float64, r float64) (interface{}, error) {
	li, err := castInt64(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt64(r)
	if err != nil {
		return nil, err
	}
	return li & ri, nil
}

func (op opAnd) OperI64I64(l int64, r int64) (interface{}, error) {
	return l & r, nil
}
func (op opAnd) OperI32I32(l int32, r int32) (interface{}, error) {
	return l & r, nil
}
func (op opAnd) OperI16I16(l int16, r int16) (interface{}, error) {
	return l & r, nil
}
func (op opAnd) OperI8I8(l int8, r int8) (interface{}, error) {
	return l & r, nil
}
func (op opAnd) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l & r, nil
}
func (op opAnd) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l & r, nil
}
func (op opAnd) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l & r, nil
}
func (op opAnd) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l & r, nil
}

type opOr struct {
}

func (op opOr) OperF32F32(l float32, r float32) (interface{}, error) {
	li, err := castInt(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt(r)
	if err != nil {
		return nil, err
	}
	return li | ri, nil
}
func (op opOr) OperF64F64(l float64, r float64) (interface{}, error) {
	li, err := castInt64(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt64(r)
	if err != nil {
		return nil, err
	}
	return li | ri, nil
}

func (op opOr) OperI64I64(l int64, r int64) (interface{}, error) {
	return l | r, nil
}
func (op opOr) OperI32I32(l int32, r int32) (interface{}, error) {
	return l | r, nil
}
func (op opOr) OperI16I16(l int16, r int16) (interface{}, error) {
	return l | r, nil
}
func (op opOr) OperI8I8(l int8, r int8) (interface{}, error) {
	return l | r, nil
}
func (op opOr) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l | r, nil
}
func (op opOr) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l | r, nil
}
func (op opOr) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l | r, nil
}
func (op opOr) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l | r, nil
}

type opShl struct {
}

func (op opShl) OperF32F32(l float32, r float32) (interface{}, error) {
	li, err := castInt(l)
	if err != nil {
		return nil, err
	}
	ri, err := castUint(r)
	if err != nil {
		return nil, err
	}
	return li << ri, nil
}
func (op opShl) OperF64F64(l float64, r float64) (interface{}, error) {
	li, err := castInt64(l)
	if err != nil {
		return nil, err
	}
	ri, err := castUint64(r)
	if err != nil {
		return nil, err
	}
	return li << ri, nil
}

func (op opShl) OperI64I64(l int64, r int64) (interface{}, error) {

	ri, err := castUint64(r)
	if err != nil {
		return nil, err
	}
	return l << ri, nil
}
func (op opShl) OperI32I32(l int32, r int32) (interface{}, error) {
	ri, err := castUint32(r)
	if err != nil {
		return nil, err
	}
	return l << ri, nil
}
func (op opShl) OperI16I16(l int16, r int16) (interface{}, error) {
	ri, err := castUint16(r)
	if err != nil {
		return nil, err
	}
	return l << ri, nil
}
func (op opShl) OperI8I8(l int8, r int8) (interface{}, error) {
	ri, err := castUint8(r)
	if err != nil {
		return nil, err
	}
	return l << ri, nil
}
func (op opShl) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l << r, nil
}
func (op opShl) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l << r, nil
}
func (op opShl) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l << r, nil
}
func (op opShl) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l << r, nil
}

type opShr struct {
}

func (op opShr) OperF32F32(l float32, r float32) (interface{}, error) {
	li, err := castInt(l)
	if err != nil {
		return nil, err
	}
	ri, err := castUint(r)
	if err != nil {
		return nil, err
	}
	return li >> ri, nil
}
func (op opShr) OperF64F64(l float64, r float64) (interface{}, error) {
	li, err := castInt64(l)
	if err != nil {
		return nil, err
	}
	ri, err := castUint64(r)
	if err != nil {
		return nil, err
	}
	return li >> ri, nil
}

func (op opShr) OperI64I64(l int64, r int64) (interface{}, error) {
	li, err := castInt64(l)
	if err != nil {
		return nil, err
	}
	ri, err := castUint64(r)
	if err != nil {
		return nil, err
	}
	return li >> ri, nil
}
func (op opShr) OperI32I32(l int32, r int32) (interface{}, error) {
	li, err := castInt32(l)
	if err != nil {
		return nil, err
	}
	ri, err := castUint32(r)
	if err != nil {
		return nil, err
	}
	return li >> ri, nil
}
func (op opShr) OperI16I16(l int16, r int16) (interface{}, error) {
	li, err := castInt16(l)
	if err != nil {
		return nil, err
	}
	ri, err := castUint16(r)
	if err != nil {
		return nil, err
	}
	return li >> ri, nil
}
func (op opShr) OperI8I8(l int8, r int8) (interface{}, error) {
	li, err := castInt8(l)
	if err != nil {
		return nil, err
	}
	ri, err := castUint8(r)
	if err != nil {
		return nil, err
	}
	return li >> ri, nil
}
func (op opShr) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l >> r, nil
}
func (op opShr) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l >> r, nil
}
func (op opShr) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l >> r, nil
}
func (op opShr) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l >> r, nil
}

type opXor struct {
}

func (op opXor) OperF32F32(l float32, r float32) (interface{}, error) {
	li, err := castInt(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt(r)
	if err != nil {
		return nil, err
	}
	return li ^ ri, nil
}
func (op opXor) OperF64F64(l float64, r float64) (interface{}, error) {
	li, err := castInt64(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt64(r)
	if err != nil {
		return nil, err
	}
	return li ^ ri, nil
}

func (op opXor) OperI64I64(l int64, r int64) (interface{}, error) {
	return l ^ r, nil
}
func (op opXor) OperI32I32(l int32, r int32) (interface{}, error) {
	return l ^ r, nil
}
func (op opXor) OperI16I16(l int16, r int16) (interface{}, error) {
	return l ^ r, nil
}
func (op opXor) OperI8I8(l int8, r int8) (interface{}, error) {
	return l ^ r, nil
}
func (op opXor) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l ^ r, nil
}
func (op opXor) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l ^ r, nil
}
func (op opXor) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l ^ r, nil
}
func (op opXor) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l ^ r, nil
}

type opAndNot struct {
}

func (op opAndNot) OperF32F32(l float32, r float32) (interface{}, error) {
	li, err := castInt(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt(r)
	if err != nil {
		return nil, err
	}
	return li &^ ri, nil
}
func (op opAndNot) OperF64F64(l float64, r float64) (interface{}, error) {
	li, err := castInt64(l)
	if err != nil {
		return nil, err
	}
	ri, err := castInt64(r)
	if err != nil {
		return nil, err
	}
	return li &^ ri, nil
}

func (op opAndNot) OperI64I64(l int64, r int64) (interface{}, error) {
	return l &^ r, nil
}
func (op opAndNot) OperI32I32(l int32, r int32) (interface{}, error) {
	return l &^ r, nil
}
func (op opAndNot) OperI16I16(l int16, r int16) (interface{}, error) {
	return l &^ r, nil
}
func (op opAndNot) OperI8I8(l int8, r int8) (interface{}, error) {
	return l &^ r, nil
}
func (op opAndNot) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l &^ r, nil
}
func (op opAndNot) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l &^ r, nil
}
func (op opAndNot) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l &^ r, nil
}
func (op opAndNot) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l &^ r, nil
}

type opLeq struct {
}

func (op opLeq) OperF32F32(l float32, r float32) (interface{}, error) {
	return l <= r, nil
}
func (op opLeq) OperF64F64(l float64, r float64) (interface{}, error) {
	return l <= r, nil
}
func (op opLeq) OperI64I64(l int64, r int64) (interface{}, error) {
	return l <= r, nil
}
func (op opLeq) OperI32I32(l int32, r int32) (interface{}, error) {
	return l <= r, nil
}
func (op opLeq) OperI16I16(l int16, r int16) (interface{}, error) {
	return l <= r, nil
}
func (op opLeq) OperI8I8(l int8, r int8) (interface{}, error) {
	return l <= r, nil
}
func (op opLeq) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l <= r, nil
}
func (op opLeq) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l <= r, nil
}
func (op opLeq) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l <= r, nil
}
func (op opLeq) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l <= r, nil
}

type opEql struct {
}

func (op opEql) OperF32F32(l float32, r float32) (interface{}, error) {
	return l == r, nil
}
func (op opEql) OperF64F64(l float64, r float64) (interface{}, error) {
	return l == r, nil
}
func (op opEql) OperI64I64(l int64, r int64) (interface{}, error) {
	return l == r, nil
}
func (op opEql) OperI32I32(l int32, r int32) (interface{}, error) {
	return l == r, nil
}
func (op opEql) OperI16I16(l int16, r int16) (interface{}, error) {
	return l == r, nil
}
func (op opEql) OperI8I8(l int8, r int8) (interface{}, error) {
	return l == r, nil
}
func (op opEql) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l == r, nil
}
func (op opEql) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l == r, nil
}
func (op opEql) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l == r, nil
}
func (op opEql) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l == r, nil
}

type opLss struct {
}

func (op opLss) OperF32F32(l float32, r float32) (interface{}, error) {
	return l < r, nil
}
func (op opLss) OperF64F64(l float64, r float64) (interface{}, error) {
	return l < r, nil
}
func (op opLss) OperI64I64(l int64, r int64) (interface{}, error) {
	return l < r, nil
}
func (op opLss) OperI32I32(l int32, r int32) (interface{}, error) {
	return l < r, nil
}
func (op opLss) OperI16I16(l int16, r int16) (interface{}, error) {
	return l < r, nil
}
func (op opLss) OperI8I8(l int8, r int8) (interface{}, error) {
	return l < r, nil
}
func (op opLss) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l < r, nil
}
func (op opLss) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l < r, nil
}
func (op opLss) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l < r, nil
}
func (op opLss) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l < r, nil
}

type opGtr struct {
}

func (op opGtr) OperF32F32(l float32, r float32) (interface{}, error) {
	return l > r, nil
}
func (op opGtr) OperF64F64(l float64, r float64) (interface{}, error) {
	return l > r, nil
}
func (op opGtr) OperI64I64(l int64, r int64) (interface{}, error) {
	return l > r, nil
}
func (op opGtr) OperI32I32(l int32, r int32) (interface{}, error) {
	return l > r, nil
}
func (op opGtr) OperI16I16(l int16, r int16) (interface{}, error) {
	return l > r, nil
}
func (op opGtr) OperI8I8(l int8, r int8) (interface{}, error) {
	return l > r, nil
}
func (op opGtr) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l > r, nil
}
func (op opGtr) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l > r, nil
}
func (op opGtr) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l > r, nil
}
func (op opGtr) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l > r, nil
}

type opGeq struct {
}

func (op opGeq) OperF32F32(l float32, r float32) (interface{}, error) {
	return l >= r, nil
}
func (op opGeq) OperF64F64(l float64, r float64) (interface{}, error) {
	return l >= r, nil
}
func (op opGeq) OperI64I64(l int64, r int64) (interface{}, error) {
	return l >= r, nil
}
func (op opGeq) OperI32I32(l int32, r int32) (interface{}, error) {
	return l >= r, nil
}
func (op opGeq) OperI16I16(l int16, r int16) (interface{}, error) {
	return l >= r, nil
}
func (op opGeq) OperI8I8(l int8, r int8) (interface{}, error) {
	return l >= r, nil
}
func (op opGeq) OperUI64UI64(l uint64, r uint64) (interface{}, error) {
	return l >= r, nil
}
func (op opGeq) OperUI32UI32(l uint32, r uint32) (interface{}, error) {
	return l >= r, nil
}
func (op opGeq) OperUI16UI16(l uint16, r uint16) (interface{}, error) {
	return l >= r, nil
}
func (op opGeq) OperUI8UI8(l uint8, r uint8) (interface{}, error) {
	return l >= r, nil
}
