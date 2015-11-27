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

	_, err := Eval("\"2\" / \"2\"", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
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
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
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
	if val.(int) != 2 {
		t.Error("Expected 2 get ", val)
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
