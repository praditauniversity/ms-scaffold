package web

type successResponse struct {
	// Code   int         `json:"code"`
	// Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type errorResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors"`
}

func SuccessResponse(/*code int, status string, */data interface{}) successResponse {
	return successResponse{
		// Code:   code,
		// Status: status,
		Data:   data,
	}
}

func ErrorResponse(code int, status string, errors interface{}) errorResponse {
	return errorResponse{
		Code:   code,
		Status: status,
		Errors: errors,
	}
}