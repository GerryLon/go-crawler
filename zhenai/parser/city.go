package parser

import (
	"github.com/GerryLon/go-crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]+)</a>`)

// 城市详情页面还有其他城市的链接
var cityUrlRe = regexp.MustCompile(`<a\s+href="(http://www.zhenai.com/zhenghun/[^"]+)"[^>]*>([^<]+)</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		name := string(match[2])
		homepage := string(match[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: homepage,
			Parser: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, homepage)
			},
		})

		// match[2] is user's nickname
		result.Items = append(result.Items, name)
	}

	// relative cities
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(match[1]), // url
			Parser: ParseCity,
		})

		result.Items = append(result.Items, string(match[2]))
	}

	return result
}
