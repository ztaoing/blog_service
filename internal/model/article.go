/**
* @Author:zhoutao
* @Date:2020/7/29 下午10:01
 */

package model

import "blog_service/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	state         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type Aritcle struct {
	List  []*Article
	Paper *app.Pager
}
