package goScript

import "testing"

func TestLAndBasics(t *testing.T) {
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

func TestLOrBasics(t *testing.T) {
	val, err := Eval("true || true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = Eval("false || true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = Eval("true || false", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = Eval("false || false", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}
}

func TestLNotBasics(t *testing.T) {
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

	val, err = Eval("!2", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("!2.0", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("!0", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = Eval("!0.0", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = Eval("!\"true\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("!\"4\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("!\"4.1\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("!\"false\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("!\"0\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected false get ", val)
	}

	val, err = Eval("!\"0.0\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected false get ", val)
	}
}
