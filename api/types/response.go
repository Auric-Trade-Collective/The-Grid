package types

type ApiResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    []any  `json:"data"`
}

func NewApiResponse(message string, data []any) ApiResponse {
	return ApiResponse{Error: false, Message: message, Data: data}
}

func NewApiError(message string) ApiResponse {
	return ApiResponse{Error: true, Message: message, Data: []any{}}
}
