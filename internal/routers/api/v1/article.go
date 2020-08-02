/**
* @Author:zhoutao
* @Date:2020/7/29 下午10:33
 */

package v1

import (
	"blog_service/global"
	"blog_service/internal/service"
	"blog_service/pkg/app"
	"blog_service/pkg/convert"
	"blog_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svs := service.New(c.Request.Context())
	article, err := svs.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf("svs.GetArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFial)
		return
	}
	response.ToResponse(article)
	return
}
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if errs != nil {
		global.Logger.Errorf("app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svs := service.New(c)
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	articles, totalRows, err := svs.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svs.GetArticleList err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFial)
		return
	}
	response.ToResponseList(articles, totalRows)
	return
}
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svs := service.New(c)
	err := svs.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svs.CreateArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFial)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svs := service.New(c)
	err := svs.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svs.UpdateArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svs := service.New(c)
	err := svs.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf("svs.DeleteArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	retur
}
