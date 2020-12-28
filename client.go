/**
 * @Time: 2020/3/30 10:38
 * @Author: solacowa@gmail.com
 * @File: client
 * @Software: GoLand
 */

package redisclient

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient interface {
	Set(ctx context.Context, k string, v interface{}, expir ...time.Duration) (err error)
	Get(ctx context.Context, k string) (v string, err error)
	Del(ctx context.Context, k string) (err error)
	Exists(ctx context.Context, keys ...string) int64
	HSet(ctx context.Context, k string, field string, v interface{}) (err error)
	HGet(ctx context.Context, k string, field string) (res string, err error)
	HGetAll(ctx context.Context, k string) (res map[string]string, err error)
	HLen(ctx context.Context, k string) (res int64, err error)
	ZCard(ctx context.Context, k string) (res int64, err error)
	ZRangeWithScores(ctx context.Context, k string, start, stop int64) (res []redis.Z, err error)
	ZAdd(ctx context.Context, k string, score float64, member interface{}) (err error)
	HDelAll(ctx context.Context, k string) (err error)
	HDel(ctx context.Context, k string, field string) (err error)
	Keys(ctx context.Context, pattern string) (res []string, err error)
	LLen(ctx context.Context, key string) int64
	RPop(ctx context.Context, key string) (res string, err error)
	LPush(ctx context.Context, key string, val interface{}) (err error)
	TypeOf(ctx context.Context, key string) (res string, err error)
	Close(ctx context.Context) error
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	Publish(ctx context.Context, channel string, message interface{}) error
	Incr(ctx context.Context, key string, exp time.Duration) error
	SetPrefix(ctx context.Context, prefix string) RedisClient
	TTL(ctx context.Context, key string) time.Duration
	Ping(ctx context.Context) error
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

	if err = rds.Ping(context.Background()); err != nil {
		return nil, err
	}
	return rds, nil
}
