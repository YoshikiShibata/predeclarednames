// Copyright © 2019 Yoshiki Shibata. All rights reserved.

package predeclarednames

import (
	"go/ast"
	"go/doc"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "predeclarednames",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "predeclarednames checks that names which match with any predeclared names"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
		(*ast.StructType)(nil),
		(*ast.FuncType)(nil),
		(*ast.InterfaceType)(nil),
		(*ast.LabeledStmt)(nil),
		(*ast.AssignStmt)(nil),
		(*ast.GenDecl)(nil),
		(*ast.FuncDecl)(nil),
	}

	reportf := func(pos token.Pos, kind, name string) {
		pass.Reportf(pos, "%s: shadowing a predeclared identifier (%s)", kind, name)
	}

	analyzeIdent := func(ident *ast.Ident, kind string) {
		if doc.IsPredeclared(ident.Name) {
			reportf(ident.NamePos, kind, ident.Name)
		}
	}

	analyzeFieldList := func(fl *ast.FieldList, kind string) {
		for _, field := range fl.List {
			for _, ident := range field.Names {
				analyzeIdent(ident, kind)
			}
		}
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.StructType:
			if n.Fields != nil {
				analyzeFieldList(n.Fields, "field")
			}
		case *ast.TypeSpec:
			analyzeIdent(n.Name, "type")
		case *ast.FuncType:
			analyzeFieldList(n.Params, "param")
			if n.Results != nil {
				analyzeFieldList(n.Results, "result")
			}
		case *ast.InterfaceType:
			if n.Methods != nil {
				analyzeFieldList(n.Methods, "method")
			}
		case *ast.LabeledStmt:
			analyzeIdent(n.Label, "label")
		case *ast.AssignStmt:
			if n.Tok != token.DEFINE {
				return
			}

			// short variable declaration
			for _, expr := range n.Lhs {
				if ident, ok := expr.(*ast.Ident); ok {
					analyzeIdent(ident, "variable")
				}
			}
		case *ast.GenDecl:
			var kind string
			switch n.Tok {
			case token.CONST:
				kind = "const"
			case token.VAR:
				kind = "variable"
			default:
				return
			}

			for _, spec := range n.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, ident := range valueSpec.Names {
						analyzeIdent(ident, kind)
					}
				}
			}
		case *ast.FuncDecl:
			if n.Recv == nil {
				analyzeIdent(n.Name, "function")
				return
			}

			analyzeIdent(n.Name, "method")
			analyzeFieldList(n.Recv, "receiver")
		}

	})

	return nil, nil
}
