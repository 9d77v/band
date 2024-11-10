package alisms

import (
	"reflect"
	"testing"

	"github.com/9d77v/band/pkg/sms"
)

func TestSMS_SendSms(t *testing.T) {
	sms, _ := NewSMS(sms.Conf{
		Type:                "alisms",
		AccessKey:           "",
		SecretKey:           "",
		SignName:            "",
		TemplateCode:        "",
		ForeignTemplateCode: "",
	})
	type args struct {
		phoneNumbers string
		captcha      string
	}
	tests := []struct {
		name        string
		s           *SMS
		args        args
		wantSmsResp map[string]interface{}
		wantErr     bool
	}{
		{"", sms, args{"+86XXXXXXXXXX", "111121"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSmsResp, err := tt.s.SendSms(tt.args.phoneNumbers, tt.args.captcha)
			if (err != nil) != tt.wantErr {
				t.Errorf("SMS.SendSms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSmsResp, tt.wantSmsResp) {
				t.Errorf("SMS.SendSms() = %v, want %v", gotSmsResp, tt.wantSmsResp)
			}
		})
	}
}
