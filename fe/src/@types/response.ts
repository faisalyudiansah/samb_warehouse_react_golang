export interface ResponseMessageOnly {
    message: string;
}

export interface ResponseApiError {
    field: string;
    message: string;
}

export interface ResponseReportResult {
    warehouse: string;
    product: string;
    qty_dus: number;
    qty_pcs: number;
    type: string;
}

export interface ResponseWarehouse {
    whsPK: number;
    whsName: string;
}

export interface ResponseProduct {
    productPK: number;
    productName: string;
}

export interface ResponseSupplier {
    supplierPK: number;
    supplierName: string;
}