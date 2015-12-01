package evalGoTest

import (
	"testing"

	"github.com/japm/evalGo"
)

func TestAddConstInt(t *testing.T) {

	val, err := evalGo.Eval("1 + 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddConstFloat(t *testing.T) {

	val, err := evalGo.Eval("1.15 + 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3.4 {
		t.Error("Expected 3.4 get ", val)
	}
}

func TestAddConstFloatInt1(t *testing.T) {

	val, err := evalGo.Eval("1 + 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3.25 {
		t.Error("Expected 3.25 get ", val)
	}
}

func TestAddConstFloatInt2(t *testing.T) {

	val, err := evalGo.Eval("2.25 + 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3.25 {
		t.Error("Expected 3.25 get ", val)
	}
}

func TestAddConstStringInt(t *testing.T) {

	val, err := evalGo.Eval("\"a\" + 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "a1" {
		t.Error("Expected a1 get ", val)
	}
}

func TestAddConstIntString(t *testing.T) {

	_, err := evalGo.Eval("1 + \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAddConstIntString2(t *testing.T) {

	val, err := evalGo.Eval("1 + \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddConstStringString(t *testing.T) {

	val, err := evalGo.Eval("\"1\" + \"2\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "12" {
		t.Error("Expected 12 get ", val)
	}
}

func TestAddConstIntBool(t *testing.T) {

	val, err := evalGo.Eval("2 + true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a + b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAddIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestAddInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a + b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAddInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestAddFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a + b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestAddFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestAddFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestAddStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z2" {
		t.Error("Expected z2 get ", val)
	}
}

func TestAddStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z2" {
		t.Error("Expected z2 get ", val)
	}
}

func TestAddStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z2" {
		t.Error("Expected z2 get ", val)
	}
}

func TestAddStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "xz" {
		t.Error("Expected xz get ", val)
	}
}

func TestAddStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "ztrue" {
		t.Error("Expected ztrue get ", val)
	}
}

func TestAddStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = false

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "zfalse" {
		t.Error("Expected ztrue get ", val)
	}
}

func TestDivConstInt(t *testing.T) {

	val, err := evalGo.Eval("2 / 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivConstFloat(t *testing.T) {

	val, err := evalGo.Eval("1.15 / 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 0.5111111111111111 {
		t.Error("Expected 0.5111111111111111  get ", val)
	}
}

func TestDivConstFloatInt1(t *testing.T) {

	val, err := evalGo.Eval("2 / 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 0.8888888888888888 {
		t.Error("Expected 0.8888888888888888  get ", val)
	}
}

func TestDivConstFloatInt2(t *testing.T) {

	val, err := evalGo.Eval("2.25 / 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1.125 {
		t.Error("Expected 1.125 get ", val)
	}
}

func TestDivConstStringInt(t *testing.T) {

	_, err := evalGo.Eval("\"a\" / 3", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivConstIntString(t *testing.T) {

	_, err := evalGo.Eval("1 / \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivConstIntString2(t *testing.T) {

	val, err := evalGo.Eval("2 / \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivConstStringString(t *testing.T) {

	_, err := evalGo.Eval("\"2\" / \"2\"", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = 2

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestDivIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(1)

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestDivInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestDivFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a / b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestDivFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = true

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestDivStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = false

	_, err := evalGo.Eval("a / b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulConstInt(t *testing.T) {

	val, err := evalGo.Eval("2 * 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulConstFloat(t *testing.T) {

	val, err := evalGo.Eval("1.15 * 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2.5875 {
		t.Error("Expected 2.5875 get ", val)
	}
}

func TestMulConstFloatInt1(t *testing.T) {

	val, err := evalGo.Eval("2 * 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4.50 {
		t.Error("Expected 4.50 get ", val)
	}
}

func TestMulConstFloatInt2(t *testing.T) {

	val, err := evalGo.Eval("2.25 * 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4.50 {
		t.Error("Expected 4.50 get ", val)
	}
}

func TestMulConstStringInt(t *testing.T) {

	val, err := evalGo.Eval("\"a\" * 3", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "aaa" {
		t.Error("Expected aaa get ", val)
	}
}

func TestMulConstIntString(t *testing.T) {

	_, err := evalGo.Eval("1 * \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulConstIntString2(t *testing.T) {

	val, err := evalGo.Eval("2 * \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulConstStringString(t *testing.T) {

	val, err := evalGo.Eval("\"2\" * \"2\"", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "22" {
		t.Error("Expected 22 get ", val)
	}
}

func TestMulIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = 2

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestMulInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 4 {
		t.Error("Expected 3 get ", val)
	}
}

func TestMulInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestMulFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 4 {
		t.Error("Expected 4 get ", val)
	}
}

func TestMulFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestMulFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestMulFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestMulStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "zz" {
		t.Error("Expected zz get ", val)
	}
}

func TestMulStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "zz" {
		t.Error("Expected zz get ", val)
	}
}

func TestMulStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a * b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "zz" {
		t.Error("Expected zz get ", val)
	}
}

func TestMulStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a * b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestRemConstInt(t *testing.T) {

	val, err := evalGo.Eval("1 % 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestRemConstFloat(t *testing.T) {

	val, err := evalGo.Eval("1.15 % 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemConstFloatInt1(t *testing.T) {

	val, err := evalGo.Eval("1 % 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemConstFloatInt2(t *testing.T) {

	val, err := evalGo.Eval("2.25 % 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestRemConstStringInt(t *testing.T) {

	_, err := evalGo.Eval("\"a\" % 1", nil)

	if err == nil {
		t.Error("Expected error on string % int")
	}
}

func TestRemConstIntString(t *testing.T) {

	_, err := evalGo.Eval("1 % \"a\" ", nil)

	if err == nil {
		t.Error("Expected error on int % string")
	}
}

func TestRemConstIntString2(t *testing.T) {

	val, err := evalGo.Eval("1 % \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestRemConstStringString(t *testing.T) {

	_, err := evalGo.Eval("\"1\" % \"2\"", nil)

	if err == nil {
		t.Error("Expected error on string % string")
	}
}

func TestRemConstIntBool(t *testing.T) {

	val, err := evalGo.Eval("2 % true", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestRemIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestRemIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 0 {
		t.Error("Expected 0 get ", val)
	}
}

func TestRemIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	_, err := evalGo.Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in int % false")
	}
}

func TestRemInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestRemInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestRemInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	_, err := evalGo.Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in int % false")
	}

}

func TestRemFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestRemFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestRemFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a % b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 0 {
		t.Error("Expected 1 get ", val)
	}
}

func TestRemFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	_, err := evalGo.Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in int % false")
	}
}

func TestRemStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = 2

	_, err := evalGo.Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in string % int")
	}
}

func TestRemStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = int64(2)

	_, err := evalGo.Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in string % int64")
	}
}

func TestRemStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = float64(2)

	_, err := evalGo.Eval("a % b", ctxt)

	if err == nil {
		t.Error("Expected error in string % flaot64")
	}
}

func TestRemStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "x"
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a % b", ctxt)
	if err == nil {
		t.Error("Expected error in string % int")
	}
}

func TestRemStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = true

	_, err := evalGo.Eval("a % b", ctxt)
	if err == nil {
		t.Error("Expected error in string % bool")
	}
}

func TestRemStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z"
	ctxt["b"] = false

	_, err := evalGo.Eval("a % b", ctxt)
	if err == nil {
		t.Error("Expected error in string % bool")
	}
}

func TestSubConstInt(t *testing.T) {

	val, err := evalGo.Eval("1 - 2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubConstFloat(t *testing.T) {

	val, err := evalGo.Eval("1.15 - 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1.10 {
		t.Error("Expected -1.10 get ", val)
	}
}

func TestSubConstFloatInt1(t *testing.T) {

	val, err := evalGo.Eval("1 - 2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1.25 {
		t.Error("Expected -1.25 get ", val)
	}
}

func TestSubConstFloatInt2(t *testing.T) {

	val, err := evalGo.Eval("2.25 - 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1.25 {
		t.Error("Expected 1.25 get ", val)
	}
}

func TestSubConstStringInt(t *testing.T) {

	val, err := evalGo.Eval("\"a\" - 1", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "a" {
		t.Error("Expected a get ", val)
	}
}

func TestSubConstIntString(t *testing.T) {

	_, err := evalGo.Eval("1 - \"a\" ", nil)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubConstIntString2(t *testing.T) {

	val, err := evalGo.Eval("1 - \"2\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubConstStringString(t *testing.T) {

	val, err := evalGo.Eval("\"abc\" - \"b\" ", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "ac" {
		t.Error("Expected ac get ", val)
	}
}

func TestSubIntInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = 2

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubIntInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a + b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 3 {
		t.Error("Expected 3 get ", val)
	}
}

func TestSubIntFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubIntString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 1
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a - b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubIntBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = true

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestSubIntBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = 2
	ctxt["b"] = false

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestSubInt64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubInt64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(1)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a - b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubInt64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestSubInt64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = int64(2)
	ctxt["b"] = false

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestSubFloat64Int(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = 2

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubFloat64Int64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected 3 get ", val)
	}
}

func TestSubFloat64Float64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != -1 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubFloat64String(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(1)
	ctxt["b"] = "z"

	_, err := evalGo.Eval("a - b", ctxt)

	if err == nil {
		t.Error("Expected err and error is nil")
	}
}

func TestSubFloat64Bool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = true

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 1 {
		t.Error("Expected 1 get ", val)
	}
}

func TestSubFloat64Bool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = float64(2)
	ctxt["b"] = false

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(float64) != 2 {
		t.Error("Expected 2 get ", val)
	}
}

func TestSubStringInt(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "2z2"
	ctxt["b"] = 2

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestSubStringInt64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "z2"
	ctxt["b"] = int64(2)

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestSubStringFloat64(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "2z"
	ctxt["b"] = float64(2)

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "z" {
		t.Error("Expected z get ", val)
	}
}

func TestSubStingString(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "xzx"
	ctxt["b"] = "z"

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "xx" {
		t.Error("Expected xx get ", val)
	}
}

func TestSubStringBool(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "2true2"
	ctxt["b"] = true

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "22" {
		t.Error("Expected 22 get ", val)
	}
}

func TestSubStringBool2(t *testing.T) {

	ctxt := make(map[string]interface{})
	ctxt["a"] = "false2false"
	ctxt["b"] = false

	val, err := evalGo.Eval("a - b", ctxt)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(string) != "2" {
		t.Error("Expected 2 get ", val)
	}
}

func TestSubUnConstInt(t *testing.T) {

	val, err := evalGo.Eval("-2", nil)

	if err != nil {
		t.Error("err not nil", err)
	}
	if val.(int64) != -2 {
		t.Error("Expected -1 get ", val)
	}
}

func TestSubUnConstFloat(t *testing.T) {

	val, err := evalGo.Eval("-2.25", nil)

	if err != nil {
		t.Error("err not nil", err)
		if val.(float64) != -2.25 {
		}
		t.Error("Expected -2.25 get ", val)
	}
}
