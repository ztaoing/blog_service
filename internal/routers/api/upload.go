/**
* @Author:zhoutao
* @Date:2020/8/1 上午10:23
* @Desc:
 */

package api

import (
	"blog_service/global"
	"blog_service/internal/service"
	"blog_service/pkg/app"
	"blog_service/pkg/convert"
	"blog_service/pkg/errcode"
	"blog_service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	//读取上传文件信息
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svs := service.New(c.Request.Context())
	fileInfo, err := svs.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svs.UploadFile err:%v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
	return
}
