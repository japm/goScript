package evalGo

import "testing"

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

	val, err := Eval("\"2\" * \"2\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "22" {
		t.Error("Expected 22 get ", val)
	}
}

func TestDivIntInt(t *testing.T) {

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

func TestDivIntInt64(t *testing.T) {

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

func TestDivIntFloat64(t *testing.T) {

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

func TestDivIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = "z"

	_, err := Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivInt64Int(t *testing.T) {

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

func TestDivInt64Int64(t *testing.T) {

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

func TestDivInt64Float64(t *testing.T) {

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

func TestDivInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivFloat64Int(t *testing.T) {

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

func TestDivFloat64Int64(t *testing.T) {

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

func TestDivFloat64Float64(t *testing.T) {

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

func TestDivFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = "z"

	_, err := Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringInt(t *testing.T) {

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

func TestDivStringInt64(t *testing.T) {

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

func TestDivStringFloat64(t *testing.T) {

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

func TestDivStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}
