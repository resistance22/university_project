package utils

type HttpResponse struct {
	Status   int
	Response any
}

func NewHttpResponse(response any, Status int) *HttpResponse {
	return &HttpResponse{
		Response: response,
		Status:   Status,
	}
}
