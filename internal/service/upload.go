/**
* @Author:zhoutao
* @Date:2020/8/1 上午10:06
* @Desc:
 */

package service

import (
	"blog_service/global"
	"blog_service/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

//上传文件
func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	//检测文件类型
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	//检测文件路径
	if upload.CheckSavePath(dst) {
		if err := upload.CreateSavePath(dst, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	//检测文件大小
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("execceded maximum file limit.")
	}
	//检测权限
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permission.")
	}
	//写入文件中
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessUrl,
	}, nil
}
