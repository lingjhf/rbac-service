package errors

type ErrorMessage struct {
	Code    string
	Message any
}

func NewError(code string, message any) *ErrorMessage {
	return &ErrorMessage{Code: code, Message: message}
}
