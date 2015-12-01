package evalGoTest

import (
	"testing"

	"github.com/japm/evalGo"
)

func TestAndBasics(t *testing.T) {
	val, err := evalGo.Eval("true && true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = evalGo.Eval("false && true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("true && false", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("false && false", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}
}

func TestOrBasics(t *testing.T) {
	val, err := evalGo.Eval("true || true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = evalGo.Eval("false || true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = evalGo.Eval("true || false", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = evalGo.Eval("false || false", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}
}

func TestNotBasics(t *testing.T) {
	val, err := evalGo.Eval("!true", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("!false", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = evalGo.Eval("!2", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("!2.0", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("!0", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = evalGo.Eval("!0.0", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected true get ", val)
	}

	val, err = evalGo.Eval("!\"true\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("!\"4\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("!\"4.1\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != false {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("!\"false\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("!\"0\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected false get ", val)
	}

	val, err = evalGo.Eval("!\"0.0\"", nil)
	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(bool) != true {
		t.Error("Expected false get ", val)
	}
}
