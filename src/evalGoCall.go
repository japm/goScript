package evalGo

import (
	"fmt"
	"go/ast"
	"reflect"
)

type callSyte struct {
	callee    interface{}
	fnName    string
	calleeVal reflect.Value
}

func evalSelectorExpr(expr *ast.SelectorExpr, context interface{}) (interface{}, error) {
	callee, err := eval(expr.X, context)
	if err != nil {
		return nil, err
	}

	calleeVal := reflect.ValueOf(callee)
	fieldVal := calleeVal
	if fieldVal.Kind() == reflect.Ptr {
		fieldVal = calleeVal.Elem() //FieldByName panics on pointers
	}
	_, ok := fieldVal.Type().FieldByName(expr.Sel.Name)
	if ok {
		return fieldVal.FieldByName(expr.Sel.Name).Interface(), nil
	}
	return callSyte{callee, expr.Sel.Name, calleeVal}, nil
}

func evalCallExpr(expr *ast.CallExpr, context interface{}) (interface{}, error) {

	val, err := eval(expr.Fun, context) //Find the type called, this calls evalSelectorExpr
	if err != nil {
		return nil, err
	}

	switch val.(type) { //The return value must be a callSyte
	case callSyte:
		break
	default:
		return nil, fmt.Errorf("Waiting callsite found %s", reflect.TypeOf(val))
	}

	callsite := val.(callSyte)
	caleeVal := callsite.calleeVal

	//-------------------Check Method------------------------
	method := caleeVal.MethodByName(callsite.fnName)
	if !method.IsValid() {
		return nil, fmt.Errorf("Method %s not found", callsite.fnName)
	}

	mType := method.Type()
	numArgs := mType.NumIn()

	if !mType.IsVariadic() {
		if len(expr.Args) != numArgs {
			return nil, fmt.Errorf("Method alguments count mismatch. Expected %d get %d", numArgs, len(expr.Args))
		}
	} else {
		numArgs = numArgs - 1
		if len(expr.Args) < numArgs {
			return nil, fmt.Errorf("Method alguments count mismatch. Expected at least %d get %d", (numArgs - 1), len(expr.Args))
		}
	}

	//-------------------Prepare call arguments ------------
	var args []reflect.Value
	if len(expr.Args) == 0 {
		args = zeroArg //Zero arg constant
	} else {
		args = make([]reflect.Value, len(expr.Args))
	}

	var vaArgsTp reflect.Type
	if mType.IsVariadic() {
		vaArgsTp = mType.In(numArgs).Elem() //Type declared
	}

	for key, value := range expr.Args {

		val, err := eval(value, context)
		if err != nil {
			return nil, err
		}
		rVal := reflect.ValueOf(val)
		var parArg reflect.Type

		//If true we are in the variadic parameters
		if key >= numArgs {
			parArg = vaArgsTp
		} else {
			parArg = mType.In(key)
		}
		if rVal.Type() != parArg {
			if !rVal.Type().ConvertibleTo(parArg) {
				return nil, fmt.Errorf("Method argument %s type mismatch. Expected %s get %s", key, parArg, rVal.Type())
			}
			rVal = rVal.Convert(parArg)
		}
		args[key] = rVal
	}

	//Call
	retVal := method.Call(args)

	//Evaluate result
	if len(retVal) == 0 {
		return nil, nil
	}
	return retVal[0].Interface(), nil
}
