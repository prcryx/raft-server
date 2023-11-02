package err

type AppError struct {
	code int
	msg  string
}

func newAppError(code int, msg string) AppError {
	return AppError{
		code: code,
		msg:  msg,
	}
}

var _ error = (*AppError)(nil)

func (appErr AppError) Error() string {
	return appErr.msg
}

func (appErr AppError) GetCode() int {
	return appErr.code
}
