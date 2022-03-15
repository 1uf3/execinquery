package main

import (
	"github.com/lufeee/execinquery"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(execinquery.Analyzer) }
