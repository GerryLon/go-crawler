package main

import (
	"github.com/GerryLon/go-crawler/engine"
	"github.com/GerryLon/go-crawler/scheduler"
	"github.com/GerryLon/go-crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:    "http://city.zhenai.com/",
		Parser: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:    "http://www.zhenai.com/zhenghun/shenzhen",
	//	Parser: parser.ParseCity,
	//})
}
