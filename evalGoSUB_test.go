package evalGo

import "testing"

func TestSubConstInt(t *testing.T) {

	val, err := Eval("1 - 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubConstFloat(t *testing.T) {

	val, err := Eval("1.15 - 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1.10 {
		t.Error("Expected -1.10 get ", val)
	}
}

func TestSubConstFloatInt1(t *testing.T) {

	val, err := Eval("1 - 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1.25 {
		t.Error("Expected -1.25 get ", val)
	}
}

func TestSubConstFloatInt2(t *testing.T) {

	val, err := Eval("2.25 - 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1.25 {
		t.Error("Expected 1.25 get ", val)
	}
}

func TestSubConstStringInt(t *testing.T) {

	val, err := Eval("\"a\" - 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "a" {
		t.Error("Expected a get ", val)
	}
}

func TestSubConstIntString(t *testing.T) {

	_, err := Eval("1 - \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubConstIntString2(t *testing.T) {

	val, err := Eval("1 - \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubConstStringString(t *testing.T) {

	val, err := Eval("\"abc\" - \"b\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "ac" {
		t.Error("Expected ac get ", val)
	}
}

func TestSubIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubIntInt64(t *testing.T) {

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

func TestSubIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := Eval("a - b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := Eval("a - b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestSubFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := Eval("a - b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "2z2"
	ctxt["b"] = 2

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestSubStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z2"
	ctxt["b"] = int64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestSubStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "2z"
	ctxt["b"] = float64(2)

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestSubStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "xzx"
	ctxt["b"] = "z"

	val, err := Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "xx" {
		t.Error("Expected xx get ", val)
	}
}
