package mail

import "testing"

func TestSendMail(t *testing.T) {
	SendMail("fang.changnian@qq.com", "hello", "<h1>golang</h1>")
	SendMailWithTmpl("fang.changnian@qq.com", "emailcode", map[string]string{
		"code": "1234",
	})
}
