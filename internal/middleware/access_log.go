/**
* @Author:zhoutao
* @Date:2020/8/1 下午5:14
* @Desc:
 */

package middleware

import (
	"blog_service/global"
	"blog_service/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		//将其赋予当前的writer写入流
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request": c.Request.PostForm.Encode(),
			"reponse": bodyWriter.body.String(),
		}
		//记录日志
		global.Logger.WithFields(fields).Infof("access log:method:%s ,status_code:%d,begin_time:%d,end_time:%d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
