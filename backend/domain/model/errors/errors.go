package errors

type ValidateError struct {
	Msg string
}

func NewValidateError(msg string) *ValidateError {
	return &ValidateError{
		Msg: msg,
	}
}

func (e *ValidateError) Error() string {
	return e.Msg
}

type UnauthorizedError struct {
	Msg string
}

func NewUnauthorizedError(msg string) *UnauthorizedError {
	return &UnauthorizedError{
		Msg: msg,
	}
}

func (e *UnauthorizedError) Error() string {
	return e.Msg
}

type InternalError struct {
	Msg string
}

func NewInternalError(msg string) *InternalError {
	return &InternalError{
		Msg: msg,
	}
}

func (e *InternalError) Error() string {
	return e.Msg
}
