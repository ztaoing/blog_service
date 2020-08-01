/**
* @Author:zhoutao
* @Date:2020/7/31 下午11:15
* @Desc:
 */

package dao

import "blog_service/internal/model"

func (d *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.GetByAID(d.engine)
}

func (d *Dao) GetArticleTagListByID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleID: tagID}
	return articleTag.ListByTID(d.engine)
}

func (d *Dao) GetArticleTagListByAIDS(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByAIDs(d.engine, articleIDs)
}

func (d *Dao) CreateArticleTag(articleID, tagID uint32, createBy string) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
		TagID:     tagID,
		Model: &model.Model{
			CreatedBy: createBy,
		},
	}
	return articleTag.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(articleID, tagId uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
	}
	values := map[string]interface{}{
		"article_id":  articleID,
		"tag_id":      tagId,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
	}
	return articleTag.DeleteOne(d.engine)
}
