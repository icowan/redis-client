/**
 * @Time : 2020/8/24 9:40 AM
 * @Author : solacowa@gmail.com
 * @File : client_test
 * @Software: GoLand
 */

package redisclient

import (
	"testing"
	"time"
)

func TestNewRedisClient(t *testing.T) {
	rds := NewRedisClient("127.0.0.1:6379", "admin", "", 1)
	defer func() {
		_ = rds.Close()
	}()

	_ = rds.Set("hello", "world", time.Second*10)
	v, err := rds.Get("hello")
	if err != nil {
		t.Error(err)
	}

	t.Log(v)
}

func TestNewRedisCluster(t *testing.T) {
	rds := NewRedisClient("127.0.0.1:6379,127.0.0.1:7379,127.0.0.1:8379", "admin", "", 0)
	defer func() {
		_ = rds.Close()
	}()

	_ = rds.Set("hello", "world", time.Second*10)
	v, err := rds.Get("hello")
	if err != nil {
		t.Error(err)
	}

	t.Log(v)
}
