package main

import (
	"github.com/YoshikiShibata/predeclarednames"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(predeclarednames.Analyzer) }