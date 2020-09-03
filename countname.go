package countname

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "countname is ..."

const maxLongNum = 19

// Analyzer search long name decl
var Analyzer = &analysis.Analyzer{
	Name: "countname",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch decl := n.(type) {
		case *ast.GenDecl:
			for i := 0; i < len(decl.Specs); i++ {
				switch spec := decl.Specs[i].(type) {
				case *ast.ValueSpec:
					var flag bool
					for _, name := range spec.Names {
						if !flag && len(name.Name) > maxLongNum {
							pass.Reportf(n.Pos(), "name is longer than %d", maxLongNum)
							flag = true
						}
					}
				}
			}
		case *ast.FuncDecl:
			if len(decl.Name.Name) > maxLongNum {
				pass.Reportf(n.Pos(), "name is longer than %d", maxLongNum)
			}
		}
	})

	return nil, nil
}
