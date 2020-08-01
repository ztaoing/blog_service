/**
* @Author:zhoutao
* @Date:2020/8/1 下午3:28
* @Desc:
 */

package app

import (
	"blog_service/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJwtSecret() string {
	return global.JwtSetting.Secret
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	exipireTime := nowTime.Add(global.JwtSetting.Expire)
	claims := Claims{
		AppKey:    appKey,
		AppSecret: appSecret,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exipireTime.Unix(),
			Issuer:    global.JwtSetting.Issuer,
		},
	}
	//根据claims和加密算法获得tokenClaims
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	//根据secret生成签名字符串
	token, err := tokenClaims.SignedString(GetJwtSecret())
	return token, err
}

//解析和校验token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})
	if tokenClaims != nil {
		//valid 验证基于时间的声明，如过期时间、签发者，生效时间
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
