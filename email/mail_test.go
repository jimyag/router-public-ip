package email

import (
	"fmt"
	"testing"
)

func TestSendEmail(t *testing.T) {
	var to = []string{"1@jimyag.cn"}
	from := "1@jimyag.cn"
	nickname := "jimyag"
	secret := "你的密码"
	host := "smtp.mxhichina.com"
	//port := 25
	port := 465
	subject := "Perfect Vue Admin 发送邮件测试"
	body := "测试内容1"
	if err := SendEmail(from, to, secret, host, nickname, subject, body, port, true); err != nil {
		fmt.Println("发送失败: ", err)
	} else {
		fmt.Println("发送成功")
	}
}
