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

func TestIdent(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"

	val, err := Eval("a", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestIdent2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"

	_, err := Eval("b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestIdent3(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	val, err := Eval("a", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestIdent4(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1.4

	val, err := Eval("a", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1.4 {
		t.Error("Expected 1.4 get ", val)
	}
}
