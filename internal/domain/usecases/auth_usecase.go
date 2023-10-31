package usecases

import (
	"github.com/prcryx/raft-server/internal/domain/entities"
	"github.com/prcryx/raft-server/internal/domain/repositories"
	"github.com/prcryx/raft-server/internal/domain/types"
)

type IAuthUseCase interface {
	SendOtp(types.OtpReqBody) (*types.OtpResBody, error)
	Login(types.OtpVerificationReqBody) (*entities.User, error)
}

type AuthUseCase struct {
	repo repositories.AuthRepository
}

var _ IAuthUseCase = (*AuthUseCase)(nil)

func NewAuthUseCase(repo repositories.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

func (usecase *AuthUseCase) SendOtp(otpReq types.OtpReqBody) (*types.OtpResBody, error) {
	otpRes, err := usecase.repo.SendOtp(otpReq)
	if err != nil {
		return nil, err
	}
	return otpRes, nil
}

func (usecase *AuthUseCase) Login(otpReq types.OtpVerificationReqBody) (*entities.User, error) {
	user, err := usecase.repo.Login(otpReq)
	if err != nil {
		return nil, err
	}
	return user, nil
}
