/*
Package goScript  MIT License (MIT)
Copyright (c) 2016 Juan Pascual
*/
package goScript

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
)

//Helper len function
func reflLen(val interface{}) int {
	valRefl := reflect.ValueOf(val)
	return valRefl.Len()
}

//Context allows custom identification resolver
type Context interface {
	GetIdent(name string) (val interface{}, err error)
}

//Internal map context
type mapContext struct {
	mp map[string]interface{}
}

func (ctx mapContext) GetIdent(name string) (val interface{}, err error) {
	val, ok := ctx.mp[name]
	if !ok {
		return nil, fmt.Errorf("Symbol %s not found", name)
	}
	return val, nil
}

//Internal reflect context
//We can store a map to cache propertys
//and aliviate the use of reflect
type reflectContext struct {
	val reflect.Value
}

func (ctx reflectContext) GetIdent(name string) (val interface{}, err error) {
	//refVal := ctx.val
	//Forced on construction
	//if refVal.Kind() == reflect.Ptr {
	//	refVal = refVal.Elem()
	//}
	fld := ctx.val.FieldByName(name)
	ok := fld.IsValid()
	if !ok {
		return nil, fmt.Errorf("Symbol %s not found", name)
	}

	return fld.Interface(), nil
}

// Expr expresion holder, allows sentence preparation
type Expr struct {
	expr ast.Expr
}

//Constant for zero arg calls
var zeroArg []reflect.Value

//Package initialization
func init() {
	zeroArg = make([]reflect.Value, 0)
}

// Prepare sentence for evaluation and reuse
func (e *Expr) Prepare(expr string) error {
	exp, err := parser.ParseExpr(expr)
	if err != nil {
		return err
	}
	e.expr = exp
	return err
}

// Eval a prepared sentence
func (e *Expr) Eval(context interface{}) (val interface{}, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Eval paniked. %s", r)
		}
	}()

	ctxt := createContext(context)
	val, err = eval(e.expr, ctxt)

	return val, err
}

//EvalNoRecover without recover on defer
func (e *Expr) EvalNoRecover(context interface{}) (interface{}, error) {

	ctxt := createContext(context)
	return eval(e.expr, ctxt)
}

// EvalInt convenient function casting return type to int
func (e *Expr) EvalInt(context interface{}) (val int, err error) {

	valI, err := e.Eval(context)
	if err != nil {
		return 0, err
	}
	return castInt(valI)
}

//Eval expression within a context
func Eval(expr string, context interface{}) (val interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Eval paniked. %s", r)
		}
	}()

	exp, err := parser.ParseExpr(expr)

	if err != nil {
		return nil, err
	}

	ctxt := createContext(context)

	val, err = eval(exp, ctxt)
	return val, err
}

// EvalInt convenient function casting return type to int
func EvalInt(expr string, context interface{}) (int, error) {
	val, err := Eval(expr, context)
	if err != nil {
		return 0, err
	}
	return castInt(val)
}

func evalInt(expr ast.Node, context Context) (int, error) {
	val, err := eval(expr, context)
	if err != nil {
		return 0, err
	}
	return castInt(val)
}

//Creates the apropiate execution identificaion resolver
func createContext(context interface{}) Context {
	var ctxt Context
	switch context.(type) {
	case *map[string]interface{}:
		ctxt = mapContext{*(context.(*map[string]interface{}))}
	case map[string]interface{}:
		ctxt = mapContext{(context.(map[string]interface{}))}
	case reflect.Value:
		rval := context.(reflect.Value)
		if rval.Kind() == reflect.Ptr {
			rval = rval.Elem()
		}
		ctxt = reflectContext{rval}
	case *reflect.Value:
		rval := *(context.(*reflect.Value))
		if rval.Kind() == reflect.Ptr {
			rval = rval.Elem()
		}
		ctxt = reflectContext{rval}
	default:
		var ok bool
		ctxt, ok = context.(Context)
		if !ok {
			rval := reflect.ValueOf(context)
			if rval.Kind() == reflect.Ptr {
				rval = rval.Elem()
			}
			ctxt = reflectContext{rval}
		}
	}

	return ctxt
}

//Evaluates all kinds of ast node types
func eval(expr ast.Node, context Context) (interface{}, error) {
	//fmt.Println(reflect.TypeOf(expr), time.Now().UnixNano()/int64(10000), expr)
	switch expr.(type) {
	/*
		case *ast.ArrayType:
			return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		case *ast.AssignStmt:
			return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		case *ast.BadDecl:
			return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		case *ast.BadExpr:
			return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		case *ast.BadStmt:
			return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	*/
	case *ast.BasicLit:
		return evalBasicLit(expr.(*ast.BasicLit), context)
	case *ast.BinaryExpr:
		return evalBinaryExpr(expr.(*ast.BinaryExpr), context)
		/*
			case *ast.BlockStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.BranchStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		*/
	case *ast.CallExpr:
		return evalCallExpr(expr.(*ast.CallExpr), context)
		/*
			case *ast.CaseClause:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.ChanType:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.CommClause:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.Comment:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.CommentGroup:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.CompositeLit:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.DeclStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.DeferStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.Ellipsis:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.EmptyStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.Field:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.FieldList:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.ForStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.FuncDecl:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.FuncLit:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.FuncType:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.GenDecl:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.GoStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		*/
	case *ast.Ident:
		return evalIdent(expr.(*ast.Ident), context)
		/*
			case *ast.IfStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.ImportSpec:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.IncDecStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		*/
	case *ast.IndexExpr:
		return evalIndexExpr(expr.(*ast.IndexExpr), context)
		/*
			case *ast.InterfaceType:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.KeyValueExpr:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.LabeledStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.MapType:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.Package:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		*/
	case *ast.ParenExpr:
		return evalParenExpr(expr.(*ast.ParenExpr), context)
		/*
			case *ast.RangeStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.ReturnStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.SelectStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		*/
	case *ast.SelectorExpr:
		return evalSelectorExpr(expr.(*ast.SelectorExpr), context)
		/*
			case *ast.SendStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		*/
	case *ast.SliceExpr:
		return evalSliceExpr(expr.(*ast.SliceExpr), context)
	case *ast.StarExpr:
		return evalStarExpr(expr.(*ast.StarExpr), context)
		/*
			case *ast.StructType:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.SwitchStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.TypeAssertExpr:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.TypeSpec:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
			case *ast.TypeSwitchStmt:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		*/
	case *ast.UnaryExpr:
		return evalUnaryExpr(expr.(*ast.UnaryExpr), context)
		/*
			case *ast.ValueSpec:
				return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
		*/
	default:
		return nil, fmt.Errorf("Default %s not suported", reflect.TypeOf(expr))
	}
}

func evalIndexExpr(expr *ast.IndexExpr, context Context) (interface{}, error) {

	val, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}

	idx, err := eval(expr.Index, context)
	if err != nil {
		return nil, err
	}
	var retVal reflect.Value
	v := reflect.ValueOf(val)
	vk := v.Kind()
	if vk == reflect.Map {
		retVal = v.MapIndex(reflect.ValueOf(idx))
	} else if vk == reflect.Array ||
		vk == reflect.Slice {
		i, err := castInt(idx)
		if err != nil {
			return nil, err
		}
		retVal = v.Index(i)

	} else {
		return nil, fmt.Errorf("Unexpected indexer [] type %s ", v)
	}
	if !retVal.IsValid() {
		return nil, nil
	}

	return retVal.Interface(), nil

}

func evalStarExpr(expr *ast.StarExpr, context Context) (interface{}, error) {

	val, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}

	refVal := reflect.ValueOf(val)
	if refVal.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("Expected pointer, found %d ", refVal.Type())
	}

	return refVal.Elem().Interface(), nil
}

func evalSliceExpr(expr *ast.SliceExpr, context Context) (interface{}, error) {

	//Get array object
	val, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}
	sl := reflect.ValueOf(val)

	//Check the type
	vk := sl.Kind()
	if vk != reflect.Array &&
		vk != reflect.Slice {
		return nil, fmt.Errorf("Expected array, found %d ", sl.Type())
	}

	//calculate i,j,j from a[i:j:k]
	var i, j, k int
	if expr.Low == nil {
		i = 0
	} else {
		i, err = evalInt(expr.Low, context)
		if err != nil {
			return nil, err
		}
	}

	if expr.High == nil {
		j = sl.Len()
	} else {
		j, err = evalInt(expr.High, context)
		if err != nil {
			return nil, err
		}
	}

	//Calculate subslice
	if expr.Slice3 {
		k, err = evalInt(expr.Max, context)
		if err != nil {
			return nil, err
		}
		return sl.Slice3(i, j, k).Interface(), nil
	}
	return sl.Slice(i, j).Interface(), nil
}

func evalParenExpr(expr *ast.ParenExpr, context Context) (interface{}, error) {
	return eval(expr.X, context)
}

func evalUnaryExpr(expr *ast.UnaryExpr, context Context) (interface{}, error) {
	val, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}
	switch expr.Op {
	case token.NOT:
		return evalUnaryExprNOT(val)
	case token.SUB:
		return evalUnaryExprSUB(val)
	case token.AND:
		return evalUnaryExprAND(val)
	default:
		return nil, fmt.Errorf("Unary operation %d not implemented", expr.Op)
	}
}

func evalBinaryExpr(expr *ast.BinaryExpr, context Context) (interface{}, error) {

	if expr.Op == token.LAND ||
		expr.Op == token.LOR {
		return evalBinaryExprOpLazy(expr, context)
	}

	left, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}

	right, err := eval(expr.Y, context)
	if err != nil {
		return nil, err
	}

	return evalBinaryExprOp(expr, left, right)
}

func evalIdent(expr *ast.Ident, context Context) (interface{}, error) {

	lname := len(expr.Name)

	//Resolve reserved value-words
	if lname == 3 {
		if expr.Name == "nil" {
			return nil, nil
		} else if expr.Name == "len" {
			return reflLen, nil
		}
	} else if lname == 4 && expr.Name == "true" {
		return true, nil
	} else if lname == 5 && expr.Name == "false" {
		return false, nil
	}
	//No context, cant return an ident
	if context == nil {
		return nil, fmt.Errorf("Context is null, no ident possible for %s", expr.Name)
	}

	return context.GetIdent(expr.Name)
}

func evalBasicLit(expr *ast.BasicLit, context Context) (interface{}, error) {
	switch expr.Kind {
	case token.INT:
		return strconv.ParseInt(expr.Value, 10, strconv.IntSize)
	case token.FLOAT:
		return strconv.ParseFloat(expr.Value, 10)
	case token.IMAG:
		return nil, fmt.Errorf("token.IMAG not suported")
	case token.CHAR:
		return expr.Value, nil
	case token.STRING:
		return expr.Value[1 : len(expr.Value)-1], nil
	default:
		return nil, fmt.Errorf("token type not suported %d %s", expr.Kind, expr.Value)
	}
}

func evalBinaryExprOp(expr *ast.BinaryExpr, left interface{}, right interface{}) (interface{}, error) {
	var op operation
	//op = nil

	switch expr.Op {
	case token.ADD: // +
		op = opAdd{}
	case token.SUB: // -
		op = opSub{}
	case token.MUL: // *
		op = opMul{}
	case token.QUO: // /
		op = opQuo{}
	case token.REM: // %
		op = opRem{}
	case token.EQL: // ==
		return evalBinaryExprEQL(left, right)
	case token.LSS: // <
		op = opLss{}
	case token.GTR: // >
		op = opGtr{}
	case token.NEQ: // !=
		return evalBinaryExprNEQ(left, right)
	case token.GEQ: // >=
		op = opGeq{}
	case token.LEQ: // <=
		op = opLeq{}
	case token.AND: // &
		op = opAnd{}
	case token.OR: // |
		op = opOr{}
	case token.SHL: // <<
		op = opShl{}
	case token.SHR: // >>
		op = opShr{}
	case token.XOR: // ^
		op = opXor{}
	case token.AND_NOT: // &^
		op = opAndNot{}
	default:
		return nil, fmt.Errorf("evalBinaryExprOp not implemented for %d", expr.Op)
	}

	tp, e := binaryOperType(left, right)
	if e != nil {
		return nil, e
	}

	return evalBinary(left, right, tp, op)
}

func evalBinaryExprOpLazy(expr *ast.BinaryExpr, context Context) (interface{}, error) {
	switch expr.Op {
	case token.LAND: // &&
		return evalBinaryExprLAND(expr, context)
	case token.LOR: // ||
		return evalBinaryExprLOR(expr, context)
	default:
		return nil, fmt.Errorf("evalBinaryExprOp not implemented for %d", expr.Op)
	}
}

func evalUnaryExprSUB(value interface{}) (interface{}, error) {
	switch value.(type) {
	case uint8:
		return -value.(uint8), nil
	case uint16:
		return -value.(uint16), nil
	case uint32:
		return -value.(uint32), nil
	case uint:
		return -value.(uint), nil
	case uint64:
		return -value.(uint64), nil
	case int8:
		return -value.(int8), nil
	case int16:
		return -value.(int16), nil
	case int32:
		return -value.(int32), nil
	case int:
		return -value.(int), nil
	case int64:
		return -value.(int64), nil
	case float32:
		return -value.(float32), nil
	case float64:
		return -value.(float64), nil
	case bool:
		return !value.(bool), nil
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented not for type %s", reflect.TypeOf(value))
}

func evalUnaryExprAND(value interface{}) (interface{}, error) {

	val := reflect.ValueOf(value)

	if !val.CanAddr() {
		return nil, fmt.Errorf("Value is not addressable %s", val)
	}

	vk := val.Kind()
	if vk == reflect.Chan || vk == reflect.Func || vk == reflect.Map ||
		vk == reflect.Ptr || vk == reflect.Slice {
		if val.IsNil() {
			return nil, fmt.Errorf("Value is nill, not addressable %s", val)
		}
	}

	return val.Addr(), nil
}
