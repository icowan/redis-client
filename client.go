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
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"strings"
	"time"
)

func NewRedisClient(hosts, password, prefix string, db int, stdTracer opentracing.Tracer) (rdb redis.UniversalClient, err error) {
	prefixFn := func(key string) string {
		if strings.EqualFold(prefix, "") {
			return key
		}
		prefix = strings.Trim(prefix, ":")
		key = strings.Trim(key, ":")
		return fmt.Sprintf("%s:%s", prefix, key)
	}
	h := strings.Split(hosts, ",")
	rdb = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            h,
		Password:         password,
		DB:               db,
		DialTimeout:      time.Second * 15,
		ReadTimeout:      time.Second * 15,
		WriteTimeout:     time.Second * 15,
		MasterName:       "",
		SentinelUsername: "",
		SentinelPassword: "",
	})
	if stdTracer != nil {
		rdb.AddHook(newPrefixHook(prefixFn, stdTracer))
	}
	if err = rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
