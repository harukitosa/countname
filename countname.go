package countname

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "countname is ..."

// Analyzer is ...
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
			// log.Printf("%+v\n", decl)
			for i := 0; i < len(decl.Specs); i++ {
				// fmt.Printf("specs:%+v\n", decl.Specs[i])
				switch spec := decl.Specs[i].(type) {
				case *ast.ValueSpec:
					// fmt.Printf("value:%+v\n", spec)
					for _, name := range spec.Names {
						// fmt.Printf("name:%v\n", name)
						if len(name.Name) >= 20 {
							pass.Reportf(n.Pos(), "name is too long")
						}
					}
				}
			}
		case *ast.FuncDecl:
			// log.Printf("%+v len:%d\n", n, len(decl.Name.Name))
			if len(decl.Name.Name) >= 20 {
				pass.Reportf(n.Pos(), "name is too long")
			}
		}
	})

	return nil, nil
}
