package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// from the example at https://pkg.go.dev/go/ast

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "data/foo.go", nil, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			s = x.Name
		}
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})
}
