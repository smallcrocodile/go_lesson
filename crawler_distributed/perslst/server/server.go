package main

import (
	"imooc/crawler_distributed/perslst"
	"imooc/crawler_distributed/rpcsupport"

	"gopkg.in/olivere/elastic.v5"
)

func main() {
	serveRpc(":1234", "dating_profile")
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServerRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
