/**
* @Author:zhoutao
* @Date:2020/7/31 上午7:00
* @desc:分页处理
 */

package app

import (
	"blog_service/global"
	"blog_service/pkg/convert"
	"github.com/gin-gonic/gin"
)

//获取页码
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

//获取每页数量
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize < 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

//获取偏移量
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) + pageSize
	}
	return result
}
