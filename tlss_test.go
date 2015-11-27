package evalGo

import "testing"

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
		t.Error("Expected false get ", val)
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
