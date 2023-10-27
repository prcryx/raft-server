package utils

import (
	"github.com/prcryx/datacheckr"
	"github.com/prcryx/raft-server/internal/domain/types"
)

func ValidatePhone(otpReq *types.OtpReqBody) bool {
	validator := datacheckr.NewValidatorInstance()

	//addvalidation rules for india only
	validator.AddValidationRules(
		datacheckr.MaxStrLenValidation(11),
	)

	return validator.Validate(otpReq.PhoneNo)
}
