/**
* @Author:zhoutao
* @Date:2020/7/29 下午9:58
 */

package model

import "blog_service/pkg/app"

//标签
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
