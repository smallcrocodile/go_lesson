package main

import (
	"net/http"

	"imooc/crawer/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawer/frontend/view")))
	http.Handle("/search", controller.CreateSearchReusltHandler("crawer/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
