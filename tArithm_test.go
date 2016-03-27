package goScript

import (
	"strconv"
	"testing"
)

func TestAddConstInt(t *testing.T) {

	val, err := Eval("1 + 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddConstFloat(t *testing.T) {

	val, err := Eval("1.15 + 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3.4 {
		t.Error("Expected 3.4 get ", val)
	}
}

func TestAddConstFloatInt1(t *testing.T) {

	val, err := Eval("1 + 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3.25 {
		t.Error("Expected 3.25 get ", val)
	}
}

func TestAddConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 + 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3.25 {
		t.Error("Expected 3.25 get ", val)
	}
}

func TestAddConstStringInt(t *testing.T) {

	val, err := Eval("\"a\" + 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "a1" {
		t.Error("Expected a1 get ", val)
	}
}

func TestAddConstIntString(t *testing.T) {

	_, err := Eval("1 + \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAddConstIntString2(t *testing.T) {

	val, err := Eval("1 + \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddConstStringString(t *testing.T) {

	val, err := Eval("\"1\" + \"2\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "12" {
		t.Error("Expected 12 get ", val)
	}
}

func TestAddConstIntBool(t *testing.T) {

	val, err := Eval("2 + true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestAddUintUint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(uint32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(uint64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestAddUintUint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUintFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUintFloat32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint32Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint32(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float32) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint64Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint8Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint8Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint8(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint8) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint8Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestAddUint8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint16Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint16Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint16Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestAddUint16Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddUint64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a + b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAddIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestAddIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestAddInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a + b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAddInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestAddFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a + b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAddFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestAddStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z2" {
		t.Error("Expected z2 get ", val)
	}
}

func TestAddStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z2" {
		t.Error("Expected z2 get ", val)
	}
}

func TestAddStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z2" {
		t.Error("Expected z2 get ", val)
	}
}

func TestAddStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "xz" {
		t.Error("Expected xz get ", val)
	}
}

func TestAddStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "ztrue" {
		t.Error("Expected ztrue get ", val)
	}
}

func TestAddStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = false

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "zfalse" {
		t.Error("Expected ztrue get ", val)
	}
}
func TestAddBoolBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = true

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestAddBoolBool1(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = false

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestAddBoolBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = true

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestAddBoolBool3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = false

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivConstInt(t *testing.T) {

	val, err := Eval("2 / 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivConstFloat(t *testing.T) {

	val, err := Eval("1.15 / 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 0.5111111111111111 {
		t.Error("Expected 0.5111111111111111  get ", val)
	}
}

func TestDivConstFloatInt1(t *testing.T) {

	val, err := Eval("2 / 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 0.8888888888888888 {
		t.Error("Expected 0.8888888888888888  get ", val)
	}
}

func TestDivConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 / 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1.125 {
		t.Error("Expected 1.125 get ", val)
	}
}

func TestDivConstStringInt(t *testing.T) {

	_, err := Eval("\"a\" / 3", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivConstIntString(t *testing.T) {

	_, err := Eval("1 / \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivConstIntString2(t *testing.T) {

	val, err := Eval("2 / \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivConstStringString(t *testing.T) {

	val, err := Eval("\"2\" / \"2\"", nil)

	if err != nil {
		t.Error("Error not nil unexpected")
	}

	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}

}

func TestDivIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = 2

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}

}

func TestDivIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = int64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = float64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = "z"

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestDivIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = 2

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = int64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(1)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestDivInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = 2

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = int64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = float64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = "z"

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestDivFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = true

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = false

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivUintUint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(uint32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(uint64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestDivUintUint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivUintFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != float64(0.5) {
		t.Error("Expected 0.5 get ", val)
	}
}

func TestDivUintFloat32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}

	if strconv.IntSize == 32 {
		if val.(float32) != float32(0.5) {
			t.Error("Expected 0.5 get ", val)
		}
	} else {
		if val.(float64) != float64(0.5) {
			t.Error("Expected 0.5 get ", val)
		}
	}
}

func TestDivUint64Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != float64(0.5) {
		t.Error("Expected 0.5 get ", val)
	}
}

func TestDivUint8Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint16) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivUint8Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint8(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint8) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivUint8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivUint8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivUint8Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(int64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestDivUint8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivUint16Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivUint16Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivUint16Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(int64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestDivUint16Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivInt8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivInt8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivInt8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivUint64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivBoolBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = true

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestDivBoolBool1(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = false

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected error")
	}

}

func TestDivBoolBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = true

	val, err := Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivBoolBool3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = false

	_, err := Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected error")
	}

}

func TestMulConstInt(t *testing.T) {

	val, err := Eval("2 * 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulConstFloat(t *testing.T) {

	val, err := Eval("1.15 * 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2.5875 {
		t.Error("Expected 2.5875 get ", val)
	}
}

func TestMulConstFloatInt1(t *testing.T) {

	val, err := Eval("2 * 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4.50 {
		t.Error("Expected 4.50 get ", val)
	}
}

func TestMulConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 * 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4.50 {
		t.Error("Expected 4.50 get ", val)
	}
}

func TestMulConstStringInt(t *testing.T) {

	val, err := Eval("\"a\" * 3", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "aaa" {
		t.Error("Expected aaa get ", val)
	}
}

func TestMulConstIntString(t *testing.T) {

	_, err := Eval("1 * \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulConstIntString2(t *testing.T) {

	val, err := Eval("2 * \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulConstStringString(t *testing.T) {

	val, err := Eval("\"2\" * \"2\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "22" {
		t.Error("Expected 22 get ", val)
	}
}

func TestMulIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = 2

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 4 {
			t.Error("Expected 4 get ", val)
		}
	} else {
		if val.(int64) != 4 {
			t.Error("Expected 4 get ", val)
		}
	}
}

func TestMulIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = int64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = float64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = "z"

	_, err := Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestMulIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(int64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestMulInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = 2

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = int64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 3 get ", val)
	}
}

func TestMulInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = float64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestMulFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = 2

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = int64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = float64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = "z"

	_, err := Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestMulStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "zz" {
		t.Error("Expected zz get ", val)
	}
}

func TestMulStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "zz" {
		t.Error("Expected zz get ", val)
	}
}

func TestMulStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "zz" {
		t.Error("Expected zz get ", val)
	}
}

func TestMulStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulUintUint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(uint32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(uint64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestMulUintUint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUintFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUintFloat32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(float32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(float64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestMulUint64Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUint8Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint16) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUint8Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint8(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint8) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUint8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUint8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUint8Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestMulUint8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUint16Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUint16Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUint16Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestMulUint16Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulInt8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulInt8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulInt8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulUint64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulBoolBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = true

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestMulBoolBool1(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = false

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestMulBoolBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = true

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestMulBoolBool3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = false

	val, err := Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestRemConstInt(t *testing.T) {

	val, err := Eval("1 % 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestRemConstFloat(t *testing.T) {

	val, err := Eval("1.15 % 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemConstFloatInt1(t *testing.T) {

	val, err := Eval("1 % 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 % 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestRemConstStringInt(t *testing.T) {

	_, err := Eval("\"a\" % 1", nil)

	if err == nil {
		t.Error("Expected error on string % int")
	}
}

func TestRemConstIntString(t *testing.T) {

	_, err := Eval("1 % \"a\" ", nil)

	if err == nil {
		t.Error("Expected error on int % string")
	}
}

func TestRemConstIntString2(t *testing.T) {

	val, err := Eval("1 % \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestRemConstStringString(t *testing.T) {

	val, err := Eval("\"1\" % \"2\"", nil)

	if err != nil {
		t.Error("Unexpected error on string % string")
	}

	if val.(int) != 1 {
		t.Error("Expected 3 get ", val)
	}

}

func TestRemConstIntBool(t *testing.T) {

	val, err := Eval("2 % true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestRemIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestRemIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestRemIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(int64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestRemIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in int % false")
	}
}

func TestRemInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestRemInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestRemInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in int % false")
	}

}

func TestRemFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestRemFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestRemFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in int % false")
	}
}

func TestRemStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in string % int")
	}
}

func TestRemStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in string % int64")
	}
}

func TestRemStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in string % flaot64")
	}
}

func TestRemStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := Eval("a % b", ctxt)
	if err == nil {
		t.Error("Expected error in string % int")
	}
}

func TestRemStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	_, err := Eval("a % b", ctxt)
	if err == nil {
		t.Error("Expected error in string % bool")
	}
}

func TestRemStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = false

	_, err := Eval("a % b", ctxt)
	if err == nil {
		t.Error("Expected error in string % bool")
	}
}

func TestRemUintUint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(uint32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(uint64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestRemUintUint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUintFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != int64(1) {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUintFloat32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestRemUint64Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != int64(1) {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUint8Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUint8Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint8(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint8) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUint8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUint8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUint8Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestRemUint8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUint16Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUint16Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUint16Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestRemUint16Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemInt8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemInt8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemInt8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemUint64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemBoolBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = true

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestRemBoolBool1(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = false

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error")
	}

}

func TestRemBoolBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = true

	val, err := Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestRemBoolBool3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = false

	_, err := Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error")
	}

}

func TestSubConstInt(t *testing.T) {

	val, err := Eval("1 - 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubConstFloat(t *testing.T) {

	val, err := Eval("1.15 - 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1.10 {
		t.Error("Expected -1.10 get ", val)
	}
}

func TestSubConstFloatInt1(t *testing.T) {

	val, err := Eval("1 - 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1.25 {
		t.Error("Expected -1.25 get ", val)
	}
}

func TestSubConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 - 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1.25 {
		t.Error("Expected 1.25 get ", val)
	}
}

func TestSubConstStringInt(t *testing.T) {

	val, err := Eval("\"a\" - 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "a" {
		t.Error("Expected a get ", val)
	}
}

func TestSubConstIntString(t *testing.T) {

	_, err := Eval("1 - \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubConstIntString2(t *testing.T) {

	val, err := Eval("1 - \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubConstStringString(t *testing.T) {

	val, err := Eval("\"abc\" - \"b\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "ac" {
		t.Error("Expected ac get ", val)
	}
}

func TestSubIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != -1 {
			t.Error("Expected -1 get ", val)
		}
	} else {
		if val.(int64) != -1 {
			t.Error("Expected -1 get ", val)
		}
	}
}

func TestSubIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestSubIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a - b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestSubIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestSubInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a - b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestSubInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestSubFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestSubFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a - b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestSubFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestSubStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "2z2"
	ctxt["b"] = 2

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestSubStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z2"
	ctxt["b"] = int64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestSubStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "2z"
	ctxt["b"] = float64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestSubStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "xzx"
	ctxt["b"] = "z"

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "xx" {
		t.Error("Expected xx get ", val)
	}
}

func TestSubStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "2true2"
	ctxt["b"] = true

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "22" {
		t.Error("Expected 22 get ", val)
	}
}

func TestSubStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "false2false"
	ctxt["b"] = false

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "2" {
		t.Error("Expected 2 get ", val)
	}
}

func TestSubUnConstInt(t *testing.T) {

	val, err := Eval("-2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -2 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUnConstFloat(t *testing.T) {

	val, err := Eval("-2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
		if val.(float64) != -2.25 {
		}
		t.Error("Expected -2.25 get ", val)
	}
}

func TestSubBoolBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = true

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubBoolBool1(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = false

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestSubBoolBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = true

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestSubBoolBool3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = false

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestSubUintUint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(uint32) != 4294967295 {
			t.Error("Expected 4294967295 get ", val)
		}
	} else {
		if val.(uint64) != 18446744073709551615 {
			t.Error("Expected 18446744073709551615 get ", val)
		}
	}

}

func TestSubUintUint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 18446744073709551615 {
		t.Error("Expected 18446744073709551615 get ", val)
	}
}

func TestSubUintFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUintFloat32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(float32) != -1 {
			t.Error("Expected -1 get ", val)
		}
	} else {
		if val.(float64) != -1 {
			t.Error("Expected -1 get ", val)
		}
	}
}

func TestSubUint64Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUint8Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint16) != 65535 {
		t.Error("Expected 65535 get ", val)
	}
}

func TestSubUint8Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint8(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint8) != 255 {
		t.Error("Expected 255 get ", val)
	}
}

func TestSubUint8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUint8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUint8Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != -1 {
			t.Error("Expected -1 get ", val)
		}
	} else {
		if val.(int64) != -1 {
			t.Error("Expected -1 get ", val)
		}
	}
}

func TestSubUint8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUint16Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUint16Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUint16Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != -1 {
			t.Error("Expected -1 get ", val)
		}
	} else {
		if val.(int64) != -1 {
			t.Error("Expected -1 get ", val)
		}
	}
}

func TestSubUint16Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUint64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 18446744073709551615 {
		t.Error("Expected 18446744073709551615 get ", val)
	}
}

func TestAndConstInt(t *testing.T) {

	val, err := Eval("1 & 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) & int64(2)) {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndConstFloat(t *testing.T) {

	val, err := Eval("1.15 & 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) & int64(2)) {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndConstFloatInt1(t *testing.T) {

	val, err := Eval("1 & 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) & int64(2)) {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 & 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(2) & int64(1)) {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndConstStringInt(t *testing.T) {

	_, err := Eval("\"a\" & 1", nil)

	if err == nil {
		t.Error("Expected error on string & int64", err)
	}

}

func TestAndConstIntString(t *testing.T) {

	_, err := Eval("1 & \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndConstIntString2(t *testing.T) {

	val, err := Eval("1 & \"1\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) & int64(1)) {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndConstStringString(t *testing.T) {

	val, err := Eval("\"1\" & \"1\"", nil)

	if err != nil {
		t.Error("Unexpected err and error is not nil")
	}

	if val.(int) != (1 & 1) {
		t.Error("Expected 1 get ", val)
	}
}

func TestAndConstIntBool(t *testing.T) {

	val, err := Eval("1 & true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) & int64(1)) {
		t.Error("Expected (int64(1) & int64(1)) get ", val)
	}
}

func TestAndIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = 2

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != (2 & 2) {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != (2 & 2) {
			t.Error("Expected 3 get ", val)
		}

	}
}

func TestAndIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(1)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestAndIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 33
	ctxt["b"] = float64(33)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 33 {
		t.Error("Expected 33 get ", val)
	}
}

func TestAndIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a & b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = true

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestAndIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(int64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestAndInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(3)
	ctxt["b"] = 3

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(3)
	ctxt["b"] = int64(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = float64(3)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a & b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 2 get ", val)
	}
}

func TestAndFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(55)
	ctxt["b"] = 55

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 55 {
		t.Error("Expected 55 get ", val)
	}
}

func TestAndFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(789)
	ctxt["b"] = int64(789)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 789 {
		t.Error("Expected 789 get ", val)
	}
}

func TestAndFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(966234)
	ctxt["b"] = float64(966234)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 966234 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a & b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = true

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAndFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	_, err := Eval("a & b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	_, err := Eval("a & b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	_, err := Eval("a & b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := Eval("a & b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	_, err := Eval("a & b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = false

	_, err := Eval("a & b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAndBoolBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = true

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndBoolBool1(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = false

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndBoolBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = true

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestAndBoolBool3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = false

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUintUint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(uint32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(uint64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestAndUintUint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUintFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUintFloat32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(int64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestAndUint64Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUint8Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint16) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUint8Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(2)
	ctxt["b"] = uint8(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint8) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestAndUint8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUint8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUint8Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(int64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestAndUint8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUint16Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUint16Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestAndUint16Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 0 {
			t.Error("Expected 0 get ", val)
		}
	} else {
		if val.(int64) != 0 {
			t.Error("Expected 0 get ", val)
		}
	}
}

func TestAndUint16Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a & b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestOrConstInt(t *testing.T) {

	val, err := Eval("1 | 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) | int64(2)) {
		t.Error("Expected 3 get ", val)
	}
}

func TestOrConstFloat(t *testing.T) {

	val, err := Eval("1.15 | 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) | int64(2)) {
		t.Error("Expected 3 get ", val)
	}
}

func TestOrConstFloatInt1(t *testing.T) {

	val, err := Eval("1 | 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) | int64(2)) {
		t.Error("Expected 3 get ", val)
	}
}

func TestOrConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 | 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(2) | int64(1)) {
		t.Error("Expected 3 get ", val)
	}
}

func TestOrConstStringInt(t *testing.T) {

	_, err := Eval("\"a\" | 1", nil)

	if err == nil {
		t.Error("Expected error on string | int64", err)
	}

}

func TestOrConstIntString(t *testing.T) {

	_, err := Eval("1 | \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestOrConstIntString2(t *testing.T) {

	val, err := Eval("1 | \"1\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) | int64(1)) {
		t.Error("Expected 2 get ", val)
	}
}

func TestOrConstStringString(t *testing.T) {

	val, err := Eval("\"1\" | \"1\"", nil)

	if err != nil {
		t.Error("Unexpected err and error is not nil")
	}

	if val.(int) != (1 | 1) {
		t.Error("Expected 1 get ", val)
	}
}

func TestOrConstIntBool(t *testing.T) {

	val, err := Eval("1 | true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != (int64(1) | int64(1)) {
		t.Error("Expected (int64(1) | int64(1)) get ", val)
	}
}

func TestOrIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = 2

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != (2 | 2) {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != (2 | 2) {
			t.Error("Expected 2 get ", val)
		}
	}

}

func TestOrIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(1)

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestOrIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 33
	ctxt["b"] = float64(33)

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 33 {
		t.Error("Expected 33 get ", val)
	}
}

func TestOrIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a | b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestOrIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = true

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestOrIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestOrInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(3)
	ctxt["b"] = 3

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestOrInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(3)
	ctxt["b"] = int64(2)

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestOrInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = float64(3)

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestOrInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a | b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestOrInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestOrInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestOrFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(55)
	ctxt["b"] = 55

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 55 {
		t.Error("Expected 55 get ", val)
	}
}

func TestOrFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(789)
	ctxt["b"] = int64(789)

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 789 {
		t.Error("Expected 789 get ", val)
	}
}

func TestOrFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(966234)
	ctxt["b"] = float64(966234)

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 966234 {
		t.Error("Expected 966234 get ", val)
	}
}

func TestOrFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a | b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestOrFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = true

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestOrFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a | b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 0 get ", val)
	}
}

func TestOrStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	_, err := Eval("a | b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestOrStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	_, err := Eval("a | b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestOrStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	_, err := Eval("a | b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestOrStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := Eval("a | b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestOrStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	_, err := Eval("a | b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestOrStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = false

	_, err := Eval("a | b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlConstInt(t *testing.T) {

	val, err := Eval("1 << 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlConstFloat(t *testing.T) {

	val, err := Eval("1.15 << 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlConstFloatInt1(t *testing.T) {

	val, err := Eval("1 << 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 << 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlConstStringInt(t *testing.T) {

	_, err := Eval("\"a\" << 1", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlConstIntString(t *testing.T) {

	_, err := Eval("1 << \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlConstIntString2(t *testing.T) {

	val, err := Eval("1 << \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlConstStringString(t *testing.T) {

	val, err := Eval("\"1\" << \"2\"", nil)

	if err != nil {
		t.Error("Unexpected err and error is not nil")
	}

	if val.(int) != (1 << 2) {
		t.Error("Expected 1 get ", val)
	}
}

func TestShlConstIntBool(t *testing.T) {

	val, err := Eval("2 << true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 4 {
			t.Error("Expected 4 get ", val)
		}
	} else {
		if val.(int64) != 4 {
			t.Error("Expected 4 get ", val)
		}
	}

}

func TestShlUintUint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(uint32) != 4 {
			t.Error("Expected 4 get ", val)
		}
	} else {
		if val.(uint64) != 4 {
			t.Error("Expected 4 get ", val)
		}
	}
}

func TestShlUintUint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUintFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUintFloat32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 4 {
			t.Error("Expected 4 get ", val)
		}
	} else {
		if val.(int64) != 4 {
			t.Error("Expected 4 get ", val)
		}
	}
}

func TestShlUint64Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUint8Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint16) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUint8Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint8(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint8) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUint8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUint8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUint8Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 4 {
			t.Error("Expected 4 get ", val)
		}
	} else {
		if val.(int64) != 4 {
			t.Error("Expected 4 get ", val)
		}
	}
}

func TestShlUint8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUint16Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUint16Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUint16Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 4 {
			t.Error("Expected 4 get ", val)
		}
	} else {
		if val.(int64) != 4 {
			t.Error("Expected 4 get ", val)
		}
	}
}

func TestShlUint16Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlInt8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlInt8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlInt8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlUint64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a << b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 4 {
			t.Error("Expected 4 get ", val)
		}
	} else {
		if val.(int64) != 4 {
			t.Error("Expected 4 get ", val)
		}
	}
}

func TestShlIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestShlInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a << b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestShlFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a << b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestShlFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestShlStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	_, err := Eval("a << b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	_, err := Eval("a << b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	_, err := Eval("a << b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}

}

func TestShlStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := Eval("a << b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	_, err := Eval("a << b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShlStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = false

	_, err := Eval("a << b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}
func TestShlBoolBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = true

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestShlBoolBool1(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = false

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShlBoolBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = true

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestShlBoolBool3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = false

	val, err := Eval("a << b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestShrConstInt(t *testing.T) {

	val, err := Eval("2 >> 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrConstFloat(t *testing.T) {

	val, err := Eval(" 2.25 >> 1.15", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrConstFloatInt1(t *testing.T) {

	val, err := Eval("2.25 >> 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrConstFloatInt2(t *testing.T) {

	val, err := Eval("2 >> 1.15", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrConstStringInt(t *testing.T) {

	_, err := Eval("\"a\" >> 1", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShrConstIntString(t *testing.T) {

	_, err := Eval("1 >> \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShrConstIntString2(t *testing.T) {

	val, err := Eval("2 >> \"1\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrConstStringString(t *testing.T) {

	val, err := Eval("\"1\" >> \"2\"", nil)

	if err != nil {
		t.Error("Unexpected err and error is not nil")
	}

	if val.(int) != (1 >> 2) {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrConstIntBool(t *testing.T) {

	val, err := Eval("2 >> true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = 1

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}

}

func TestShrUintUint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(2)
	ctxt["b"] = uint(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(uint32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(uint64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestShrUintUint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(2)
	ctxt["b"] = uint64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUintFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(2)
	ctxt["b"] = float64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUintFloat32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(2)
	ctxt["b"] = float32(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestShrUint64Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(2)
	ctxt["b"] = float32(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUint8Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(2)
	ctxt["b"] = uint16(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUint8Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(2)
	ctxt["b"] = uint8(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint8) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUint8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(2)
	ctxt["b"] = int8(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUint8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(2)
	ctxt["b"] = int16(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUint8Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(2)
	ctxt["b"] = int(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestShrUint8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(2)
	ctxt["b"] = int64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUint16Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(2)
	ctxt["b"] = int8(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUint16Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(2)
	ctxt["b"] = int16(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUint16Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(2)
	ctxt["b"] = int(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestShrUint16Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(2)
	ctxt["b"] = int64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrInt8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(2)
	ctxt["b"] = int16(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrInt8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(2)
	ctxt["b"] = int8(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrInt8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(2)
	ctxt["b"] = int64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrUint64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(2)
	ctxt["b"] = uint(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = int64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = float64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a >> b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShrIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 1 {
			t.Error("Expected 1 get ", val)
		}
	} else {
		if val.(int64) != 1 {
			t.Error("Expected 1 get ", val)
		}
	}
}

func TestShrIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestShrInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = 1

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = int64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = float64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a >> b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShrInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestShrFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = 1

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = int64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = float64(1)

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a >> b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShrFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestShrStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	_, err := Eval("a >> b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShrStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	_, err := Eval("a >> b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShrStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	_, err := Eval("a >> b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}

}

func TestShrStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := Eval("a >> b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShrStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	_, err := Eval("a >> b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestShrStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = false

	_, err := Eval("a >> b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}
func TestShrBoolBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = true

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestShrBoolBool1(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = false

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestShrBoolBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = true

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestShrBoolBool3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = false

	val, err := Eval("a >> b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestXorConstInt(t *testing.T) {

	val, err := Eval("1 ^ 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorConstFloat(t *testing.T) {

	val, err := Eval("1.15 ^ 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorConstFloatInt1(t *testing.T) {

	val, err := Eval("1 ^ 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 ^ 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorConstStringInt(t *testing.T) {

	_, err := Eval("\"a\" ^ 1", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestXorConstIntString(t *testing.T) {

	_, err := Eval("1 ^ \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestXorConstIntString2(t *testing.T) {

	val, err := Eval("1 ^ \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorConstStringString(t *testing.T) {

	val, err := Eval("\"1\" ^ \"2\"", nil)

	if err != nil {
		t.Error("Unexpected err and error is not nil")
	}

	if val.(int) != (1 ^ 2) {
		t.Error("Expected 1 get ", val)
	}
}

func TestXorConstIntBool(t *testing.T) {

	val, err := Eval("2 ^ true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}

}

func TestXorUintUint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(uint32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(uint64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestXorUintUint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUintFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUintFloat32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestXorUint64Float32(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = float32(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUint8Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUint8Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = uint8(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint8) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUint8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUint8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUint8Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestXorUint8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUint16Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUint16Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUint16Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestXorUint16Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint16(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorInt8Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int16) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorInt8Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int8) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorInt8Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int8(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorUint64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = uint64(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(uint64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a ^ b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestXorIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 3 {
			t.Error("Expected 3 get ", val)
		}
	} else {
		if val.(int64) != 3 {
			t.Error("Expected 3 get ", val)
		}
	}
}

func TestXorIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if strconv.IntSize == 32 {
		if val.(int32) != 2 {
			t.Error("Expected 2 get ", val)
		}
	} else {
		if val.(int64) != 2 {
			t.Error("Expected 2 get ", val)
		}
	}
}

func TestXorInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a ^ b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestXorInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestXorFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a ^ b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestXorFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestXorFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestXorStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	_, err := Eval("a ^ b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestXorStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	_, err := Eval("a ^ b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestXorStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	_, err := Eval("a ^ b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}

}

func TestXorStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := Eval("a ^ b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestXorStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	_, err := Eval("a ^ b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestXorStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = false

	_, err := Eval("a ^ b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}
func TestXorBoolBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = true

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestXorBoolBool1(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = false

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestXorBoolBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = true
	ctxt["b"] = true

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestXorBoolBool3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = false
	ctxt["b"] = false

	val, err := Eval("a ^ b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}
