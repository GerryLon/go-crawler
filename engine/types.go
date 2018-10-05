package engine

// 对爬虫而言，一个请求的格式
// 一个地址对应解析器
type Request struct {
	Url string
	Parser func(contents []byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
}

// this parser do nothing
func NoopParser(contents []byte) ParseResult {
	return ParseResult{}
}