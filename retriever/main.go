package main

import (
	"fmt"
	"time"

	"imooc/retriever/mock"
	"imooc/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"fake url"}
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

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println(r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}
