package dtos

type ResponseMessageOnly struct {
	Message string `json:"message"`
}

type ResponseApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}
