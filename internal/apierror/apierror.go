package apierror

import "fmt"

type ApiError struct {
	Status int
	Message string
}

func (ae *ApiError) Error() string {
	return fmt.Sprintf("Status: %d Message: %s", ae.Status, ae.Message)
}

func New(s int, m string) *ApiError {
	return &ApiError{
		Status:  s,
		Message: m,
	}
}