/*
The MIT License (MIT)
Copyright (c) 2016 Juan Pascual
*/
package goScript

import (
	"go/ast"
	"fmt"
	//"reflect"
	//"time"

)

func isEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, c := range s {
		if c != ' ' && c != '\t' {
			return false
		}
	}
	return true
}

///ConstNode new ast.Node for constant pre evaluation
type ConstNodeBasicLit struct {
	ast.BasicLit
	value interface{}
	node  *ast.Node
}

type ConstNodeIdent struct {
	ast.Ident
	value interface{}
	node  *ast.Node
}

type ConstNodeUnaryExpr struct {
	ast.UnaryExpr
	value interface{}
	node  *ast.Node
}

type ConstNodeBinaryExpr struct {
	ast.BinaryExpr
	value interface{}
	node  *ast.Node
}


//Evaluates all kinds of ast node types
func replaceConstants(expr ast.Node) ast.Node {
	switch expr.(type) {
	case *ast.BasicLit:
		v, e := evalBasicLit(expr.(*ast.BasicLit), nil)
		if e != nil {
			return expr
		}
		return &ConstNodeBasicLit{*expr.(*ast.BasicLit), v, &expr}
	case *ast.BinaryExpr:
		bexp := expr.(*ast.BinaryExpr)
		bexp.X = replaceConstants(bexp.X).(ast.Expr)
		bexp.Y = replaceConstants(bexp.Y).(ast.Expr)
		v, e := evalBinaryExpr(expr.(*ast.BinaryExpr), nilContext{})
		if e != nil {
			return expr
		}
		return &ConstNodeBinaryExpr{*expr.(*ast.BinaryExpr), v, &expr}
	case *ast.ParenExpr:
		pexp := expr.(*ast.ParenExpr)
		pexp.X = replaceConstants(pexp.X).(ast.Expr)
		return expr
	case *ast.Ident:
		v, e := evalIdent(expr.(*ast.Ident), nilContext{})
		if e != nil {
			return expr
		}
		return &ConstNodeIdent{*expr.(*ast.Ident), v, &expr}
	case *ast.UnaryExpr:
		v, e := evalUnaryExpr(expr.(*ast.UnaryExpr), nilContext{})
		if e != nil {
			fmt.Println(e)
			return expr
		}
		return &ConstNodeUnaryExpr{*expr.(*ast.UnaryExpr), v, &expr}
	default:
		return expr
	}
}
