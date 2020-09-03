package main

import (
	"github.com/harukitosa/merucari/countname"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(countname.Analyzer) }

