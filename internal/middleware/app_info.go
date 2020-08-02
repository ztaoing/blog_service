/**
* @Author:zhoutao
* @Date:2020/8/1 下午6:39
* @Desc:服务信息存储 根据不同的租户号获取不同的数据库实例对象
 */

package middleware

import "github.com/gin-gonic/gin"

//setter getter 是gin.Context提供的元数据管理
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog-service")
		c.Set("app-version", "1.0.0")
		c.Next()
	}
}
