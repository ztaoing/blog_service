/**
* @Author:zhoutao
* @Date:2020/8/1 上午8:52
* @Desc:
 */

package upload

import (
	"blog_service/global"
	"blog_service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeTmage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	//去掉后缀
	fileName := strings.TrimSuffix(name, ext)
	//对文件名进行加密
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

//扩展名
func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

//检查文件目录是否存在
func CheckSavePath(dst string) bool {
	// Stat returns a FileInfo describing the named file.
	//获取文件的描述信息
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

//检查是否是支持的类型
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeTmage:
		//检查是否支持此图片类型
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
				return true
			}
		}

	}
	return false
}

//检查文件大小
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeTmage:
		if size >= global.AppSetting.MaxPageSize*1024*1024 {
			return true
		}
	}
	return false
}

//检查权限
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

//创建保存上传文件的目录
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

//保存上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	//打开原地址文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建目标地址文件
	out, err := os.Create(dst)
	defer out.Close()

	_, err = io.Copy(out, src)
	return err

}
