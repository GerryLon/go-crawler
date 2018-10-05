package parser

import (
	"github.com/GerryLon/go-crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]+)</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]), // url
			Parser: engine.NoopParser,
		})

		// match[2] is user's nickname
		result.Items = append(result.Items, string(match[2]))
	}

	return result
}
