package goScript

import (
	"testing"
)

func TestEvalBasics(t *testing.T) {
	exp := Expr{}
	//err := exp.Prepare("a.Day()")
	err := exp.Prepare("1+1")

	if err != nil {
		t.Error("err not nil", err)
	}

	val, err := exp.Eval(nil)

	if err != nil {
		t.Error("err not nil", err)
	}

	if val.(int64) != 2 {
		t.Error("expected 2 get ", val)
	}
}
