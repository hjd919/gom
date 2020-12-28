package gom

import "testing"

func TestSms_Send(t *testing.T) {
	e := NewSms(&SmsConf{
		AccessKeyId:     "LTAI4FbitYh5ExGoED99UUNQ",
		AccessKeySecret: "--",
		SignName:        "师生汇",
		TemplateCode:    "SMS_200185878",
	})
	if err := e.Send(&SmsSendParam{
		Phones: []string{"18500223089"},
		TemplateParam: map[string]interface{}{
			"code": 12345,
		},
	}); err != nil {
		t.Errorf("Sms.Send() error = %v", err)
	}
}
