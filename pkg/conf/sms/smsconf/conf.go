package smsconf

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Type                string `yaml:"type"`
	AccessKey           string `yaml:"access_key"`
	SecretKey           string `yaml:"secret_key"`
	SignName            string `yaml:"sign_name"`
	TemplateCode        string `yaml:"template_code"`
	ForeignTemplateCode string `yaml:"foreign_template_code"`
}

func FromEnv() Conf {
	return Conf{
		Type:                env.String("SMS_TYPE"),
		AccessKey:           env.String("SMS_ACCESS_KEY"),
		SecretKey:           env.String("SMS_SECRET_KEY"),
		SignName:            env.String("SMS_SIGN_NAME"),
		TemplateCode:        env.String("SMS_TEMPLAET_CODE"),
		ForeignTemplateCode: env.String("SMS_FOREIGN_TEMPLAET_CODE"),
	}
}
