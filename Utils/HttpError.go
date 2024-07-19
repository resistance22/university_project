package utils

type HttpError struct {
	Message     string
	Status      int
	Description string
}

func NewHttpError(Message string, Status int, Description string) *HttpError {
	return &HttpError{
		Message:     Message,
		Status:      Status,
		Description: Description,
	}
}
