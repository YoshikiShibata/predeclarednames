package a

import "fmt"

const new = 10 // want "const"

func foo(make string) { // want "param"
	var len string // want "variable"

	fmt.Println(len)
}

func foo2() (string string) { // want "result"
	return ""
}

type bar1 interface {
	make(int) // want "method"
}

type len interface { // want "type"
}

type X struct {
	new int // want "field"
}

func foo3() {
new: // want "label"
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			continue new
		}
	}
}

func foo4() {
	new := 10 // want "variable"
	fmt.Println(new)
}

func (x X) len() { // want "method"
}

func (new X) foo() { // want "receiver"
}
