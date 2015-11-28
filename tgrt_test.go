package evalGo

import "testing"

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
