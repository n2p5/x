package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

// from the example at https://pkg.go.dev/go/ast

func main() {
	exp1()
	exp2()
}

func exp2() {
	// Parse the file
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "data/foo.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Modify the AST
	// For example, let's rename all functions that start with "old" to start with "new"
	ast.Inspect(f, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			name := funcDecl.Name.Name
			if len(name) >= 3 && name[:3] == "old" {
				funcDecl.Name.Name = "new" + name[3:]
			}
		}
		return true
	})

	// Create output file
	outputFile, err := os.Create("data/output/modified_foo.go")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Write modified AST to file
	err = printer.Fprint(outputFile, fset, f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully modified and saved the file!")
}

func exp1() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "data/foo.go", nil, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		var nodeType string

		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
			nodeType = "BasicLit"
		case *ast.Ident:
			s = x.Name
			nodeType = "Ident"
		case *ast.FuncDecl:
			s = x.Name.Name
			nodeType = "FuncDecl"
		case *ast.AssignStmt:
			s = fmt.Sprintf("%d assignments", len(x.Lhs))
			nodeType = "AssignStmt"
		case *ast.CallExpr:
			nodeType = "CallExpr"
			switch fun := x.Fun.(type) {
			case *ast.Ident:
				s = fun.Name
			case *ast.SelectorExpr:
				if sel, ok := fun.X.(*ast.Ident); ok {
					s = sel.Name + "." + fun.Sel.Name
				}
			}
		case *ast.BinaryExpr:
			s = x.Op.String()
			nodeType = "BinaryExpr"
		case *ast.SelectorExpr:
			nodeType = "SelectorExpr"
			if id, ok := x.X.(*ast.Ident); ok {
				s = id.Name + "." + x.Sel.Name
			}
		}
		if s != "" {
			fmt.Printf("%s:\t[%s]\t%s\n", fset.Position(n.Pos()), nodeType, s)
		}
		return true
	})
}
