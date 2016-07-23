/*
The MIT License (MIT)
Copyright (c) 2016 Juan Pascual
*/
package goScript

import (
	"fmt"
	"go/ast"
	"reflect"
)

type callSyte struct {
	callee interface{}
	fnName    string
	calleeVal reflect.Value
}

func evalSelectorExpr(expr *ast.SelectorExpr, context Context) (interface{}, error) {
	callee, err := eval(expr.X, context)
	if err != nil {
		return nilInterf, err
	}

	calleeVal := reflect.ValueOf(callee)
	fieldVal := calleeVal
	if fieldVal.Kind() == reflect.Ptr {
		fieldVal = calleeVal.Elem() //FieldByName panics on pointers
	}
	sName := expr.Sel.Name
	_, ok := fieldVal.Type().FieldByName(sName) //Faster than  fieldVal.FieldByName(sName)
	if ok {
		return fieldVal.FieldByName(sName).Interface(), nil
	}
	//Not a field, must be a function
	return callSyte{callee, sName, calleeVal}, nil
}

func evalCallExpr(expr *ast.CallExpr, context Context) (interface{}, error) {

	//Can be optimized if expr is allways evalSelectorExpr or evalIdent
	//The return value must be a callsite, 10% performance increase on evalSelectorExpr calls
	//with this optimization
	val, err := eval(expr.Fun, context) //Find the type called, this calls evalSelectorExpr
	if err != nil {
		return nilInterf, err
	}

	//-------------------Check Method------------------------
	var vaArgsTp reflect.Type
	var method reflect.Value

	callsite, ok := val.(callSyte)

	if !ok {
		//Not a callSyte, must be a Ident(args), so Ident must be a function
		method = reflect.ValueOf(val)
		if method.Kind() != reflect.Func {
			return nilInterf, fmt.Errorf("Waiting callsite found %s", reflect.TypeOf(val))
		}
	} else {
		//A callSyte so must be f.x(args)
		caleeVal := callsite.calleeVal
		method = caleeVal.MethodByName(callsite.fnName)
		if !method.IsValid() {
			return nilInterf, fmt.Errorf("Method %s not found", callsite.fnName)
		}
	}
	mType := method.Type()
	numArgs := mType.NumIn()

	if !mType.IsVariadic() {
		if len(expr.Args) != numArgs {
			return nilInterf, fmt.Errorf("Method alguments count mismatch. Expected %d get %d", numArgs, len(expr.Args))
		}
	} else {
		numArgs = numArgs - 1
		vaArgsTp = mType.In(numArgs).Elem() //Type declared
		if len(expr.Args) < numArgs {
			return nilInterf, fmt.Errorf("Method alguments count mismatch. Expected at least %d get %d", (numArgs - 1), len(expr.Args))
		}
	}

	//-------------------Prepare call arguments ------------
	var args []reflect.Value
	if len(expr.Args) == 0 {
		args = zeroArg //Zero arg constant
	} else {
		args = make([]reflect.Value, len(expr.Args))
	}

	for key, value := range expr.Args {

		val, err := eval(value, context)
		if err != nil {
			return nilInterf, err
		}
		rVal := reflect.ValueOf(val)
		var tArg reflect.Type //Method argument type

		//If true we are in the variadic parameters
		if key >= numArgs {
			tArg = vaArgsTp
		} else {
			tArg = mType.In(key)
		}
		tVal := rVal.Type() //Passed parameter type
		if tVal != tArg {
			if !tVal.ConvertibleTo(tArg) {
				return nilInterf, fmt.Errorf("Method argument %d type mismatch. Expected %s get %s", key, tArg, tVal)
			}
			rVal = rVal.Convert(tArg)
		}
		args[key] = rVal
	}

	//Call
	retVal := method.Call(args)

	//Evaluate result
	if len(retVal) == 0 {
		return nilInterf, nil
	}
	return retVal[0].Interface(), nil
}
