package mail

import "testing"

func TestSendLogMail(t *testing.T) {
	loggerMail.SendMail("fang.changnian@qq.com", "hello", "<h1>golang</h1>")
}
