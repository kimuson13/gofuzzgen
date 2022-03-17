package gofuzzgen

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/gostaticanalysis/codegen"
	"github.com/kimuson13/showfuzz"
	"golang.org/x/tools/go/analysis"
)

type InputData struct {
	PkgName   string
	Fuzzables []showfuzz.Event
}

var (
	flagOutput string
)

func init() {
	Generator.Flags.StringVar(&flagOutput, "o", "", "output file name")
}

var Generator = &codegen.Generator{
	Name: "gofuzzgen",
	Doc:  "gofuzzgen generate a go fuzzing test template",
	Run:  run,
	Requires: []*analysis.Analyzer{
		showfuzz.Analyzer,
	},
}

func run(pass *codegen.Pass) error {
	if strings.Contains(pass.Pkg.Name(), "_test") {
		return nil
	}

	sfResults := pass.ResultOf[showfuzz.Analyzer].(*showfuzz.Results).Events
	sfrExported := make([]showfuzz.Event, 0, len(sfResults))
	for _, r := range sfResults {
		if unicode.IsUpper(rune(r.Name[0])) {
			sfrExported = append(sfrExported, r)
		}
	}
	data := InputData{pass.Pkg.Name(), sfrExported}

	t, err := template.New("fuzz").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("template.New.Parse: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return fmt.Errorf("template.Excute: %w", err)
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("format.Source: %w", err)
	}

	if flagOutput == "" {
		pass.Print(string(src))
		return nil
	}

	f, err := os.Create(fmt.Sprintf("%s_fuzz_test.go", flagOutput))
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	if _, err := fmt.Fprint(f, string(src)); err != nil {
		return fmt.Errorf("unexpected err: %w", err)
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

var tmpl = `
// This file is generated by gofuzzgen. 
// Only generate fuzzing template.
package {{ .PkgName }}_test

import "testing"

{{ range $i, $v := .Fuzzables }}
func Fuzz{{ $v.Name }}(f *testing.F) {
	{{- if eq (len $v.Args) 1 -}}
		{{- if (index $v.Args 0).IsByteArr }}
			testcases := [][]{{ (index $v.Args 0).UnderlyingName }}{
				// Add seed corpus here
			}

			for _, tc := range testcases {
				f.Add(tc) // Use f.Add to provide a seed corpus
			}

			f.Fuzz(func(t *testing.T, orig []{{ (index $v.Args 0).UnderlyingName }}) {
				// implement fuzzing test code
			})
		{{ else }}
			testcases := []{{ (index $v.Args 0).UnderlyingName }}{
				// Add seed corpus here
			}

			for _, tc := range testcases {
				f.Add(tc)
			}

			f.Fuzz(func(t *testing.T, orig {{ (index $v.Args 0).UnderlyingName }}) {
				// implemnt fuzzing test code
			})
		{{ end }}
	{{ else -}}
		testcases := []struct{
			{{ range $vi, $va := $v.Args -}}
				{{ if $va.IsByteArr -}}
				arg{{ $vi }} []{{ $va.UnderlyingName}}
				{{ else -}}
				arg{{ $vi }} {{ $va.UnderlyingName }}
				{{ end -}}
			{{ end }}
		}{
			// Add seed corpus here
		}

		f.Fuzz(func(t *testing.T, {{ range $vi, $va := $v.Args}}org{{ $vi }} {{ if $va.IsByteArr }}[]{{ $va.UnderlyingName -}}{{ else }}{{ $va.UnderlyingName -}}{{ end }}, {{ end }}) {
				// implemnt fuzzing test code
			})
	{{ end }}
}
{{ end }}
`
