package execinquery

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "execinquery is a linter about query string checker in Query function which reads your Go src files and warning it finds"

// Analyzer is checking database/sql pkg Query's function
var Analyzer = &analysis.Analyzer{
	Name: "execinquery",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	result := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	result.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CallExpr:
			if len(n.Args) < 1 {
				return
			}

			selector, ok := n.Fun.(*ast.SelectorExpr)
			if !ok {
				return
			}

			if "database/sql" != pass.TypesInfo.Uses[selector.Sel].Pkg().Path() {
				return
			}

			if !strings.Contains(selector.Sel.Name, "Query") {
				return
			}

			var i int
			if strings.Contains(selector.Sel.Name, "Context") {
				i = 1
			}

			var s string
			switch arg := n.Args[i].(type) {
			case *ast.BasicLit:
				s = strings.Replace(arg.Value, "\"", "", -1)

			case *ast.Ident:

				switch arg2 := arg.Obj.Decl.(type) {
				case *ast.AssignStmt:
					for _, stmt := range arg2.Rhs {
						basicLit, ok := stmt.(*ast.BasicLit)
						if !ok {
							continue
						}

						s = strings.Replace(basicLit.Value, "\"", "", -1)
					}
				case *ast.ValueSpec:
					basicLit, ok := arg2.Values[0].(*ast.BasicLit)
					if !ok {
						return
					}

					s = strings.TrimLeftFunc(basicLit.Value, func(r rune) bool {
						return !unicode.IsLetter(r) && !unicode.IsNumber(r)
					})
					s = strings.Replace(s, "\"", "", -1)
				}

			default:
				return
			}

			if strings.HasPrefix(strings.ToLower(s), "select") {
				return
			}

			s = strings.ToTitle(strings.SplitN(s, " ", 2)[0])

			pass.Reportf(n.Fun.Pos(), "It's better to use Execute method instead of %s method to execute `%s` query", selector.Sel.Name, s)
		}
	})

	return nil, nil
}
