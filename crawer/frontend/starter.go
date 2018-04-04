package main

import (
	"net/http"

	"imooc/crawer/frontend/controller"
)

func main() {
	http.Handle("/search", controller.CreateSearchReusltHandler("/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
