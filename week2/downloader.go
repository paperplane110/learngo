package main

import (
	"fmt"
	"github.com/paperplane110/learngo/week2/infra"
)

// Something that can Get

func getRetriever() retriever {
	return infra.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	r := getRetriever()
	url := "https://www.imooc.com"
	fmt.Println(r.Get(url))
}
