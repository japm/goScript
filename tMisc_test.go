package goScript

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

func (a *helper) TestPtr1() int {
	return 1
}

func (a helper) Test6(x float64, y float64) float64 {
	return x + y
}

func TestCallBasics(t *testing.T) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = helper{1}
	ctxt["a1"] = new(helper)

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

	val, err = Eval("a.TestPtr1()", ctxt)
	if err == nil {
		t.Error("Expected err and error is nil")
	}

	val, err = Eval("a1.TestPtr1()", ctxt)
	if err != nil {
		t.Errorf("err not nil", err)
	}
	if val.(int) != 1 {
		t.Errorf("Expected 1 returned %d", val)
	}

	val, err = Eval("a.Test6(2, 3)", ctxt)
	if err != nil {
		t.Errorf("err not nil", err)
	}
	if val.(float64) != 5 {
		t.Errorf("Expected 5 returned %d", val)
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

func TestSlice1(t *testing.T) {
	ctxt := make(map[string]interface{})
	a := make([]int, 1)
	a[0] = 3
	ctxt["a"] = a

	val, err := Eval("a[0]", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestSlice2(t *testing.T) {
	ctxt := make(map[string]interface{})
	a := make([]int, 3)
	a[0] = 3
	ctxt["a"] = a

	val, err := Eval("a[0:2]", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if len(val.([]int)) != 2 {
		t.Error("Expected slice of 2 get ", val)
	}
}
func TestArray1(t *testing.T) {
	ctxt := make(map[string]interface{})
	a := []int{3, 1, 2, 3}
	ctxt["a"] = a

	val, err := Eval("a[0]", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestArray2(t *testing.T) {
	ctxt := make(map[string]interface{})
	a := []int{0, 1, 2, 3}
	ctxt["a"] = a

	val, err := Eval("a[0:2]", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if len(val.([]int)) != 2 {
		t.Error("Expected slice of 2 get ", val)
	}
}

func TestMap1(t *testing.T) {
	ctxt := make(map[string]interface{})
	a := make(map[string]int)
	a["a"] = 3
	ctxt["a"] = a

	val, err := Eval("a[\"a\"]", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestStar1(t *testing.T) {
	ctxt := make(map[string]interface{})
	a := new(int)
	*a = 3
	ctxt["a"] = a

	val, err := Eval("*a", ctxt)
	if err != nil {
		t.Error("err not nil", err)
	}

	if val.(int) != 3 {
		t.Error("Expected 3 get ", val)
	}

}
