package main

import (
	"github.com/GerryLon/go-crawler/engine"
	"github.com/GerryLon/go-crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:    "http://city.zhenai.com/",
		Parser: parser.ParseCityList,
	})
}
