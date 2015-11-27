package evalGo

import "testing"

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
