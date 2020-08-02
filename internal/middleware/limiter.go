/**
* @Author:zhoutao
* @Date:2020/8/1 下午8:31
* @Desc:
 */

package middleware

import (
	"blog_service/pkg/app"
	"blog_service/pkg/errcode"
	"blog_service/pkg/limiter"
	"github.com/gin-gonic/gin"
)

//将限流器与对应的中间件串联起来
//注意：入参是LimiterInterface接口类型，只要符合该接口类型的具体限流器实现都可以传入并使用
func ReteLimiter(l limiter.LimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			//返回删除的令牌数
			count := bucket.TakeAvailable(1)
			//没有令牌的时候
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}

		}
		c.Next()
	}
}
