package evalGo

import "testing"

func TestAndBasics(t *testing.T) {
	val, err := Eval("true && true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = Eval("false && true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("true && false", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("false && false", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}
}
