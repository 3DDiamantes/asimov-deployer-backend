package apierror

type ApiError struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func New(s int, m string) *ApiError {
	return &ApiError{
		Status:  s,
		Message: m,
	}
}