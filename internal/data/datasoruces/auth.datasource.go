package datasoruces

import (
	"fmt"

	"github.com/prcryx/raft-server/internal/domain/entities"
	"github.com/prcryx/raft-server/internal/domain/types"
	"github.com/prcryx/raft-server/internal/infrastructure/twilio"
	"gorm.io/gorm"
)

type IAuthDataSource interface {
	SendOtp(types.OtpReqBody) (*types.OtpResBody, error)
	Login(types.OtpVerificationReqBody) (*entities.UserEntity, error)
}

type AuthDataSource struct {
	db        *gorm.DB
	twilioApp *twilio.TwilioApp
}

var _ IAuthDataSource = (*AuthDataSource)(nil)

func NewAuthDataSource(db *gorm.DB, app *twilio.TwilioApp) (*AuthDataSource, error) {
	return &AuthDataSource{db: db, twilioApp: app}, nil
}

// send otp

func (authDataSoruce *AuthDataSource) SendOtp(otpReq types.OtpReqBody) (*types.OtpResBody, error) {
	return authDataSoruce.twilioApp.SendOtp(otpReq)
}

// verify and then find or create user

func (authDataSoruce *AuthDataSource) Login(otpVerificationReq types.OtpVerificationReqBody) (*entities.UserEntity, error) {
	user := new(entities.UserEntity)
	otpVerificationRes, otpVerificationErr := authDataSoruce.twilioApp.VerifyOtp(otpVerificationReq)
	if otpVerificationErr != nil {
		return nil, otpVerificationErr
	}
	if otpVerificationRes.VerificationStatus == types.Approved {
		//find or create user with the given phoneNo from otp verification response
		// authDataSoruce.db.findOrCreateUser(phoneNo)
		fmt.Println(otpVerificationRes.PhoneNo)
	}
	return user, nil
}
