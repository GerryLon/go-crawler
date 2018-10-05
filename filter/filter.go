package filter

// 对爬取过程中的url去重

type Filter interface {
	Has(key string) bool
	Get(key string) string
	Set(key string) bool
}

// de duplication filter
type DedupFilter struct {
}
