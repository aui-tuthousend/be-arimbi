package utils

type APIResponse[T any] struct {
	Code   int    `json:"code"`
	Message string `json:"status"`
	Data   *T      `json:"data"`
}


func SuccessResponse[T any](data *T) APIResponse[T] {
	return APIResponse[T]{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
}

func ErrorResponse[T any](code int, message string) APIResponse[T] {
	return APIResponse[T]{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}