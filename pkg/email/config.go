package email

import "os"

type Conf struct {
	SMTPHost    string
	SMTPPort    string
	Username    string
	Password    string
	FromAddress string
}

func FromEnv() Conf {
	return Conf{
		SMTPHost:    os.Getenv("EMAIL_SMTP_HOST"),
		SMTPPort:    os.Getenv("EMAIL_SMTP_PORT"),
		Username:    os.Getenv("EMAIL_USERNAME"),
		Password:    os.Getenv("EMAIL_PASSWORD"),
		FromAddress: os.Getenv("EMAIL_FROM_ADDRESS"),
	}
}
