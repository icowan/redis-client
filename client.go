/**
 * @Time: 2020/3/30 10:38
 * @Author: solacowa@gmail.com
 * @File: client
 * @Software: GoLand
 */

package redisclient

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(hosts, password, prefix string, db int, stdTracer opentracing.Tracer) (rdb redis.Cmdable, err error) {
	prefixFn := func(key string) string {
		prefix = strings.Trim(prefix, ":")
		key = strings.Trim(key, ":")
		return fmt.Sprintf("%s:%s", prefix, key)
	}
	h := strings.Split(hosts, ",")
	if len(h) > 1 {
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        h,
			Password:     password,
			DialTimeout:  time.Second * 15,
			ReadTimeout:  time.Second * 15,
			WriteTimeout: time.Second * 15,
		})
		if stdTracer != nil {
			client.AddHook(newPrefixHook(prefixFn, stdTracer))
		}
		rdb = client
	} else {
		client := redis.NewClient(&redis.Options{
			Addr:     hosts,
			Password: password, // no password set
			DB:       db,       // use default DB
		})
		if stdTracer != nil {
			client.AddHook(newPrefixHook(prefixFn, stdTracer))
		}
		rdb = client
	}

	if err = rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
