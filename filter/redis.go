package filter

import (
	"fmt"
	"github.com/GerryLon/go-crawler/config"

	"github.com/GerryLon/go-crawler/utils"
	"github.com/GerryLon/go-crawler/utils/text"
	redigo "github.com/gomodule/redigo/redis"
	"log"
)

// de duplication using redis
type RedisDedupFilter struct {
	DuplicateFilter
	pool *redigo.Pool
}

// init redis connection pool
func (filter *RedisDedupFilter) init() {
	filter.pool = &redigo.Pool{
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort))
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", config.RedisPassword); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
		MaxActive: config.RedisPoolMaxActive,
		MaxIdle:   config.RedisPoolMaxIdle,
	}
}

// get one connection from redis connection pool
func getRedisConn(filter *RedisDedupFilter) redigo.Conn {
	if filter.pool == nil {
		filter.init()
	}

	conn := filter.pool.Get()

	return conn
}

// redis: exists key
func (filter *RedisDedupFilter) Has(key string) bool {
	conn := getRedisConn(filter)
	defer conn.Close()

	reply, err := conn.Do("HEXISTS", config.RedisHSetKey, text.MD5(key))
	if err != nil {
		log.Printf("error occured when calling Has: %v", err)
	}

	return utils.Itob(int(reply.(int64)))
}

//redis: get mdt(key)
func (filter *RedisDedupFilter) Get(key string) string {
	conn := getRedisConn(filter)
	defer conn.Close()
	md5OfKey := text.MD5(key)
	reply, err := conn.Do("HGET", config.RedisHSetKey, md5OfKey)

	if err != nil {
		log.Printf("error occured when calling Get(%s): %v", key, err)
	}

	if reply == nil {
		return ""
	}

	return string(reply.([]uint8))
}

// redis: setnx key
func (filter *RedisDedupFilter) setNX(key string) bool {
	conn := getRedisConn(filter)
	defer conn.Close()
	md5OfKey := text.MD5(key)
	reply, err := conn.Do("HSETNX", config.RedisHSetKey, md5OfKey, key)

	if err != nil {
		log.Printf("error occured when calling Get(%s): %v", key, err)
	}

	return utils.Itob(int(reply.(int64)))
}

// set using redis: setnx key
func (filter *RedisDedupFilter) Set(key string) bool {
	return filter.setNX(key)
}

func (filter *RedisDedupFilter) Del(key string) bool {
	return false
}
