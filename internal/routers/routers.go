/**
* @Author:zhoutao
* @Date:2020/7/29 下午10:24
 */

package routers

import (
	_ "blog_service/docs"
	"blog_service/global"
	"blog_service/internal/middleware"
	"blog_service/internal/routers/api"
	v1 "blog_service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.TransLations())
	//_ "blog_service/docs"需要初始化doc
	//注册swagger路由
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()

	//文件上传
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	//增加对静态资源的访问
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	//auth
	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	//令牌校验
	apiv1.Use(middleware.Jwt()) //middleware.JWT()
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
