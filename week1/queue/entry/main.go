package main

import (
	"fmt"
	"github.com/paperplane110/learngo/week1/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
}
