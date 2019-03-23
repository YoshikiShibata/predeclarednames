# predeclarednames
a simple static checker for using predeclared names in Go source files

## **How To Install:**
`go get github.com/YoshikiShibata/predeclaredname`

## **Example**
```
$ predeclarednames text/...
/Users/yoshiki/tools/go/src/text/scanner/scanner.go:328:19: method: shadowing a predeclared identifier (error)
/Users/yoshiki/tools/go/src/text/tabwriter/tabwriter.go:413:18: method: shadowing a predeclared identifier (append)
/Users/yoshiki/tools/go/src/text/template/parse/node.go:88:20: method: shadowing a predeclared identifier (append)
/Users/yoshiki/tools/go/src/text/template/parse/node.go:158:20: method: shadowing a predeclared identifier (append)
/Users/yoshiki/tools/go/src/text/template/parse/node.go:247:23: method: shadowing a predeclared identifier (append)
/Users/yoshiki/tools/go/src/text/template/parse/node.go:498:33: param: shadowing a predeclared identifier (true)
/Users/yoshiki/tools/go/src/text/template/parse/node.go:539:3: variable: shadowing a predeclared identifier (rune)
/Users/yoshiki/tools/go/src/text/template/parse/parse.go:162:16: method: shadowing a predeclared identifier (error)
/Users/yoshiki/tools/go/src/text/template/parse/parse.go:190:16: method: shadowing a predeclared identifier (recover)
/Users/yoshiki/tools/go/src/text/template/template.go:108:20: method: shadowing a predeclared identifier (copy)
/Users/yoshiki/tools/go/src/text/template/template.go:213:30: param: shadowing a predeclared identifier (new)
/Users/yoshiki/tools/go/src/text/template/exec_test.go:226:21: param: shadowing a predeclared identifier (error)
```
