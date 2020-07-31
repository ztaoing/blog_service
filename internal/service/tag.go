/**
* @Author:zhoutao
* @Date:2020/7/31 上午9:56
* @desc:针对入参校验增加绑定和验证结构体
* @form :表单的映射字段名  binding:入参校验规则
 */

package service

type CountTagRequest struct {
	State uint8 `form:"state:default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1""`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
