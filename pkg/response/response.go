package response

type Response struct {
	Success      bool   `json:"success"`
	ResponseCode int    `json:"response_code"`
	Message      string `json:"message"`
	Data         any    `json:"data,omitempty"`
	Error        any    `json:"error,omitempty"`
	Meta         *Meta  `json:"meta,omitempty"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

// NewSuccessResponse creates a standardized success response
func ApiSuccessResponse(responseCode int, message string, data any) Response {
	return Response{
		Success:      true,
		ResponseCode: responseCode,
		Message:      message,
		Data:         data,
	}
}

// ApiErrorResponse creates a standardized error response
func ApiErrorResponse(responseCode int, message string, err any) Response {
	// If it's a standard error, convert to string or use a specific format
	var errPayload any
	if e, ok := err.(error); ok {
		errPayload = e.Error()
	} else {
		errPayload = err
	}

	return Response{
		Success:      false,
		ResponseCode: responseCode,
		Message:      message,
		Error:        errPayload,
	}
}
