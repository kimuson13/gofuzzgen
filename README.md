# gofuzzgen
gofuzzgen is generate template of fuzzing test code
## Caution
gofuzzgen is intended to be used for Go1.18 or higher
# Install
```
$ go install github.com/kimuson13/gofuzzgen/cmd/gofuzzgen@latest
```
# Situation to use
When you want to generate a template code of go standard fuzzing test, `gofuzzgen` help that.
# How to use
If you want to generate fuzzing test code with a package, give the package path of interest as the first
argument:
```
$ gofuzzgen github.com/kimuson13/gofuzzgen
```
To generate fuzzing test code with all packages beneath the current directory:
```
$ gofuzzgen ./...
```
# Demo
If you are on a directoty like that
```
$ tree .
sample
├── cmd
│   └── sample
│       └── main.go
├── go.mod
├── go.sum
└── sample.go
```
```
$ cat sample.go
package sample

func CanFuzzFunc(a int, b int) {
    return a * b
}
```
If you run `gofuzzgen`
```
# Future Outlook
# License
The source code is licensed MIT. The website content is licensed CC BY 4.0,see LICENSE.