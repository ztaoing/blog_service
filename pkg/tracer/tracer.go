/**
* @Author:zhoutao
* @Date:2020/8/1 下午9:49
* @Desc:
 */

package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName, agenHostPort string) (opentracing.Tracer, io.Closer, error) {
	//jaeger client config
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{ //采样
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,            //是否启用LoggingReporter
			BufferFlushInterval: time.Second * 1, //刷新缓冲区的频率
			LocalAgentHostPort:  agenHostPort,    //上报的agent地址
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	//设置全局的tracer，并不是某个供应商的追踪系统的对象s
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil

}
