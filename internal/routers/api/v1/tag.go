/**
* @Author:zhoutao
* @Date:2020/7/29 下午10:32
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

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

//@Summary 获取多个标签
//@Produce json
//@Param name query string false "标签名称" maxlength(100)
//@Param state query int false "状态" Enums(0,1) default(1)
//@Param page query int false "页码"
//@Param page_size query int false "每页数量"
//@Success 200 {object} model.TagSwagger "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {

	param := service.TagListRequest{}
	response := app.NewResponse(c)
	//参数校验、绑定、
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svs := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	//获取标签总数
	totalRows, err := svs.CountTag(&service.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		global.Logger.Errorf(c, "svs.CountTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	//获取标签列表
	tags, err := svs.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svs.GetTagList err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	//序列化结果集
	response.ToResponseList(tags, totalRows)
	return

}

//@Summary 新建标签
//@Produce json
//@Param name body string true "标签名称" minlength(3) maxlength(100)
//@Param state body int false "状态" Enum(0,1) default(1)
//@Param created_by body string false "创建者" minlength(3) maxlength(100)
//@Success 200 {object} model.TagSwagger "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svs := service.New(c.Request.Context())
	err := svs.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svs.CreateTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

//@Summary 更新标签
//@Produce json
//@Param id path int true "标签ID"
//@Param name body string false "标签名称" minlength(3) maxlength(100)
//@Param state body int false "状态" Enum(0,1) default(1)
//@Param create_by body string false "创建者" minlength(3) maxlength(100)
//@Success 200 {object} model.TagSwagger "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svs := service.New(c.Request.Context())
	err := svs.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svs.UpdateTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(param)
	return

}

//@Summary 删除标签
//@Produce json
//@param id  path int true "标签ID"
//@Success 200 {string} string "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Faulure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}
	svs := service.New(c.Request.Context())
	err := svs.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svs.DeleteTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
