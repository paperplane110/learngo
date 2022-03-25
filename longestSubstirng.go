package main

import "fmt"

func lengthOfNonRepeatingSubsting(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	for index, ch := range []rune(s) {
		lastIndex, ok := lastOccured[ch]
		if ok == true && lastIndex >= start {
			start = lastIndex + 1
		}
		if index-start+1 > maxLength {
			maxLength = index - start + 1
		}
		lastOccured[ch] = index
	}
	return maxLength
}

func main() {
	s := "一二三三二一四"
	l := lengthOfNonRepeatingSubsting(s)
	fmt.Println(l)
}
