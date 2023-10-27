package types

type OtpRequest struct {
	CountryCode        string      `json:"countryCode"`
	PhoneNo            string      `json:"phoneNo"`
}

type AuthenticationResponse struct {
	Token string `json:"token"`
}

type OtpVerificationRequestBody struct {
	Token string
	Otp   string
}

type TokenRefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}
