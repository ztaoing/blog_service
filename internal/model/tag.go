/**
* @Author:zhoutao
* @Date:2020/7/29 下午9:58
* @Desc:对标签模块的模型操作进行封装，并且只与实体产生关系
 */

package model

import (
	"blog_service/pkg/app"
	"github.com/jinzhu/gorm"
)

//标签
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name=? ", t.Name)
	}
	db = db.Where("state= ?", t.State)
	//统计
	if err := db.Model(&t).Where("is_del=?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name= ?", t.Name)
	}
	db = db.Where("state = ?", t.Name)
	//查找
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	//更新
	return db.Model(&t).Where("id = ? AND is_del= ?", t.ID, 0).Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	//删除
	return db.Where("id= ? AND is_del= ?", t.ID, 0).Delete(&t).Error
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ? ", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}
	return tag, nil
}
