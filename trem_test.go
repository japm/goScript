package evalGo

import "testing"

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

	_, err := Eval("\"1\" % \"2\"", nil)

	if err == nil {
		t.Error("Expected error on string % string")
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
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
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
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
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
