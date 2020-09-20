/**
* @Author:zhoutao
* @Date:2020/8/1 下午5:39
* @Desc:异常捕获处理+发送异常邮件提醒
 */

package middleware

import (
	"blog_service/global"
	"blog_service/pkg/app"
	"blog_service/pkg/email"
	"blog_service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	Mailer := email.NewEmail(
		&email.SEMTPInfo{
			Host:     global.EmailSetting.Host,
			Port:     global.EmailSetting.Port,
			IsSSL:    global.EmailSetting.IsSSL,
			UserName: global.EmailSetting.UserName,
			Password: global.EmailSetting.Password,
			From:     global.EmailSetting.From,
		},
	)
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//设置当前整个栈的调用信息
				global.Logger.WithCallerFrames().Errorf(ctx, "panic recover err:%v", err)
				//发送邮件提醒
				//邮件模可以定制
				err = Mailer.SendMail(global.EmailSetting.To,
					fmt.Sprintf("异常，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
				)
				if err != nil {
					global.Logger.Errorf(ctx, " Mailer.SendMail err:%v", err)
				}
				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort()
			}
		}()
		// It executes the pending handlers in the chain inside the calling handler
		ctx.Next()
	}
}
