package apierror

import "fmt"

type ApiError struct {
	status int `json:"status"`
	message string `json:"message"`
}

func (ae *ApiError) Error() string {
	return fmt.Sprintf("Status: %d Message: %s", ae.status, ae.message)
}

func (ae *ApiError) Status() int {
	return ae.status
}

func (ae *ApiError) Message() string {
	return ae.message
}

func New(s int, m string) *ApiError {
	return &ApiError{
		status:  s,
		message: m,
	}
}