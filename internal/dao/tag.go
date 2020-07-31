/**
* @Author:zhoutao
* @Date:2020/7/31 下午5:07
 */

package dao

import (
	"blog_service/internal/model"
	"blog_service/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pagesize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pagesize)
	return tag.List(d.engine, pageOffset, pagesize)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id, ModifiedBy: modifiedBy,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	return tag.Delete(d.engine)
}
