package repositories

import (
	"github.com/prcryx/raft-server/internal/domain/entities"
	"github.com/prcryx/raft-server/internal/domain/types"
)

type AuthRepository interface {
	SendOtp(types.OtpReqBody) (*types.OtpResBody, error)
	Login(types.OtpVerificationReqBody) (*entities.User, error)
}
