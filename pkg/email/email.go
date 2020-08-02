/**
* @Author:zhoutao
* @Date:2020/8/1 下午5:58
* @Desc:邮件报警处理
 */

package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SEMTPInfo
}

type SEMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SEMTPInfo) *Email {
	return &Email{
		SEMTPInfo: info,
	}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject) //邮件主题
	m.SetBody("text/html", body)
	//连接SMPP服务器
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: e.IsSSL,
	}
	return dialer.DialAndSend(m)
}
