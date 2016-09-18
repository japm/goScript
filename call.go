//Package goScript, javascript Eval() for go
//The MIT License (MIT)
//Copyright (c) 2016 Juan Pascual

package goScript

import (
	"fmt"
	"go/ast"
	"reflect"
)

type callSite struct {
	callee  interface{}
	fnName  string
	isValid bool
}

func evalSelectorExpr(expr *ast.SelectorExpr, context Context) (interface{}, error) {
	callee, err := eval(expr.X, context)
	if err != nil {
		return nilInterf, err
	}
	sName := expr.Sel.Name

	//If callee is a context, then resolve with getIdent
	calleeContext, ok := callee.(Context)
	if ok {
		return calleeContext.GetIdent(sName)
	}
	calleeContextPtr, ok := callee.(*Context)
	if ok {
		return (*calleeContextPtr).GetIdent(sName)
	}

	calleeVal := reflect.ValueOf(callee)
	fieldVal := calleeVal
	if fieldVal.Kind() == reflect.Ptr {
		fieldVal = calleeVal.Elem() //FieldByName panics on pointers
	}

	fbnVal, ok := fieldVal.Type().FieldByName(sName) //Faster than  fieldVal.FieldByName(sName)
	if ok {
		return fieldVal.FieldByIndex(fbnVal.Index).Interface(), nil
	}

	mval := fieldVal.MethodByName(sName)
	if !mval.IsValid() {
		return nilInterf, fmt.Errorf("Selector %s not field nor method", sName)
	}

	//Return function pointer
	return mval.Interface(), nil
}

func evalSelectorExprCall(expr *ast.SelectorExpr, context Context) (interface{}, callSite, error) {
	callee, err := eval(expr.X, context)
	if err != nil {
		return nilInterf, callSite{isValid: false}, err
	}
	return nilInterf, callSite{callee, expr.Sel.Name, true}, nil
}

func evalCallExpr(expr *ast.CallExpr, context Context) (interface{}, error) {

	//Find the type called, this calls evalSelectorExpr/ evalIdent
	val, callsite, err := evalFromCall(expr.Fun, context)
	if err != nil {
		return nilInterf, err
	}

	//-------------------Check Method------------------------
	var vaArgsTp reflect.Type
	var method reflect.Value
	var methodCallable Callable
	var isCallable bool

	if !callsite.isValid {
		//Not a callSite, must be a Ident(args), so Ident must be a function
		if val == nil {
			return nilInterf, fmt.Errorf("Waiting reflect.Func/Callable found %v", "nil")
		}
		methodCallable, isCallable = val.(Callable)
		if !isCallable {
			method = reflect.ValueOf(val)
			if method.Kind() != reflect.Func {
				return nilInterf, fmt.Errorf("Waiting reflect.Func found %v", reflect.TypeOf(val))
			}
		}
	} else {
		callableContext, ok := callsite.callee.(CallableContext)
		if ok {
			methodCallable, err = callableContext.GetCallable(callsite.fnName)
			isCallable = (err == nil)
		} else {
			isCallable = false
		}
		if !isCallable {
			//A callSite so must be f.x(args)
			caleeVal := reflect.ValueOf(callsite.callee)
			method = caleeVal.MethodByName(callsite.fnName)
			if !method.IsValid() {
				return nilInterf, fmt.Errorf("Method %s not found", callsite.fnName)
			}
		}
	}
	if isCallable {
		var args []interface{}
		if len(expr.Args) == 0 {
			args = zeroArgInterf //Zero arg constant
		} else {
			args = make([]interface{}, len(expr.Args))

			for key, value := range expr.Args {

				val, err := eval(value, context)
				if err != nil {
					return nilInterf, err
				}
				args[key] = val
			}
		}
		return methodCallable.Call(args)
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
	}
	//Call
	retVal := method.Call(args)

	//Evaluate result
	if len(retVal) == 0 {
		return nilInterf, nil
	}
	return retVal[0].Interface(), nil

}
