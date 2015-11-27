package evalGo

import "testing"

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
