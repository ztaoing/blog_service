/**
* @Author:zhoutao
* @Date:2020/7/31 上午10:06
 */

package service

import (
	"blog_service/internal/model"
	"blog_service/pkg/app"
)

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

//validator
type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string `form:"create_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" binding:"gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string `form:"create_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//以文章信息为主题，以标签内容为附属
type Article struct {
	ID             uint32     `json:"id"`
	Title          string     `json:"title"`
	Desc           string     `json:"desc"`
	Content        string     `json:"content"`
	ConverImageUrl string     `json:"conver_image_url"`
	State          uint8      `json:"state"`
	Tag            *model.Tag `json:"tag"`
}

func (svs *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := svs.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}
	articleTag, err := svs.dao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}
	tag, err := svs.dao.GetTag(articleTag.ID, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}
	return &Article{
		ID:             article.ID,
		Title:          article.Title,
		Desc:           article.Desc,
		Content:        article.Content,
		ConverImageUrl: article.CoverImageUrl,
		State:          article.State,
		Tag:            &tag,
	}, nil
}

func (svs *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {

}

func (svs *Service) CreateArticle(param *CreateArticleRequest) error {

}

func (svs *Service) UpdateArticle(param *UpdateArticleRequest) error {

}

func (svs *Service) DeleteArticle(param *DeleteArticleRequest) error {

}
