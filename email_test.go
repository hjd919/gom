package gom

import "testing"

func TestEmail_Send(t *testing.T) {
	config := &EmailConf{
		FromMail: "297538600@qq.com",
		Host:     "smtp.qq.com",
		Port:     465,
		FromName: "小抿一口",
		Username: "297538600@qq.com",
		Password: "--",
	}
	e := NewEmail(config)
	// toEmails := []ToEmail{}
	// toEmails = append(toEmails, ToEmail{Address: "hjdweapp@163.com", Name: "hjdweapp"})
	p := &EmailSendParam{
		ToEmails: []ToEmail{ToEmail{Address: "hjdweapp@163.com", Name: "hjdweapp"}},
		Subject:  "主题subject",
		Body:     "多个邮件人",
	}
	if err := e.Send(p); err != nil {
		t.Errorf("Email.Send() error = %v", err)
	}
}

func TestEmail_SendOne(t *testing.T) {
	config := &EmailConf{
		FromMail: "297538600@qq.com",
		Host:     "smtp.qq.com",
		Port:     465,
		FromName: "小抿一口",
		Username: "297538600@qq.com",
		Password: "--",
	}
	e := NewEmail(config)
	p := &EmailSendOneParam{
		ToEmail: ToEmail{Address: "hjdweapp@163.com", Name: "hjdweapp"},
		Subject: "主题subject",
		Body:    "单个邮件人",
	}
	if err := e.SendOne(p); err != nil {
		t.Errorf("Email.Send() error = %v", err)
	}
}
