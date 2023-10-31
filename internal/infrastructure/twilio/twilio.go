package twilio

import (
	"encoding/json"
	"fmt"

	"github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/internal/domain/types"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type twilioAppParams struct {
	accountSid string
	authToken  string
	verifySid  string
}

type ITwilioApp interface {
	SendOtp(types.OtpReqBody) (*types.OtpResBody, error)
	VerifyOtp(types.OtpVerificationReqBody) (*types.OtpVerificationResBody, error)
}

type TwilioApp struct {
	client *twilio.RestClient
	params twilioAppParams
}

func NewTwilioApp(conf *config.EnvConfig) *TwilioApp {
	appParams := twilioAppParams{
		accountSid: conf.Twilio.AccountSid,
		authToken:  conf.Twilio.AuthToken,
		verifySid:  conf.Twilio.VerifySid,
	}
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: appParams.accountSid,
		Password: appParams.authToken,
	})

	// create twilio app

	return &TwilioApp{
		client: client,
		params: appParams,
	}
}

var _ ITwilioApp = (*TwilioApp)(nil)

func (app *TwilioApp) SendOtp(otpReq types.OtpReqBody) (*types.OtpResBody, error) {

	var (
		verifySid  string = app.params.verifySid
		receiver   string = fmt.Sprintf("%v%v", otpReq.CountryCode, otpReq.PhoneNo)
		resultData map[string]interface{}
	)

	params := &verify.CreateVerificationParams{}
	params.SetTo(receiver)
	params.SetChannel("sms")

	resp, err := app.client.VerifyV2.CreateVerification(verifySid, params)
	if err != nil {
		return nil, err
	}

	response, resErr := json.Marshal(*resp)
	if resErr != nil {
		return nil, resErr
	}

	resultErr := json.Unmarshal(response, &resultData)
	if resultErr != nil {
		return nil, resultErr
	}

	return &types.OtpResBody{
		VerificationStatus: types.Status(fmt.Sprintf("%v", resultData["status"])),
		PhoneNo:            fmt.Sprintf("%v", resultData["to"]),
	}, nil
}

func (app *TwilioApp) VerifyOtp(otpReqBody types.OtpVerificationReqBody) (*types.OtpVerificationResBody, error) {
	var (
		verificationSid string = app.params.verifySid
		otp             string = otpReqBody.Otp
		to              string = fmt.Sprintf("%v", otpReqBody.PhoneNo)
		resultData      map[string]interface{}
	)
	params := &verify.CreateVerificationCheckParams{}
	params.SetCode(otp)
	params.SetTo(to)

	resp, err := app.client.VerifyV2.CreateVerificationCheck(verificationSid, params)
	if err != nil {
		return nil, err
	}
	response, resErr := json.Marshal(*resp)
	if resErr != nil {
		return nil, resErr
	}

	resultErr := json.Unmarshal(response, &resultData)
	if resultErr != nil {
		return nil, resultErr
	}

	return &types.OtpVerificationResBody{
		PhoneNo:            fmt.Sprintf("%v", resultData["to"]),
		VerificationStatus: types.Status(fmt.Sprintf("%v", resultData["status"])),
	}, nil
}
