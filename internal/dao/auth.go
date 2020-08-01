/**
* @Author:zhoutao
* @Date:2020/8/1 下午3:50
* @Desc:token
 */

package dao

import "blog_service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
