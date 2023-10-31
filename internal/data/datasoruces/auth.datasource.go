package datasoruces

import (
	"errors"
	"fmt"
	"log"

	"github.com/prcryx/raft-server/internal/domain/entities"
	"github.com/prcryx/raft-server/internal/domain/types"
	"github.com/prcryx/raft-server/internal/infrastructure/jwt"
	"github.com/prcryx/raft-server/internal/infrastructure/twilio"
	"gorm.io/gorm"
)

type IAuthDataSource interface {
	SendOtp(types.OtpReqBody) (*types.OtpResBody, error)
	Login(types.OtpVerificationReqBody) (*entities.User, error)
}

type AuthDataSource struct {
	db          *gorm.DB
	twilioApp   *twilio.TwilioApp
	jwtStrategy *jwt.JwtStrategy
}

var _ IAuthDataSource = (*AuthDataSource)(nil)

func NewAuthDataSource(db *gorm.DB, app *twilio.TwilioApp, jwtStrategy *jwt.JwtStrategy) (*AuthDataSource, error) {
	return &AuthDataSource{
		db:          db,
		twilioApp:   app,
		jwtStrategy: jwtStrategy,
	}, nil
}

// send otp

func (authDataSoruce *AuthDataSource) SendOtp(otpReq types.OtpReqBody) (*types.OtpResBody, error) {
	return authDataSoruce.twilioApp.SendOtp(otpReq)
}

// verify and then find or create user

func (authDataSoruce *AuthDataSource) Login(otpVerificationReq types.OtpVerificationReqBody) (*entities.User, error) {
	userEntity := new(entities.UserEntity)
	otpVerificationRes, otpVerificationErr := authDataSoruce.twilioApp.VerifyOtp(otpVerificationReq)
	if otpVerificationErr != nil {
		return nil, otpVerificationErr
	}
	if otpVerificationRes.VerificationStatus == types.Approved {
		//find or create user with the given phoneNo from otp verification response
		var err error
		userEntity, err = authDataSoruce.findOrCreateUser(otpVerificationRes.PhoneNo)
		if err != nil {
			return nil, err
		}
		
		// we need to generate the access and refresh token
		user := new(entities.User)
		accessToken := authDataSoruce.jwtStrategy.GenerateToken(userEntity.UID, userEntity.PhoneNo)
		fmt.Println("AccessToken:", accessToken)
		user = userEntity.ToUser(accessToken)
		fmt.Println(user)
		return user, nil
		// return nil, errors.New("user does not exists")
	}
	// if verification canceled
	return nil, errors.New("login failed")
}

//get exisiting user or create new user

func (authDataSource *AuthDataSource) findOrCreateUser(phoneNo string) (*entities.UserEntity, error) {
	// Check if the user already exists in the database
	var user entities.UserEntity
	err := authDataSource.db.Where(&entities.UserEntity{PhoneNo: phoneNo}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// user does not exist
			log.Printf("%v does not exists", phoneNo)
			user = entities.UserEntity{
				PhoneNo: phoneNo,
			}
			if err := authDataSource.db.Create(&user).Error; err != nil {
				// An error occurred while creating the user
				log.Printf("error occurred while creating user with : %v ", phoneNo)

				return nil, fmt.Errorf("error occurred while creating user with : %v ", phoneNo)
			}
			return &user, nil
		} else {
			// errors other than not found error
			return nil, err
		}
	}

	return &user, nil

}
