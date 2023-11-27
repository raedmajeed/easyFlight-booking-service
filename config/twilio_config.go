package config

import (
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioVerify struct {
	Client *twilio.RestClient
	Cfg    *ConfigParams
}

func SetupTwilio(cfg *ConfigParams) *TwilioVerify {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.SID,
		Password: cfg.TOKEN,
	})
	return &TwilioVerify{
		Client: client,
		Cfg:    cfg,
	}
}

func (t *TwilioVerify) SentTwilioOTP(phone string) (*verify.VerifyV2Verification, error) {
	params := verify.CreateVerificationParams{}
	params.SetTo("+91" + phone)
	params.SetChannel("sms")

	resp, err := t.Client.VerifyV2.CreateVerification(t.Cfg.SERVICETOKEN, &params)
	return resp, err
}

func (t *TwilioVerify) VerifyTwilioOTP(phone, code string) (*verify.VerifyV2VerificationCheck, error) {
	params := verify.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)

	resp, err := t.Client.VerifyV2.CreateVerificationCheck(t.Cfg.SERVICETOKEN, &params)
	return resp, err
}
