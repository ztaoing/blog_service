/**
* @Author:zhoutao
* @Date:2020/8/1 下午6:52
* @Desc:
 */

package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}

//根据requestURI切割出核心路由作为键值对名称
func (m MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

//选择要使用的令牌桶
func (m MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := m.limiterBuckets[key]
	return bucket, ok
}

func (m MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterInterface {
	for _, rule := range rules {
		if _, ok := m.limiterBuckets[rule.Key]; !ok {
			m.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capcity, rule.Quantum)
		}
	}
	return m
}

func NewMethodLimiter() LimiterInterface {
	return MethodLimiter{
		&Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}
