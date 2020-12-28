/**
 * @Time : 2020/8/24 9:40 AM
 * @Author : solacowa@gmail.com
 * @File : client_test
 * @Software: GoLand
 */

package redisclient

import (
	"context"
	"testing"
	"time"
)

var (
	ctx context.Context
)

func TestNewRedisClient(t *testing.T) {
	rds, err := NewRedisClient("127.0.0.1:6379", "admin", "", 1)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = rds.Close(ctx)
	}()

	_ = rds.Set(ctx, "hello", "world", time.Second*10)
	v, err := rds.Get(ctx, "hello")
	if err != nil {
		t.Error(err)
	}

	t.Log(v)
}

func TestNewRedisCluster(t *testing.T) {
	ctx = context.Background()
	rds, err := NewRedisClient("localhost:32769", "", "", 0)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = rds.Close(ctx)
	}()

	_ = rds.Set(ctx, "hello", "world", time.Second*10)
	v, err := rds.Get(ctx, "hello")
	if err != nil {
		t.Error(err)
	}

	t.Log(v)
}
