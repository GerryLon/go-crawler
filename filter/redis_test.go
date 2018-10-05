package filter

import (
	"fmt"
	"github.com/GerryLon/go-crawler/utils/text"
	"testing"
)

func TestMD5(t *testing.T) {
	fmt.Println(text.MD5("hello"))
}

func TestRedisFilter(t *testing.T) {
	filter := RedisDedupFilter{}
	key := "http://www.zhenai.com/zhenghun/xian"

	// redis.conf bind 127.0.0.1
	// calling Has: dial tcp 192.168.31.65:6379: connectex: No connection could be made because the target machine actively refused it.
	// panic: interface conversion: interface {} is nil, not bool [recovered]
	hasKey := filter.Has(key)
	fmt.Println("hasKey", hasKey)

	setted := filter.Set(key)
	fmt.Println("setted: ", setted)

	value := filter.Get(key)
	fmt.Println("value:", value)
}
