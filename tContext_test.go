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

type customCtxt struct {
}

func (a customCtxt) GetIdent(name string) (val interface{}, err error) {
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

func TestCtxIdent(t *testing.T) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = customCtxt{}
	ctxt["c"] = &customCtxt{}

	val, _ := Eval("a.b", ctxt)
	if val.(string) != "b" {
		t.Error("Expected b get ", val)
	}

	val, _ = Eval("c.b", ctxt)
	if val.(string) != "b" {
		t.Error("Expected b get ", val)
	}

}

func TestCtxContextBasics(t *testing.T) {
	ctxt := customCtxt{}
	val, err := Eval("azf", ctxt)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "azf" {
		t.Error("Expected a get ", val)
	}

	ctxt2 := &customCtxt{}
	val, err = Eval("azf", ctxt2)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "azf" {
		t.Error("Expected a get ", val)
	}
}
