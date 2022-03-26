package main

import (
	"fmt"
	"math"
)

// var aa = 3
// var ss = "skldfjds"
// var bb = true

var (
	aa = 3
	ss = "skldfjds"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "sdfds"
	fmt.Println(a, b, s)
}

func variableDeduction() {
	var a, b, s = 1, 2, "lsdjfksd"
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, s := 4, 5, "lsdjfksd"
	fmt.Println(a, b, s)
}

func consts() {
	const filename = "sldkfjsld.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	const (
		cpp = iota
		java
		python
	)

	// b, kb, mb, gb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(cpp, java, python)
	fmt.Println(b, kb, mb, gb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	consts()
	enums()
}
