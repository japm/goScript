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
