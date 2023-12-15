package sms_factory

import (
	"sync"

	"github.com/9d77v/band/pkg/sms"
	"github.com/9d77v/band/pkg/sms/impl/alisms"
)

var (
	client sms.SMS
	once   sync.Once
)

var (
	TypeAlisms = "alisms"
)

func NewSMS(conf sms.Conf) (sms.SMS, error) {
	var err error
	var client sms.SMS
	switch conf.Type {
	default:
		client, err = alisms.NewSMS(conf)
	}
	return client, err
}

func SMSSingleton(conf sms.Conf) (sms.SMS, error) {
	var err error
	once.Do(func() {
		client, err = NewSMS(conf)
	})
	return client, err
}
