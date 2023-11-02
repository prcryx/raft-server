package err

// error codes for low-level exceptions
// signature: 55XY
// descp: 4 <= X <= 9 and Y = m%2 == 1

type DebugErrorCode int

const (
	FailedToMarshal   DebugErrorCode = 5541
	FailedToUnmarshal DebugErrorCode = 5543
	NotMatched        DebugErrorCode = 5545
)
