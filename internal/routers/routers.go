/**
* @Author:zhoutao
* @Date:2020/7/29 下午10:24
 */

package routers

import (
	_ "blog_service/docs"
	v1 "blog_service/internal/routers/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//_ "blog_service/docs"需要初始化doc
	//注册swagger路由
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)       //增加标签
		apiv1.DELETE("/tags/:id", tag.Delete) //删除指定标签
		apiv1.PUT("/tags/:id", tag.Update)    //更新指定标签
		apiv1.GET("/tags", tag.Get)           //获取标签列表

		apiv1.POST("/articles", article.Create)            //增加文章
		apiv1.DELETE("/articles/:id", article.Delete)      //删除指定文章
		apiv1.PUT("/articles/:id", article.Update)         //更新指定文章
		apiv1.PATCH("/articles/:id/state", article.Update) //更新指定文章的状态
		apiv1.GET("/articles/:id", article.Get)            //获取指定文章
		apiv1.GET("/articles", article.Get)                //获取文章列表
	}
	return r
}
