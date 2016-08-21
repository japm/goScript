package goScript

import (
	"fmt"
	"testing"
)

var ret interface{}

func BenchmarkIdent1(b *testing.B) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	exp := &Expr{}
	err := exp.Prepare("a")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkIdent2(b *testing.B) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	exp := &Expr{}
	err := exp.Prepare("1")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkUnary1(b *testing.B) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	exp := &Expr{}
	err := exp.Prepare("+1")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkSum1(b *testing.B) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	exp := &Expr{}
	err := exp.Prepare("a+a")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkSum2(b *testing.B) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	exp := &Expr{}
	err := exp.Prepare("a+1")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkSum3(b *testing.B) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	exp := &Expr{}
	err := exp.Prepare("a+(a+a)")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkSum4(b *testing.B) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	exp := &Expr{}
	err := exp.Prepare("a+(1+1)")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkSumConstant(b *testing.B) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	exp := &Expr{}
	err := exp.Prepare("1+0+(1-1)")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkMul1(b *testing.B) {
	ctxt := make(map[string]interface{})
	ctxt["a"] = 1

	exp := &Expr{}
	err := exp.Prepare("2 * a")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkFunc1(b *testing.B) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = helper{1}

	exp := &Expr{}
	err := exp.Prepare("a.Test()")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}

func BenchmarkFunc2(b *testing.B) {

	ctxt := make(map[string]interface{})
	ctxt["OneFn"] = OneFn

	exp := &Expr{}
	err := exp.Prepare("OneFn()")

	if err != nil {
		fmt.Println(err)
		return
	}

	var val interface{}
	for n := 0; n < b.N; n++ {
		val, _ = exp.EvalNoRecover(ctxt)
	}
	ret = val
}
