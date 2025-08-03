package utils

type ApiResponses struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	TotalRecords int `json:"total_records"`
	CurrentPage int `json:"current_page"`
	TotalPage int `json:"total_page"`
	PerPage int `json:"per_page"`
}
