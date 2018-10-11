package engine

import (
	"github.com/GerryLon/go-crawler/config"
	"github.com/GerryLon/go-crawler/filter"
)

// dedup interface
type Deduper interface {
	ConfigFilter(filter.Filter) // config dedup filter: memory, redis, or more...
	isDuplicate(string) bool
	DeduperInfo
}

// 为了避免将Len等非核心的方法放在Deduper中
// 因为去重必须实现isDuplicate方法， 而Len方法不是必须的
type DeduperInfo interface {
	Len() int
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
