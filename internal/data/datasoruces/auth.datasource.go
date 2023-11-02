package datasoruces

import (
	"errors"
	"fmt"
	"log"

	e "github.com/prcryx/raft-server/internal/common/err"
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
	var receiver string = fmt.Sprintf("%v%v", otpReq.CountryCode, otpReq.PhoneNo)
	return authDataSoruce.twilioApp.SendOtp(receiver)
}

// verify and then find or create user

func (authDataSoruce *AuthDataSource) Login(otpVerificationReq types.OtpVerificationReqBody) (*entities.User, error) {
	userEntity := new(entities.UserEntity)

	otpVerificationRes, otpVerificationErr := authDataSoruce.twilioApp.VerifyOtp(otpVerificationReq.Otp, otpVerificationReq.PhoneNo)
	if otpVerificationErr != nil {
		return nil, otpVerificationErr
	}
	if otpVerificationRes.VerificationStatus != types.Approved {
		// if verification canceled
		return nil, e.UnauthorizedException()
	}
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

			//create user

			user = entities.UserEntity{
				PhoneNo: phoneNo,
			}
			if err := authDataSource.db.Create(&user).Error; err != nil {
				// An error occurred while creating the user
				log.Printf("\nuser creation failed %v\n errors: %v\n", phoneNo, err)
				return nil, e.UserCreationFailedException()
			}
			return &user, nil
		} else {
			// errors other than not found error
			log.Printf("\ncorrupted data for %v\n errors: %v\n", phoneNo, err)
			return nil, e.CourruptedUserDataException()
		}
	}

	return &user, nil

}
