package repository_impl

import (
	"github.com/prcryx/raft-server/internal/data/datasoruces"
	"github.com/prcryx/raft-server/internal/domain/entities"
	"github.com/prcryx/raft-server/internal/domain/repositories"
	"github.com/prcryx/raft-server/internal/domain/types"
)

type AuthRepositoryImpl struct {
	authDataSrc *datasoruces.AuthDataSource
}

var _ repositories.AuthRepository = (*AuthRepositoryImpl)(nil)

func NewAuthRepositoryImpl(authDataSrc *datasoruces.AuthDataSource) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		authDataSrc: authDataSrc,
	}
}

func (repoImpl *AuthRepositoryImpl) SendOtp(otpReq types.OtpReqBody) (*types.OtpResBody, error) {
	return repoImpl.authDataSrc.SendOtp(otpReq)
}

func (repoImpl *AuthRepositoryImpl) Login(otpVerificationReq types.OtpVerificationReqBody) (*entities.UserEntity, error) {
	return repoImpl.authDataSrc.Login(otpVerificationReq)
}
