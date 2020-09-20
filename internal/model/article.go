/**
* @Author:zhoutao
* @Date:2020/7/29 下午10:01
 */

package model

import (
	"blog_service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type Aritcle struct {
	List  []*Article
	Paper *app.Pager
}

func (a *Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(a).Update(values).Where("id = ? AND is_del = ? ", a.ID, 0).Error; err != nil {
		return err
	}
	return nil
}

func (a *Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Model(&a).Where("id = ? AND state = ? AND is_del= ? ", a.ID, a.State, a.IsDel)
	//First查询单条记录
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

func (a *Article) Delete(db *gorm.DB) error {
	if err := db.Model(a).Where("id = ? AND is_del = ?", a.ID, 0).Delete(a).Error; err != nil {
		return err
	}
	return nil
}

//使用关联查询 tag和article
type ArticleRow struct {
	ArticleID     uint32
	TagID         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

func (a *Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title", "ar.desc AS article_desc",
		"ar_cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+" AS at ").Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t"+
		" ON at.tag_id = t.id").Joins(" LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").Where(
		"at.`tag_id`= ? AND ar.state = ? AND ar.is_del = ? ", tagID, a.State, 0).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}

//文章列表总数
func (a *Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id ").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil

}
