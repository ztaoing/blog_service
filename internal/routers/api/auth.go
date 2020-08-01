/**
* @Author:zhoutao
* @Date:2020/8/1 下午4:22
* @Desc:
 */

package api

import (
	"blog_service/global"
	"blog_service/internal/service"
	"blog_service/pkg/app"
	"blog_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		global.Logger.Errorf("app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svs := service.New(c.Request.Context())
	//check
	err := svs.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svs.CheckAuth err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	//generate
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"token": token})
	return

}
