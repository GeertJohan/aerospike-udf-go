package udf

type udfError struct {
	message string
	code    int
}

func (u *udfError) Error() string {
	return u.message
}

var (
	ErrInvalidRegistration  = &udfError{message: "invalid registration (not a func)", code: 10}
	ErrFunctionNotAvailable = &udfError{message: "requested function is not available", code: 20}
)
