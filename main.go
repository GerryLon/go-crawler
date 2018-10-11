package main

import (
	"github.com/GerryLon/go-crawler/engine"
	"github.com/GerryLon/go-crawler/filter"
	"github.com/GerryLon/go-crawler/scheduler"
	"github.com/GerryLon/go-crawler/zhenai/parser"
)

func main() {
	deduper := engine.DefaultDeduper{}

	// config dedup filter
	deduper.ConfigFilter(&filter.RedisDedupFilter{})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,

		// 配置去重过滤器
		Deduper: &deduper,
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
