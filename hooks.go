/**
 * @Time : 2022/3/30 11:02 AM
 * @Author : solacowa@gmail.com
 * @File : hooks
 * @Software: GoLand
 */

package redisclient

import (
	"context"
	"reflect"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/redis/go-redis/v9"
)

type prefixHook struct {
	prefixFn func(key string) string
	tracer   opentracing.Tracer
}

func (s prefixHook) DialHook(next redis.DialHook) redis.DialHook {
	return next
}

func (s prefixHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return next
}

func (s prefixHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return next
}

func (s prefixHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "BeforeProcess", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "redisClient.hooks",
	})
	defer func() {
		span.LogKV("cmd", cmd.Args())
		span.Finish()
	}()
	if len(cmd.Args()) > 1 && reflect.TypeOf(cmd.Args()[1]).Kind() == reflect.String {
		cmd.Args()[1] = s.prefixFn(cmd.Args()[1].(string))
	}
	return ctx, nil
}

func (s prefixHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	return nil
}

func (s prefixHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (s prefixHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}

func newPrefixHook(prefixFn func(key string) string, stdTracer opentracing.Tracer) redis.Hook {
	return &prefixHook{prefixFn, stdTracer}
}
