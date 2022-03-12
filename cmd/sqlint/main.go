package main

import (
	"github.com/lufeee/sqlint"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(sqlint.Analyzer) }
