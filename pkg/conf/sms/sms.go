package sms

import (
	"sync"

	"github.com/9d77v/band/pkg/conf/sms/impl/alisms"
	"github.com/9d77v/band/pkg/conf/sms/smsconf"
)

type SMS interface {
	SendSms(phoneNumbers, captcha string) (smsResp map[string]interface{}, err error)
}

var (
	client SMS
	once   sync.Once
)

var (
	TypeAlisms = "alisms"
)

func NewSMS(conf smsconf.Conf) (SMS, error) {
	var err error
	once.Do(func() {
		switch conf.Type {
		default:
			client, err = alisms.NewSMS(conf)
		}
	})
	return client, err
}
