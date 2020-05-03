package services

func failResponse() *Response {
	return &Response{
		Success: false,
		Message: "invalid response",
	}
}

func successResponse() *Response {
	return &Response{
		Success: true,
		Message: "success response",
	}
}
