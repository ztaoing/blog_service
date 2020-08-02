/**
* @Author:zhoutao
* @Date:2020/8/1 下午8:43
* @Desc:超时控制
 */

package middleware

import (
	"context"
	"github.com/gin-gonic/gin"

	"time"
)

func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
