package main

import (
	"imooc/craw/engine"
	"imooc/craw/persist"
	"imooc/craw/scheduler"
	"imooc/craw/zhenai/parser"
)

func main() {
	//http://www.zhenai.com/zhenghun
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParaserFunc: parser.ParseCityList})
	//e.Run(engine.Request{
	//	Url:         "http://www.zhenai.com/zhenghun/shanghai",
	//	ParaserFunc: parser.ParseCity,
	//})
}
