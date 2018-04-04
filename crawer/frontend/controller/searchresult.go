package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"imooc/crawer/frontend/view"

	"gopkg.in/olivere/elastic.v5"
)

type SearchResultHandler struct {
	view   view.SearchReusltView
	client *elastic.Client
}

//localhost:9200/search?q=男 &from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	fmt.Fprintf(w, "q=%s,form=%d", q, from)
}

func CreateSearchReusltHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}
