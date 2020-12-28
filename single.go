/**
 * @Time: 2020/3/30 10:39
 * @Author: solacowa@gmail.com
 * @File: single
 * @Software: GoLand
 */

package redisclient

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type single struct {
	client *redis.Client
	prefix string
}

func (c *single) Ping(_ context.Context) error {
	return c.client.Ping().Err()
}

func (c *single) LPush(_ context.Context, key string, val interface{}) (err error) {
	return c.client.LPush(key, val).Err()
}

func (c *single) RPop(_ context.Context, key string) (res string, err error) {
	return c.client.RPop(key).Result()
}

func (c *single) LLen(_ context.Context, key string) int64 {
	return c.client.LLen(key).Val()
}

func (c *single) TypeOf(_ context.Context, key string) (res string, err error) {
	return c.client.Type(key).Result()
}

func (c *single) Keys(_ context.Context, pattern string) (res []string, err error) {
	return c.client.Keys(pattern).Result()
}

func (c *single) ZAdd(_ context.Context, k string, score float64, member interface{}) (err error) {
	return c.client.ZAdd(c.setPrefix(k), redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

func (c *single) ZCard(_ context.Context, k string) (res int64, err error) {
	return c.client.ZCard(c.setPrefix(k)).Result()
}

func (c *single) ZRangeWithScores(_ context.Context, k string, start, stop int64) (res []redis.Z, err error) {
	return c.client.ZRangeWithScores(c.setPrefix(k), start, stop).Result()
}

func (c *single) HLen(_ context.Context, k string) (res int64, err error) {
	return c.client.HLen(c.setPrefix(k)).Result()
}

func (c *single) HGetAll(_ context.Context, k string) (res map[string]string, err error) {
	return c.client.HGetAll(c.setPrefix(k)).Result()
}

func (c *single) Exists(_ context.Context, keys ...string) int64 {
	return c.client.Exists(keys...).Val()
}

func (c *single) TTL(_ context.Context, key string) time.Duration {
	return c.client.TTL(key).Val()
}

func NewRedisSingle(host, password, prefix string, db int) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &single{client: client, prefix: prefix}
}

func (c *single) Set(_ context.Context, k string, v interface{}, expir ...time.Duration) (err error) {
	var val string
	switch v.(type) {
	case string:
		val = v.(string)
	default:
		b, _ := json.Marshal(v)
		val = string(b)
	}

	exp := expiration
	if len(expir) == 1 {
		exp = expir[0]
	}

	return c.client.Set(c.setPrefix(k), val, exp).Err()
}

func (c *single) Get(_ context.Context, k string) (v string, err error) {
	return c.client.Get(c.setPrefix(k)).Result()
}

func (c *single) Del(_ context.Context, k string) (err error) {
	return c.client.Del(c.setPrefix(k)).Err()
}

func (c *single) HSet(_ context.Context, k string, field string, v interface{}) (err error) {
	var val string
	switch v.(type) {
	case string:
		val = v.(string)
	default:
		b, _ := json.Marshal(v)
		val = string(b)
	}
	return c.client.HSet(c.setPrefix(k), field, val).Err()
}

func (c *single) HGet(_ context.Context, k string, field string) (res string, err error) {
	return c.client.HGet(c.setPrefix(k), field).Result()
}

func (c *single) HDelAll(_ context.Context, k string) (err error) {
	res, err := c.client.HKeys(c.setPrefix(k)).Result()
	if err != nil {
		return
	}
	return c.client.HDel(c.setPrefix(k), res...).Err()
}

func (c *single) HDel(_ context.Context, k string, field string) (err error) {
	return c.client.HDel(c.setPrefix(k), field).Err()
}

func (c *single) setPrefix(s string) string {
	return c.prefix + s
}

func (c *single) Close(_ context.Context) error {
	return c.client.Close()
}

func (c *single) Subscribe(_ context.Context, channels ...string) *redis.PubSub {
	return c.client.Subscribe(channels...)
}

func (c *single) Publish(_ context.Context, channel string, message interface{}) error {
	return c.client.Publish(channel, message).Err()
}

func (c *single) Incr(_ context.Context, key string, expiration time.Duration) error {
	defer func() {
		c.client.Expire(c.setPrefix(key), expiration)
	}()
	return c.client.Incr(c.setPrefix(key)).Err()
}

func (c *single) SetPrefix(_ context.Context, prefix string) RedisClient {
	c.prefix = prefix
	return c
}
