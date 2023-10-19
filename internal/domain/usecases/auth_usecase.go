package usecases

import (
	"github.com/prcryx/raft-server/internal/domain/entities"
	"github.com/prcryx/raft-server/internal/domain/repositories"
)

type IAuthUseCase interface {
	SignUpWithEmailAndPassword(string, string) (*entities.UserEntity, error)
}

type AuthUseCase struct {
	repo repositories.AuthRepository
}

func NewAuthUseCase(repo repositories.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

func (usecase *AuthUseCase) SignUpWithEmailAndPassword(email, password string) (*entities.UserEntity, error) {
	userRecord, err := usecase.repo.SignUpWithEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}
	return &entities.UserEntity{
		UID:         userRecord.UID,
		DisplayName: userRecord.DisplayName,
		Email:       userRecord.Email,
		PhotoUrl:    userRecord.PhotoURL,
	}, nil
}
