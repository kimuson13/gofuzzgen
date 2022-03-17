package gofuzzgen

import (
	"github.com/gostaticanalysis/codegen"
	"github.com/kimuson13/showfuzz"
	"golang.org/x/tools/go/analysis"
)

var Generator = &codegen.Generator{
	Name:     "gofuzzgen",
	Doc:      "gofuzzgen generate a go fuzzing test template",
	Run:      run,
	Requires: []*analysis.Analyzer{showfuzz.Analyzer},
}

func run(pass *codegen.Pass) error {
	return nil
}
