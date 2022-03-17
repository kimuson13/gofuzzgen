package main

import (
	"github.com/gostaticanalysis/codegen/singlegenerator"
	"github.com/kimuson13/gofuzzgen"
)

func main() {
	singlegenerator.Main(gofuzzgen.Generator)
}
