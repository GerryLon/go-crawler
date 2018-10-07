package filter

import (
	"fmt"
	"testing"
)

func TestMemoryDedupFilter(t *testing.T) {
	key := "http://www.zhenai.com/zhenghun/xian"

	filter := MemoryDedupFilter{}
	// filter.Set(key) // 模拟已经有这个key了

	hasKey := filter.Has(key)
	fmt.Println("hasKey", hasKey)

	setted := filter.Set(key)
	fmt.Println("setted: ", setted)

	hasKey = filter.Has(key)
	fmt.Println("hasKey", hasKey)

	value := filter.Get(key)
	fmt.Println("value:", value)

	fmt.Println("len:", filter.Len())
}
