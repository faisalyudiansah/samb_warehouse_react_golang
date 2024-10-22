package dtos

type ResponseMessageOnly struct {
	Message string `json:"message"`
}

type ResponseApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

type ResponseReportResult struct {
	Warehouse string `json:"warehouse"`
	Product   string `json:"product"`
	QtyDus    int    `json:"qty_dus"`
	QtyPcs    int    `json:"qty_pcs"`
	Type      string `json:"type"`
}
