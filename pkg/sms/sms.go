package sms

type SMS interface {
	SendSms(phoneNumbers, captcha string) (smsResp map[string]interface{}, err error)
}
