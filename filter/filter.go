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

func (filter *DuplicateFilter) Has(key string) bool {
	return false
}

func (filter *DuplicateFilter) Get(key string) string {
	return ""
}

func (filter *DuplicateFilter) Set(key string) bool {
	return false
}

func (filter *DuplicateFilter) Del(key string) bool {
	return false
}
