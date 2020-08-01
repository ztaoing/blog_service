/**
* @Author:zhoutao
* @Date:2020/8/1 下午4:17
* @Desc:
 */

package service

import "errors"

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svs *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svs.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist.")
}
