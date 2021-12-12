package oncenv

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const doc = "oncenv detects os.Getenv in inappropriate place"

var Analyzer = &analysis.Analyzer{
	Name: "oncenv",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		if strings.Contains(f.Name.Name, "_test") {
			continue
		}
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok && decl.Name.Name != "init" {
				ast.Inspect(decl, func(n ast.Node) bool {
					if fun, ok := n.(*ast.SelectorExpr); ok {
						if x, ok := fun.X.(*ast.Ident); ok {
							if x.Name == "os" && fun.Sel.Name == "Getenv" {
								pass.Report(analysis.Diagnostic{
									Pos:     fun.Pos(),
									Message: "os.Getenv() is used in inappropriate place",
								})
							}
						}
					}
					return true
				})
			}
		}
	}
	return nil, nil
}
