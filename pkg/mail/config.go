package mail

import (
	"gopkg.in/gomail.v2"
)

type Config struct {
	Key       string
	UserName  string
	UserEmail string
	Password  string
	Host      string
	Port      int
}

var _ Mail = (*Config)(nil)
var configs []*Config

func (c *Config) SendMail(to, subject, body string, attachs ...*Attach) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", msg.FormatAddress(c.UserEmail, c.UserName))
	msg.SetHeader("To", to)           //发送给多个用户
	msg.SetHeader("Subject", subject) //设置邮件主题
	msg.SetBody("text/html", body)    //设置邮件正文
	// 附件
	for _, attach := range attachs {
		msg.Attach(attach.Path, gomail.Rename(attach.Name))
	}
	d := gomail.NewDialer(c.Host, c.Port, c.UserEmail, c.Password)
	err := d.DialAndSend(msg)
	return err
}

func Init(cs []*Config) {
	configs = cs
}
