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

type ResponseWarehouse struct {
	WhsPK   int64  `json:"whs_pK"`
	WhsName string `json:"whs_name"`
}

type ResponseProduct struct {
	ProductPK   int64  `json:"product_pK"`
	ProductName string `json:"product_name"`
}

type ResponseSupplier struct {
	SupplierPK   int64  `json:"supplier_pK"`
	SupplierName string `json:"supplier_name"`
}
