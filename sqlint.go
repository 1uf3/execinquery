package sqlint

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "sqlint is golang-sql-linter"

// Analyzer is checking database/sql pkg Query's function
var Analyzer = &analysis.Analyzer{
	Name: "sqlint",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CallExpr:
			selector, ok := n.Fun.(*ast.SelectorExpr)
			if !ok {
				break
			}

			if selector.Sel.Name != "QueryRowContext" {
				break
			}

			for _, arg := range n.Args {
				basicLit, ok := arg.(*ast.BasicLit)
				if !ok {
					continue
				}
				s := strings.Replace(basicLit.Value, "\"", "", -1)
				if strings.HasPrefix(strings.ToLower(s), "select") {
					continue
				}
				s = strings.ToTitle(strings.Split(s, " ")[0])
				pass.Reportf(n.Fun.Pos(), "QueryRowContext() not recommended execute `%s` query", s)
			}
		}
	})
	return nil, nil
}
