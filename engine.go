/*
The MIT License (MIT)
Copyright (c) 2016 Juan Pascual
*/
package goScript

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sync"
)

type Engine struct {
	files map[string]*ast.File
	lock  sync.Mutex
}

func (e Engine) AddSrc(name string, src string) error {
	if isEmpty(name) {
		return fmt.Errorf("Empty name")
	}

	if isEmpty(src) {
		return fmt.Errorf("Empty source")
	}

	fset := token.FileSet{}

	fl, err := parser.ParseFile(&fset, "", src, parser.AllErrors)
	if err != nil {
		return err
	}

	e.lock.Lock()
	defer func() {
		e.lock.Unlock()
	}()

	e.files[name] = fl

	return nil

}
