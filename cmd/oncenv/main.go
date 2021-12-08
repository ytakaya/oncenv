package main

import (
	"github.com/ytakaya/oncenv"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(oncenv.Analyzer) }
