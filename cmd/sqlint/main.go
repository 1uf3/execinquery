package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"lufe.jp/sqlint"
)

func main() { unitchecker.Main(sqlint.Analyzer) }
