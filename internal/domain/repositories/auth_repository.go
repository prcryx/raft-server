package repositories

import (
	"firebase.google.com/go/auth"
)

type AuthRepository interface {
	SignUpWithEmailAndPassword(string, string) (*auth.UserRecord, error)
}
