package engine

import (
	"github.com/GerryLon/go-crawler/config"
	"github.com/GerryLon/go-crawler/filter"
)

// dedup interface
type Deduper interface {
	ConfigFilter(filter.Filter) // config dedup filter: memory, redis, or more...
	isDuplicate(string) bool

	// Len 方法不应该强制要求实现， 毕竟不是核心方法
	//Len() int
}

type DefaultDeduper struct {
	filter filter.Filter
}

func (d *DefaultDeduper) ConfigFilter(filter filter.Filter) {
	d.filter = filter
}

func (d *DefaultDeduper) isDuplicate(url string) bool {
	if config.WillDeDup {
		if d.filter.Has(url) {
			return true
		} else {
			d.filter.Set(url)
			return false
		}
	}
	return false
}

func (d *DefaultDeduper) Len() int {
	return d.filter.Len()
}
