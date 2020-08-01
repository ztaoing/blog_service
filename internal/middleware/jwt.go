/**
* @Author:zhoutao
* @Date:2020/8/1 下午4:38
* @Desc:
 */

package middleware

import (
	"blog_service/pkg/app"
	"blog_service/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := ctx.GetQuery("token"); exist {
			token = s
		} else {
			token = ctx.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}
		if ecode != errcode.Success {
			response := app.NewResponse(ctx)
			response.ToErrorResponse(ecode)
			return
		}
		ctx.Next()
	}
}
