package sqlint

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "sqlint is ..."

// Analyzer is ...
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

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {

		switch n := n.(type) {
		case *ast.CallExpr:
			selector, ok := n.Fun.(*ast.SelectorExpr)
			if !ok {
				break
			}

			// 識別子がQueryRowContextという名前でなければ無視
			if selector.Sel.Name != "QueryRowContext" {
				break
			}

			for _, arg := range n.Args {
				basicLit, ok := arg.(*ast.BasicLit)
				if !ok {
					continue
				}
				// QueryContext関数の中にselectという文字列が入っている場合は無視
				s := strings.Replace(basicLit.Value, "\"", "", -1)
				if strings.HasPrefix(strings.ToLower(s), "select") {
					break
				}

				s = strings.ToTitle(strings.Split(s, " ")[0])
				pass.Reportf(n.Fun.Pos(), "QueryRowContext() can not use \"%s\" query", s)
			}
		}
	})

	return nil, nil
}
