package err

// error codes for low-level exceptions
// signature: 55XY
// descp: 4 <= X <= 9 and Y = m%2 == 1

type LogErrorCode int

const (
	FailedToMarshal   LogErrorCode = 5541
	FailedToUnmarshal LogErrorCode = 5543
	NotMatched        LogErrorCode = 5545
)
