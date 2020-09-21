package di

import (
	"goweb/internal/conf"
	"goweb/pkg/mail"
	"testing"
)

func TestMail(t *testing.T) {
	mail.SendMail("demo@qq.com", "golang mail", "<h1>hello</h1> best wishes")
	t.Log(conf.Conf.Mail)
}
