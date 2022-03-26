package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div(a, b)
		return q
	default:
		panic("Unsupoort operation: " + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func swap(a, b int) {
	b, a = a, b
}

func pointerSwap(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	fmt.Println(eval(13, 4, "/"))
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b) // won't work, because pass-by-value

	pointerSwap(&a, &b)
	fmt.Println(a, b)
}
