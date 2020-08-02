/**
* @Author:zhoutao
* @Date:2020/8/1 下午10:15
* @Desc:链路追踪 gin和tracer的衔接
 */

package middleware

import (
	"blog_service/global"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path)
		}
		defer span.Finish()
		//获取traceID和SpanId
		var traceID string
		var SpanId string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			traceID = spanContext.(jaeger.SpanContext).TraceID().String()
			SpanId = spanContext.(jaeger.SpanContext).SpanID().String()
		}
		//设置到上下文中元数据中
		c.Set("X-trace-ID", traceID)
		c.Set("X-span-ID", SpanId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
