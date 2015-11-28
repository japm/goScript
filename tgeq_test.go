package evalGo

import "testing"

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
