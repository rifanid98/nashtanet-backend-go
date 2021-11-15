package response

type Success struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func NewSuccess(data interface{}, status int) *Success {
	return &Success{
		StatusCode: status,
		Data:       data,
	}
}
