package filter

// 对爬取过程中的url去重

type Filter interface {
	Has(key string) bool
	Get(key string) string
	Set(key string) bool
	Del(key string) bool
}

// dereplication filter
type DuplicateFilter struct {
}

func (DuplicateFilter) Has(key string) bool {
	return false
}

func (DuplicateFilter) Get(key string) string {
	return ""
}

func (DuplicateFilter) Set(key string) bool {
	return false
}

func (DuplicateFilter) Del(key string) bool {
	return false
}
