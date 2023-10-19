package repository_impl

import (
	"firebase.google.com/go/auth"
	"github.com/prcryx/raft-server/internal/data/datasoruces"
)

type AuthRepositoryImpl struct {
    authDataSrc *datasoruces.AuthDataSource
}

func NewAuthRepositoryImpl(authDataSrc *datasoruces.AuthDataSource) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
        authDataSrc: authDataSrc,
    }
}

func (repoImpl *AuthRepositoryImpl) SignUpWithEmailAndPassword(email, password string) (*auth.UserRecord, error) {
	return repoImpl.authDataSrc.SignUpWithEmailAndPassword(email, password)
}
