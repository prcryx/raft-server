package err

import (
	"net/http"
)

func InternalServerError() AppError {
	return newAppError(http.StatusInternalServerError, internalServerError)
}

func UnauthorizedException() AppError {
	return newAppError(http.StatusUnauthorized, unauthorized)
}

func InvalidBodyRequestException() AppError {
	return newAppError(http.StatusBadRequest, invalidBodyRequest)
}

// func OtpServiceFailedException() error {
// 	return errors.New(otpServiceFailed)
// }
// func OtpVerificationFailedException() error {
// 	return errors.New(otpVerificationFailed)
// }

// func UserCreationFailedException() error {
// 	return errors.New(failedToCreateUser)
// }

// func CourruptedUserDataException() error {
// 	return errors.New(courruptedUserData)
// }
