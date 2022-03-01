package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func main() {
	// ファイルごとのトークンの位置を記録するFileSetを作成する
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "_db.go", nil, 0)
	// f, err := parser.ParseFile(fset, "_gopher.go", nil, 0)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// 	ast.Print(nil, f)

	//抽象構文木を深さ優先で探索する
	ast.Inspect(f, func(n ast.Node) bool {
		// 識別子ではない場合は無視
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selector, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		// 識別子がQueryRowContextという名前でなければ無視
		if selector.Sel.Name != "QueryRowContext" {
			return true
		}

		for _, arg := range callExpr.Args {
			basicLit, ok := arg.(*ast.BasicLit)
			if !ok {
				continue
			}
			// QueryContext関数の中にselectという文字列が入っている場合は無視
			if strings.HasPrefix(strings.ToLower(basicLit.Value), "select") {
				return true
			}
			break
		}

		fmt.Println(fset.Position(callExpr.Fun.Pos()))
		return true
	})

	fmt.Println("finish")
}
