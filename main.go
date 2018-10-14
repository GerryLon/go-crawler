package main

import (
	"github.com/GerryLon/go-crawler/config"
	"github.com/GerryLon/go-crawler/engine"
	"github.com/GerryLon/go-crawler/filter"
	"github.com/GerryLon/go-crawler/persist"
	"github.com/GerryLon/go-crawler/scheduler"
	"github.com/GerryLon/go-crawler/zhenai/parser"
)

func main() {
	// config dedup filter
	deduper := engine.DefaultDeduper{}
	deduper.ConfigFilter(&filter.RedisDedupFilter{})

	// config elastic search
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	// use concurrent engine
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		// 配置去重过滤器
		Deduper: &deduper,
	}

	// use simple engine
	//e := engine.SimpleEngine{
	//	Deduper:  &deduper,
	//	ItemChan: itemChan,
	//}

	e.Run(engine.Request{
		Url:    "http://city.zhenai.com/",
		Parser: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:    "http://www.zhenai.com/zhenghun/shenzhen",
	//	Parser: parser.ParseCity,
	//})
}
