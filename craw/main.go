package main

import (
	"imooc/craw/engine"
	"imooc/craw/zhenai/parser"
)

func main() {
	//http://www.zhenai.com/zhenghun

	engine.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParaserFunc: parser.ParseCityList})

}
