/**
* @Author:zhoutao
* @Date:2020/8/1 下午3:45
* @Desc:token
 */

package model

import "github.com/jinzhu/gorm"

//获取客户端传入的app_key和app_secret后，根据传入的认证信息进行验证，以此判断是否真的存在这条数据
func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}
