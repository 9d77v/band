package alisms

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/9d77v/band/pkg/sms"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SMS struct {
	client *dysmsapi20170525.Client
	conf   sms.Conf
}

func NewSMS(conf sms.Conf) (*SMS, error) {
	config := &openapi.Config{
		AccessKeyId:     &conf.AccessKey,
		AccessKeySecret: &conf.SecretKey,
	}
	config.Endpoint = new("dysmsapi.aliyuncs.com")
	client, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	return &SMS{
		client: client,
		conf:   conf,
	}, err
}

func (s *SMS) SendSms(phoneNumbers, captcha string) (smsResp map[string]any, err error) {
	templateParam := fmt.Sprintf("{\"code\":\"%s\"}", captcha)
	templateCode := s.conf.ForeignTemplateCode
	if strings.HasPrefix(phoneNumbers, "+86") {
		templateCode = s.conf.TemplateCode
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      &s.conf.SignName,
		TemplateCode:  &templateCode,
		PhoneNumbers:  &phoneNumbers,
		TemplateParam: &templateParam,
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, err := s.client.SendSmsWithOptions(sendSmsRequest, runtime)
		data, _ := json.Marshal(&resp)
		json.Unmarshal(data, &smsResp)
		return err
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = new(tryErr.Error())
		}
		_, err := util.AssertAsString(error.Message)
		if err != nil {
			return nil, err
		}
	}
	return smsResp, tryErr
}
