package evalGo

import "testing"

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

func TestAddIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 3 {
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
