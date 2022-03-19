package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/lufeee/execinquery"
)

func main() { unitchecker.Main(execinquery.Analyzer) }
