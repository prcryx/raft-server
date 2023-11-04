package utils

func ArrayLengthValidator[T any](array []T, expectedLength int) bool {
	if len(array) == expectedLength {
		return true
	} else {
		return false
	}
}
