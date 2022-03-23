package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func intToBinary(n int) string {
	ans := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		ans = strconv.Itoa(lsb) + ans
	}
	return ans
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		intToBinary(5),
		intToBinary(231231),
		intToBinary(0),
	)

	printFile("abc.txt")
}
