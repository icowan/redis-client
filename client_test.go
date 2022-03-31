/**
 * @Time : 2020/8/24 9:40 AM
 * @Author : solacowa@gmail.com
 * @File : client_test
 * @Software: GoLand
 */

package redisclient

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	"io"
	"testing"
	"time"
)

var (
	ctx context.Context
)

func TestNewRedisClient(t *testing.T) {
	tracer, _, _ := newJaegerTracer()
	ctx = context.Background()
	rds, err := NewRedisClient("127.0.0.1:32768", "", "rds", 1, tracer)
	if err != nil {
		fmt.Println("ping......", err)
		t.Error(err)
	}
	err = rds.Set(ctx, "hello", "world", time.Second*50).Err()
	v := rds.Get(ctx, "hello").Val()
	t.Log(v)
}

// 使用jaeger
func newJaegerTracer() (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},
		Reporter: &jaegerConfig.ReporterConfig{
			//QueueSize:          200, // 缓冲区越大内存消耗越大,默认100
			LogSpans:           true,
			LocalAgentHostPort: "",
		},
		ServiceName: fmt.Sprintf("%s.%s", "redis", "client"),
	}
	metricsFactory := prometheus.New()
	tracer, closer, err = cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger), jaegerConfig.Metrics(metricsFactory))
	if err != nil {
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}
