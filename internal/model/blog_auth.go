/**
* @Author:zhoutao
* @Date:2020/8/1 下午3:20
* @Desc:
 */

package model

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}
