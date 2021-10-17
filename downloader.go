package main

import (
	"fmt"

	"github.com/cjiao100/learngo/testing"
)

//func retrieve(url string) string {
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//
//	defer resp.Body.Close()
//
//	bytes, _ := ioutil.ReadAll(resp.Body)
//
//	return string(bytes)
//}

func getRetriever() retriever {
	return testing.Retriever{}
}

// something that can "Get"
type retriever interface {
	Get(string) string
}

func main() {
	// 对于main函数来说，不需要关心retriever具体是来自于那个包，不要他们能实现接口要求的就可以
	var r retriever = getRetriever()

	fmt.Println(r.Get("https://www.imooc.com"))
}
