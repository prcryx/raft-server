package err

import (
	"errors"
	"fmt"
)

func UnexpectedException(code DebugErrorCode) error {
	return fmt.Errorf("%v code: %v", unexpectedError, code)
}

func OtpServiceFailedException() error {
	return errors.New(otpServiceFailed)
}
func OtpVerificationFailedException() error {
	return errors.New(otpVerificationFailed)
}

func UnauthorizedException() error {
	return errors.New(unauthorized)
}

func UserCreationFailedException() error {
	return errors.New(failedToCreateUser)
}

func CourruptedUserDataException() error {
	return errors.New(courruptedUserData)
}
func InvalidBodyRequestException() error {
	return errors.New(invalidBodyRequest)
}
