package goScript

import (
	"testing"
)

func TestEqlConstInt(t *testing.T) {

	val, err := Eval("1 == 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlConstInt2(t *testing.T) {

	val, err := Eval("1 == 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlConstFloat(t *testing.T) {

	val, err := Eval("1.15 == 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}
func TestEqlConstFloat2(t *testing.T) {

	val, err := Eval("2.25 == 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlConstFloatInt1(t *testing.T) {

	val, err := Eval("1 == 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlConstFloatInt2(t *testing.T) {

	val, err := Eval("1 == 1.0", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlConstFloatInt3(t *testing.T) {

	val, err := Eval("2.25 == 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlConstFloatInt4(t *testing.T) {

	val, err := Eval("1.0 == 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlConstStringInt(t *testing.T) {

	val, err := Eval("\"a\" == 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlConstStringInt2(t *testing.T) {

	val, err := Eval("\"1\" == 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlConstIntString(t *testing.T) {

	_, err := Eval("1 == \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestEqlConstIntString2(t *testing.T) {

	val, err := Eval("1 == \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlConstStringString(t *testing.T) {

	val, err := Eval("\"a\" == \"b\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlConstStringString2(t *testing.T) {

	val, err := Eval("\"a\" == \"a\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlConstIntBool(t *testing.T) {

	val, err := Eval("2 == true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlConstIntBool2(t *testing.T) {

	val, err := Eval("0 == true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a == b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestEqlIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlInt64Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = uint8(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}
func TestEqlInt64Uint8_2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = uint8(1)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlInt64Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlInt64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlInt64Uint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlInt64Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlInt64Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a == b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestEqlInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64Uint8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = uint8(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64Uint16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = uint16(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64Uint(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = uint(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64Uint64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = uint64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64Int8(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int8(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64Int16(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int16(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a == b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestEqlFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestEqlFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestEqlStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "false"
	ctxt["b"] = false

	val, err := Eval("a == b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstInt(t *testing.T) {

	val, err := Eval("1 >= 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqConstInt2(t *testing.T) {

	val, err := Eval("1 >= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstFloat(t *testing.T) {

	val, err := Eval("1.15 >= 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqConstFloat2(t *testing.T) {

	val, err := Eval("2.25 >= 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstFloatInt1(t *testing.T) {

	val, err := Eval("1 >= 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqConstFloatInt2(t *testing.T) {

	val, err := Eval("1 >= 1.0", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstFloatInt3(t *testing.T) {

	val, err := Eval("2.25 >= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstFloatInt4(t *testing.T) {

	val, err := Eval("1.0 >= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstStringInt(t *testing.T) {

	val, err := Eval("\"a\" >= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstStringInt2(t *testing.T) {

	val, err := Eval("\"1\" >= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstIntString(t *testing.T) {

	_, err := Eval("1 >= \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestGeqConstIntString2(t *testing.T) {

	val, err := Eval("1 >= \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqConstStringString(t *testing.T) {

	val, err := Eval("\"a\" >= \"b\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqConstStringString2(t *testing.T) {

	val, err := Eval("\"a\" >= \"a\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstIntBool(t *testing.T) {

	val, err := Eval("2 >= true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqConstIntBool2(t *testing.T) {

	val, err := Eval("0 >= true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a >= b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestGeqIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a >= b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestGeqInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a >= b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestGeqFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGeqStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGeqStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected trye get ", val)
	}
}

func TestGeqStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "false"
	ctxt["b"] = false

	val, err := Eval("a >= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtConstInt(t *testing.T) {

	val, err := Eval("1 > 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstInt2(t *testing.T) {

	val, err := Eval("1 > 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstFloat(t *testing.T) {

	val, err := Eval("1.15 > 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}
func TestGrtConstFloat2(t *testing.T) {

	val, err := Eval("2.25 > 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstFloatInt1(t *testing.T) {

	val, err := Eval("1 > 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstFloatInt2(t *testing.T) {

	val, err := Eval("1 > 1.0", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstFloatInt3(t *testing.T) {

	val, err := Eval("2.25 > 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtConstFloatInt4(t *testing.T) {

	val, err := Eval("1.0 > 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstStringInt(t *testing.T) {

	val, err := Eval("\"a\" > 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtConstStringInt2(t *testing.T) {

	val, err := Eval("\"1\" > 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstIntString(t *testing.T) {

	_, err := Eval("1 > \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestGrtConstIntString2(t *testing.T) {

	val, err := Eval("1 > \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstStringString(t *testing.T) {

	val, err := Eval("\"a\" > \"b\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstStringString2(t *testing.T) {

	val, err := Eval("\"a\" > \"a\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtConstIntBool(t *testing.T) {

	val, err := Eval("2 > true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtConstIntBool2(t *testing.T) {

	val, err := Eval("0 > true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a > b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestGrtIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a > b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestGrtInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a > b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestGrtFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestGrtStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestGrtStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "false"
	ctxt["b"] = false

	val, err := Eval("a > b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqConstInt(t *testing.T) {

	val, err := Eval("1 <= 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected false get ", val)
	}
}

func TestLeqConstInt2(t *testing.T) {

	val, err := Eval("1 <= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqConstFloat(t *testing.T) {

	val, err := Eval("1.15 <= 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}
func TestLeqConstFloat2(t *testing.T) {

	val, err := Eval("2.25 <= 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqConstFloatInt1(t *testing.T) {

	val, err := Eval("1 <= 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqConstFloatInt2(t *testing.T) {

	val, err := Eval("1 <= 1.0", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqConstFloatInt3(t *testing.T) {

	val, err := Eval("2.25 <= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqConstFloatInt4(t *testing.T) {

	val, err := Eval("1.0 <= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqConstStringInt(t *testing.T) {

	val, err := Eval("\"a\" <= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqConstStringInt2(t *testing.T) {

	val, err := Eval("\"1\" <= 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqConstIntString(t *testing.T) {

	_, err := Eval("1 <= \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestLeqConstIntString2(t *testing.T) {

	val, err := Eval("1 <= \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true  get ", val)
	}
}

func TestLeqConstStringString(t *testing.T) {

	val, err := Eval("\"a\" <= \"b\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqConstStringString2(t *testing.T) {

	val, err := Eval("\"a\" <= \"a\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqConstIntBool(t *testing.T) {

	val, err := Eval("2 <= true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqConstIntBool2(t *testing.T) {

	val, err := Eval("0 <= true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a <= b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestLeqIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a <= b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestLeqInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a <= b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestLeqFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLeqStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLeqStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "false"
	ctxt["b"] = false

	val, err := Eval("a <= b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssConstInt(t *testing.T) {

	val, err := Eval("1 < 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstInt2(t *testing.T) {

	val, err := Eval("1 < 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstFloat(t *testing.T) {

	val, err := Eval("1.15 < 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}
func TestLssConstFloat2(t *testing.T) {

	val, err := Eval("2.25 < 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstFloatInt1(t *testing.T) {

	val, err := Eval("1 < 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssConstFloatInt2(t *testing.T) {

	val, err := Eval("1 < 1.0", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstFloatInt3(t *testing.T) {

	val, err := Eval("2.25 < 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstFloatInt4(t *testing.T) {

	val, err := Eval("1.0 < 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstStringInt(t *testing.T) {

	val, err := Eval("\"a\" < 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstStringInt2(t *testing.T) {

	val, err := Eval("\"1\" < 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstIntString(t *testing.T) {

	_, err := Eval("1 < \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestLssConstIntString2(t *testing.T) {

	val, err := Eval("1 < \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true  get ", val)
	}
}

func TestLssConstStringString(t *testing.T) {

	val, err := Eval("\"a\" < \"b\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssConstStringString2(t *testing.T) {

	val, err := Eval("\"a\" < \"a\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstIntBool(t *testing.T) {

	val, err := Eval("2 < true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssConstIntBool2(t *testing.T) {

	val, err := Eval("0 < true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a < b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestLssIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a < b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestLssInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a < b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestLssFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}

func TestLssStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == false {
		t.Error("Expected true get ", val)
	}
}

func TestLssStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	val, err := Eval("a <b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected true get ", val)
	}
}

func TestLssStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "false"
	ctxt["b"] = false

	val, err := Eval("a < b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) == true {
		t.Error("Expected false get ", val)
	}
}
