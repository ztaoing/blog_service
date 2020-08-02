/**
* @Author:zhoutao
* @Date:2020/8/1 下午6:44
* @Desc:
 */

package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

//定义吸纳刘去必须的方法
type LimiterInterface interface {
	Key(c *gin.Context) string //获取对应的限流器的键值对名称
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterInterface //增加多个令牌桶
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

//限流令牌桶规则
type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration //间隔多久时间方N个令牌
	Capcity      int64         //令牌桶容量
	Quantum      int64         //每次到达间隔时间后所放的具体令牌数量
}
