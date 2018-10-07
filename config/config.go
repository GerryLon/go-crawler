package config

// config for redigo(https://github.com/gomodule/redigo)
const (
	RedisHost          = "192.168.31.65"
	RedisPort          = 6379
	RedisPassword      = "gerrylon"
	RedisPoolMaxActive = 10
	RedisPoolMaxIdle   = 5
	RedisHSetKey       = "zhenaiwang"
	QPS                = 20
	WillDeDup          = false // 是否启用去重
)
