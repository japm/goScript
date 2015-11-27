package evalGo

import "testing"

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
	if val.(int) != 4 {
		t.Error("Expected 4 get ", val)
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
	if val.(int) != 2 {
		t.Error("Expected 2 get ", val)
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
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
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
