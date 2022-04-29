package adapter

type BadRequestError struct {
	Msg string
}

func NewBadRequestError(msg string) *BadRequestError {
	return &BadRequestError{
		Msg: msg,
	}
}

func (e *BadRequestError) Error() string {
	return e.Msg
}

type MethodNotAllowedError struct{}

func NewMethodNotAllowedError() *MethodNotAllowedError {
	return &MethodNotAllowedError{}
}

func (e *MethodNotAllowedError) Error() string {
	return "method not allowed"
}
