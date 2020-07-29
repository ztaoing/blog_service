/**
* @Author:zhoutao
* @Date:2020/7/29 下午9:58
 */

package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
