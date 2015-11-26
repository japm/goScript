package evalGo

import "testing"

func TestCtxMapBasics(t *testing.T) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	val, err := Eval("a", ctxt)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}

	val, err = Eval("b", ctxt)
	if err == nil {
		t.Error("Expected err", err)
	}
}
