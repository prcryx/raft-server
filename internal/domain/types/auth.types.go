package types

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Canceled Status = "canceled"
)

type OtpReqBody struct {
	CountryCode string `json:"countryCode"`
	PhoneNo     string `json:"phoneNo"`
}

type OtpResBody struct {
	PhoneNo            string `json:"phoneNo"`
	VerificationStatus Status `json:"verificationStatus"`
}

type OtpVerificationReqBody struct {
	PhoneNo string `json:"phoneNo"`
	Otp     string `json:"otp"`
}
type OtpVerificationResBody struct {
	PhoneNo            string `json:"phoneNo"`
	CreatedAt          string `json:"date_created"`
	VerificationStatus Status `json:"status"`
}

type AuthenticationResponse struct {
	UID         uint   `json:"uid"`
	AccessToken string `json:"accessToken"`
	PhoneNo     string `json:"phoneNo"`
	CreatedAt   int64  `json:"createdAt"`
}

type TokenRefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}
