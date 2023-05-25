package smtp

import (
	"net/smtp"
	"strings"
)

const (
	HOST     = "smtp.exmail.qq.com:25"
	USER     = "linglu0212@163.com"
	PASSWORD = "Jishu123456"
)

type Email struct {
	To       string `json:"to"`       // 多个用;分开
	Subject  string `json:"subject"`  // 主题
	Msg      string `json:"msg"`      // 消息
	MailType string `json:"mailtype"` // 类型
}

// SendEmail 发送邮件
func SendEmail(email *Email) (err error) {
	hp := strings.Split(HOST, ":")
	auth := smtp.PlainAuth("", USER, PASSWORD, hp[0])
	sendTo := strings.Split(email.To, ";")
	done := make(chan error, 1024)
	var contentType string
	if email.MailType == "html" {
		contentType = "Content-Type: text/" + email.MailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	go func() {
		defer close(done)
		for _, v := range sendTo {
			str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.Subject+"~"+contentType+"~~", "~", "\r\n", -1) + email.Msg
			err := smtp.SendMail(
				HOST,
				auth,
				USER,
				[]string{v},
				[]byte(str),
			)
			done <- err
		}
	}()
	for i := 0; i < len(sendTo); i++ {
		err = <-done
		//发送一个成功 就算成功
		if err == nil {
			break
		}
	}
	return err
}
