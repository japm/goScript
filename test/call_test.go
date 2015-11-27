package evalGo

import "testing"

func (a helper) Test() int {
	return 1
}

func (a helper) Test2(z int) int {
	return z
}

func (a helper) Test3(z ...interface{}) []interface{} {
	return z
}

func (a helper) Test4(z ...int) []int {
	return z
}

func (a helper) Test5(x int, z ...interface{}) []interface{} {
	return z
}

func TestCallBasics(t *testing.T) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = helper{1}
	val, err := Eval("a.Test()", ctxt)
	if err != nil {
		t.Errorf("err not nil", err)
	}
	if val.(int) != 1 {
		t.Errorf("Expected 1 returned %d", val)
	}

	val, err = Eval("a.Test2(2)", ctxt)
	if err != nil {
		t.Errorf("err not nil", err)
	}
	if val.(int) != 2 {
		t.Errorf("Expected 2 returned %d", val)
	}

	val, err = Eval("a.Test3(2)", ctxt)
	if err != nil {
		t.Errorf("err not nil", err)
	}
	if len(val.([]interface{})) != 1 {
		t.Errorf("Expected arr returned %d", val)
	}

	val, err = Eval("a.Test4(2)", ctxt)
	if err != nil {
		t.Error("err not nil", err)
	}
	if len(val.([]int)) != 1 {
		t.Errorf("Expected arr returned %d", val)
	}

	val, err = Eval("a.Test5(2,4)", ctxt)
	if err != nil {
		t.Errorf("err not nil", err)
	}
	if len(val.([]interface{})) != 1 {
		t.Errorf("Expected arr returned %d", val)
	}
}
