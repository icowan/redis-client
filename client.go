/**
 * @Time: 2020/3/30 10:38
 * @Author: solacowa@gmail.com
 * @File: client
 * @Software: GoLand
 */

package redisclient

import (
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient interface {
	Set(k string, v interface{}, expir ...time.Duration) (err error)
	Get(k string) (v string, err error)
	Del(k string) (err error)
	Exists(keys ...string) int64
	HSet(k string, field string, v interface{}) (err error)
	HGet(k string, field string) (res string, err error)
	HGetAll(k string) (res map[string]string, err error)
	HLen(k string) (res int64, err error)
	ZCard(k string) (res int64, err error)
	ZRangeWithScores(k string, start, stop int64) (res []redis.Z, err error)
	ZAdd(k string, score float64, member interface{}) (err error)
	HDelAll(k string) (err error)
	HDel(k string, field string) (err error)
	Keys(pattern string) (res []string, err error)
	LLen(key string) int64
	RPop(key string) (res string, err error)
	LPush(key string, val interface{}) (err error)
	TypeOf(key string) (res string, err error)
	Close() error
	Subscribe(channels ...string) *redis.PubSub
	Publish(channel string, message interface{}) error
	Incr(key string, exp time.Duration) error
	SetPrefix(prefix string) RedisClient
	TTL(key string) time.Duration
	Ping() error
}

const (
	RedisCluster = "cluster"
	RedisSingle  = "single"
	expiration   = 600 * time.Second
)

func NewRedisClient(hosts, password, prefix string, db int) (rds RedisClient, err error) {
	h := strings.Split(hosts, ",")
	if len(h) > 1 {
		rds = NewRedisCluster(
			h,
			password,
			prefix,
		)
	} else {
		rds = NewRedisSingle(
			hosts,
			password,
			prefix,
			db,
		)
	}

	if err = rds.Ping(); err != nil {
		return nil, err
	}
	return rds, nil
}
