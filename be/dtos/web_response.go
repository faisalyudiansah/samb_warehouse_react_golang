package dtos

type WebResponse[T any] struct {
	Message string        `json:"message,omitempty"`
	Data    T             `json:"data,omitempty"`
	Paging  *PageMetaData `json:"paging,omitempty"`
	Errors  []FieldError  `json:"errors,omitempty"`
}

type WebResponseSeek[T any] struct {
	Message string            `json:"message,omitempty"`
	Data    T                 `json:"data"`
	Paging  *SeekPageMetaData `json:"paging,omitempty"`
	Errors  []FieldError      `json:"errors,omitempty"`
}

type WebResponseDataEmptyArray[T any] struct {
	Message string        `json:"message,omitempty"`
	Data    T             `json:"data"`
	Paging  *PageMetaData `json:"paging,omitempty"`
	Errors  []FieldError  `json:"errors,omitempty"`
}

type PageMetaData struct {
	Page      int64  `json:"page"`
	Size      int64  `json:"size"`
	TotalItem int64  `json:"total_item"`
	TotalPage int64  `json:"total_page"`
	Links     *Links `json:"links"`
}

type SeekPageMetaData struct {
	Size int64  `json:"size"`
	Last string `json:"last"`
	Next bool   `json:"next"`
}

type Links struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
