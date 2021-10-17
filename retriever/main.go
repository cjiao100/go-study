package main

import (
	"fmt"
	"time"

	"github.com/cjiao100/learngo/retriever/real"

	"github.com/cjiao100/learngo/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("https://www.imooc.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		},
	)
}

// RetrieverPoster 接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})

	return s.Get(url)
}

func main() {
	// 它里面会有当前的类型和值
	var r Retriever

	inspect(r)
	//fmt.Printf("%T %v", r, r)

	r = &mock.Retriever{Contents: "this is a fake imooc.com"}

	inspect(r)
	//fmt.Printf("%T %v", r, r)
	//fmt.Println(download(r))

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	//fmt.Printf("%T %v", r, r)
	//fmt.Println(download(r))

	inspect(r)

	// r.(mock.Retriever) => type assertion 可以通过.来获取r现在的类型，如果不是当前的类型，会报错。或者可以加第二个参数获取是否正确
	if realRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//session()

	fmt.Println("Try a session")
	s := mock.Retriever{Contents: "this is a fake imooc.com"}
	fmt.Println(session(&s))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
