package jwt

import "testing"

func TestGenerateToken(t *testing.T) {
	var uid uint = 1
	phoneNo := "+919366123354"
	jwtStrategy := JwtStrategy{
		secretKey: []byte("b2dded49"),
	}

	tokenStr := jwtStrategy.GenerateToken(uid, phoneNo)
	if tokenStr == "" {
		t.Errorf("got tokenStr: %q, wanted not empty", tokenStr)
	}
}
