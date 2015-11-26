package evalGo

import "testing"

func TestNotBasics(t *testing.T) {
	val, err := Eval("!true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("!false", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}
}
