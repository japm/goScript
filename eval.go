package evalGo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
)

// Expr expresion holder, allows sentence preparation
type Expr struct {
	expr ast.Expr
}

var zeroArg []reflect.Value

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

// Eval a pepared sentence
func (e *Expr) Eval(context interface{}) (val interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Eval paniked. %s", r)
		}
	}()

	val, err = eval(e.expr, context)
	return val, err
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

	var ctxt interface{}
	switch context.(type) {
	case *map[string]interface{}:
		ctxt = context
	case map[string]interface{}:
		ctxt = context
	case reflect.Value:
		ctxt = context
	default:
		ctxt = reflect.ValueOf(context)
	}

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

func evalInt(expr ast.Node, context interface{}) (int, error) {
	val, err := eval(expr, context)
	if err != nil {
		return 0, err
	}
	return castInt(val)
}

func eval(expr ast.Node, context interface{}) (interface{}, error) {
	//fmt.Println(reflect.TypeOf(expr), time.Now().UnixNano()/int64(10000), expr)
	switch expr.(type) {
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
	case *ast.BasicLit:
		return evalBasicLit(expr.(*ast.BasicLit), context)
	case *ast.BinaryExpr:
		return evalBinaryExpr(expr.(*ast.BinaryExpr), context)
	case *ast.BlockStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.BranchStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.CallExpr:
		return evalCallExpr(expr.(*ast.CallExpr), context)
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
	case *ast.Ident:
		return evalIdent(expr.(*ast.Ident), context)
	case *ast.IfStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.ImportSpec:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.IncDecStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.IndexExpr:
		return evalIndexExpr(expr.(*ast.IndexExpr), context)
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
	case *ast.ParenExpr:
		return evalParenExpr(expr.(*ast.ParenExpr), context)
	case *ast.RangeStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.ReturnStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.SelectStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.SelectorExpr:
		return evalSelectorExpr(expr.(*ast.SelectorExpr), context)
	case *ast.SendStmt:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	case *ast.SliceExpr:
		return evalSliceExpr(expr.(*ast.SliceExpr), context)
	case *ast.StarExpr:
		return evalStarExpr(expr.(*ast.StarExpr), context)
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
	case *ast.UnaryExpr:
		return evalUnaryExpr(expr.(*ast.UnaryExpr), context)
	case *ast.ValueSpec:
		return nil, fmt.Errorf("%s not suported", reflect.TypeOf(expr))
	default:
		return nil, fmt.Errorf("Default %s not suported", reflect.TypeOf(expr))
	}
}

func evalIndexExpr(expr *ast.IndexExpr, context interface{}) (interface{}, error) {

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
	if v.Kind() == reflect.Map {
		retVal = v.MapIndex(reflect.ValueOf(idx))
	} else if v.Kind() == reflect.Array ||
		v.Kind() == reflect.Slice {
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

func evalStarExpr(expr *ast.StarExpr, context interface{}) (interface{}, error) {

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

func evalSliceExpr(expr *ast.SliceExpr, context interface{}) (interface{}, error) {

	//Get array object
	val, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}
	sl := reflect.ValueOf(val)

	//Check the type
	if sl.Kind() != reflect.Array &&
		sl.Kind() != reflect.Slice {
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

func evalParenExpr(expr *ast.ParenExpr, context interface{}) (interface{}, error) {
	return eval(expr.X, context)
}

func evalUnaryExpr(expr *ast.UnaryExpr, context interface{}) (interface{}, error) {
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

func evalBinaryExpr(expr *ast.BinaryExpr, context interface{}) (interface{}, error) {

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

func evalIdent(expr *ast.Ident, context interface{}) (interface{}, error) {

	lname := len(expr.Name)

	if lname == 3 && expr.Name == "nil" {
		return nil, nil
	} else if lname == 4 && expr.Name == "true" {
		return true, nil
	} else if lname == 5 && expr.Name == "false" {
		return false, nil
	}
	if context == nil {
		return nil, fmt.Errorf("Context is null, no ident possible for %s", expr.Name)
	}
	var ok bool
	var val interface{}

	switch context.(type) {
	case map[string]interface{}:
		val, ok = context.(map[string]interface{})[expr.Name]
	case *map[string]interface{}:
		val, ok = (*context.(*map[string]interface{}))[expr.Name]
	case reflect.Value:
		refVal := context.(reflect.Value)
		if refVal.Kind() == reflect.Ptr {
			refVal = refVal.Elem()
		}
		fld := refVal.FieldByName(expr.Name)
		ok = fld.IsValid()
		if ok {
			val = fld.Interface()
		}
	default:
		refVal := reflect.ValueOf(context)
		if refVal.Kind() == reflect.Ptr {
			refVal = refVal.Elem()
		}
		fld := refVal.FieldByName(expr.Name)
		ok = fld.IsValid()
		if ok {
			val = fld.Interface()
		}
	}
	if !ok {
		return nil, fmt.Errorf("Symbol %s not found", expr.Name)
	}
	return val, nil
}

func evalBasicLit(expr *ast.BasicLit, context interface{}) (interface{}, error) {
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
	switch expr.Op {
	case token.ADD:
		return evalBinaryExprADD(left, right)
	case token.SUB:
		return evalBinaryExprSUB(left, right)
	case token.MUL:
		return evalBinaryExprMUL(left, right)
	case token.QUO:
		return evalBinaryExprQUO(left, right)
	case token.REM:
		return evalBinaryExprREM(left, right)
	case token.EQL: // ==
		return evalBinaryExprEQL(left, right)
	case token.LSS: // <
		return evalBinaryExprLSS(left, right)
	case token.GTR: // >
		return evalBinaryExprGTR(left, right)
	case token.NEQ: // !=
		return evalBinaryExprNEQ(left, right)
	case token.GEQ: // >=
		return evalBinaryExprGEQ(left, right)
	case token.LEQ: // <=
		return evalBinaryExprLEQ(left, right)
	case token.AND:
		return evalBinaryExprAND(left, right)
	case token.OR:
		return evalBinaryExprOR(left, right)
	case token.SHL:
		return evalBinaryExprSHL(left, right)
	case token.SHR:
		return evalBinaryExprSHR(left, right)
	case token.XOR:
		return evalBinaryExprXOR(left, right)
	case token.AND_NOT:
		return evalBinaryExprANDNOT(left, right)
	default:
		return nil, fmt.Errorf("evalBinaryExprOp not implemented for %d", expr.Op)
	}

}

func evalBinaryExprOpLazy(expr *ast.BinaryExpr, context interface{}) (interface{}, error) {
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
	case uint:
		return -value.(uint), nil
	case uint64:
		return -value.(uint64), nil
	case int8:
		return -value.(int8), nil
	case int16:
		return -value.(int16), nil
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
