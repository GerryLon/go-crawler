package engine

// 对爬虫而言，一个请求的格式
// 一个地址对应一个解析器
type Request struct {
	Url    string
	Parser func(contents []byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

// ParseResult中的一个条目
type Item struct {
	Id      string // 目标的id， 如uid值
	Type    string // 可配置， 比如“zhenai"
	Url     string // 一般为详情页面（如个人主页）
	Payload interface{}
}

// this parser do nothing
func NoopParser(contents []byte) ParseResult {
	return ParseResult{}
}
