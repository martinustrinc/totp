package dtoModel

// CustomError is a custom error type that implements the error interface.
type CustomError struct {
	Message string
}

// Error returns the error message.
func (e *CustomError) Error() string {
	return e.Message
}
