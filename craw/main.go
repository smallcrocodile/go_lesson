package main

import (
	"imooc/craw/engine"
	"imooc/craw/scheduler"
	"imooc/craw/zhenai/parser"
)

func main() {
	//http://www.zhenai.com/zhenghun
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParaserFunc: parser.ParseCityList})

}
