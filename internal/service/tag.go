/**
* @Author:zhoutao
* @Date:2020/7/31 上午9:56
* @desc:处理标签模块的业务逻辑。针对入参校验增加绑定和验证结构体
* @form :表单的映射字段名  binding:入参校验规则
 */

package service

import (
	"blog_service/internal/model"
	"blog_service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state:default=1" binding:"oneof=0 1"`
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

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svs *Service) CountTag(param *CountTagRequest) (int, error) {
	return svs.dao.CountTag(param.Name, param.State)
}

func (svs *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svs.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svs *Service) CreateTag(param *CreateTagRequest) error {
	return svs.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svs *Service) UpdateTag(param *UpdateTagRequest) error {
	return svs.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svs *Service) DeleteTag(param *DeleteTagRequest) error {
	return svs.dao.DeleteTag(param.ID)
}
