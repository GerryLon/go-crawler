package parser

import (
	"github.com/GerryLon/go-crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://city.zhenai.com/[a-z\d]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(match[1]), // url
			Parser: ParseCity,
		})

		// match[2] is city name
		// result.Items = append(result.Items, string(match[2]))
	}

	return result
}
