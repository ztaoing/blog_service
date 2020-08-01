/**
* @Author:zhoutao
* @Date:2020/7/31 下午5:50
* @Desc:针对标签模块
 */

package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")

	ErrorDeleteTagFail = NewError(20010004, "删除标签失败")
	ErrorCountTagFail  = NewError(20010005, "统计标签失败")

	ErrorGetArticleFial    = NewError(20020001, "获取单个文章失败")
	ErrorGetArticlesFial   = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFial = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)
