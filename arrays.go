package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 3, 4, 5, 6}
	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]int // 4 rows 5 columns
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr1 {
		fmt.Println(i)
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, v := range arr2 {
		fmt.Println(v)
	}

	printArray(arr1)
}
