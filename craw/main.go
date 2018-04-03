package main

import (
	"imooc/craw/engine"
	"imooc/craw/persist"
	"imooc/craw/scheduler"
	"imooc/craw/zhenai/parser"
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
