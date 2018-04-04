package main

import (
	"imooc/crawer/engine"
	"imooc/crawer/persist"
	"imooc/crawer/scheduler"
	"imooc/crawer/zhenai/parser"
)

func main() {
	//http://www.zhenai.com/zhenghun
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParserFunc: parser.ParseCityList})

}
