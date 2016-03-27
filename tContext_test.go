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

	_, err = Eval("b", &ctxt)
	if err == nil {
		t.Error("Expected err", err)
	}
}

type helper struct {
	A int
}

type helper2 struct {
}

func (a helper2) GetIdent(name string) (val interface{}, err error) {
	return name, nil
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

	_, err = Eval("b", ctxt2)
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

	_, err = Eval("b", ctxt)
	if err == nil {
		t.Error("Expected err", err)
	}
}

func TestCtxContextBasics(t *testing.T) {
	ctxt := helper2{}
	val, err := Eval("azf", ctxt)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "azf" {
		t.Error("Expected a get ", val)
	}

	ctxt2 := &helper2{}
	val, err = Eval("azf", ctxt2)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "azf" {
		t.Error("Expected a get ", val)
	}
}
