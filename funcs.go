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

func main() {
	fmt.Println(eval(13, 4, "/"))
}
