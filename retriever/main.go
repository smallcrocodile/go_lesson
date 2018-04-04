package main

import (
	fmt "fmt"
	"time"

	"imooc/retriever/mock"
	"imooc/retriever/real"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{
		"name":   "jack",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"fake url"}
	r = &retriever

	inspect(r)
	r = &real.Retriever{UserAgent: "Mozile/5.0", TimeOut: time.Minute}
	inspect(r)
	//fmt.Printf(download(r))

	//Type assertion
	if mockRetriver, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriver.Contents)
	} else {
		fmt.Println("not mock retriver")

	}
	fmt.Println("try a session with mockRetriever")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)

	fmt.Printf("%T %v\n", r, r)
	fmt.Println(r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}
