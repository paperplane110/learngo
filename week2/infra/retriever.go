package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct{}

func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // close io stream when process is finished

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
