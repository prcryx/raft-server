package repository_impl

import (
	"firebase.google.com/go/auth"
	"github.com/prcryx/raft-server/internal/data/datasoruces"
	"github.com/prcryx/raft-server/internal/domain/repositories"
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

func (repoImpl *AuthRepositoryImpl) SignUpWithEmailAndPassword(email, password string) (*auth.UserRecord, error) {
	return repoImpl.authDataSrc.SignUpWithEmailAndPassword(email, password)
}
