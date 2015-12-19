package goScript

import (
	"testing"
)

func TestCtxIMapBasics(t *testing.T) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = helper{1}
	val, err := Eval("a", &ctxt)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(helper).A != 1 {
		t.Error("Expected 1 get ", val)
	}

	val, err = Eval("b", &ctxt)
	if err == nil {
		t.Error("Expected err", err)
	}
}

type helper struct {
	A int
}

func TestCtxInterfBasics(t *testing.T) {
	ctxt := helper{1}
	val, err := Eval("A", ctxt)
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

	ctxt2 := &helper{1}
	val, err = Eval("A", ctxt2)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}

	val, err = Eval("b", ctxt2)
	if err == nil {
		t.Error("Expected err", err)
	}
}

func TestCtxMapBasics(t *testing.T) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = helper{1}
	val, err := Eval("a", ctxt)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(helper).A != 1 {
		t.Error("Expected 1 get ", val)
	}

	val, err = Eval("b", ctxt)
	if err == nil {
		t.Error("Expected err", err)
	}
}
