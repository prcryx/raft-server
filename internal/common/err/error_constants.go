package err

const (
	unexpectedError       string = "unexpected error occurred"
	internalServerError   string = "internal server error"
	invalidBodyRequest    string = "invalid request body"
	otpServiceFailed      string = "otp service failed"
	otpVerificationFailed string = "otp verification failed"
	failedToCreateUser    string = "failed to create user"
	unauthorized          string = "unauthorized"
	forbidden             string = "forbidden"
	signingError             string = "signing error"
	courruptedUserData    string = "corrupted user data"
)
