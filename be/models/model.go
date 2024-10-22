package models

import "time"

type MasterSupplier struct {
	SupplierPK   int64
	SupplierName string
}

type MasterCustomer struct {
	CustomerPK   int64
	CustomerName string
}

type MasterProduct struct {
	ProductPK   int64
	ProductName string
}

type MasterWarehouse struct {
	WhsPK   int64
	WhsName string
}

type TransaksiPenerimaanBarangHeader struct {
	TrxInPK      int64
	TrxInNo      string
	WhsIdf       int64
	TrxInDate    time.Time
	TrxInSuppIdf int64
	TrxInNotes   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     *time.Time
}

type TransaksiPenerimaanBarangDetail struct {
	TrxInDPK         int64
	TrxInIDF         int64
	TrxInDProductIdf int64
	TrxInDQtyDus     int
	TrxInDQtyPcs     int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeleteAt         *time.Time
}

type TransaksiPengeluaranBarangHeader struct {
	TrxOutPK      int64
	TrxOutNo      string
	WhsIdf        int64
	TrxOutDate    time.Time
	TrxOutSuppIdf int64
	TrxOutNotes   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeleteAt      *time.Time
}

type TransaksiPengeluaranBarangDetail struct {
	TrxOutDPK         int64
	TrxOutIDF         int64
	TrxOutDProductIdf int64
	TrxOutDQtyDus     int
	TrxOutDQtyPcs     int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeleteAt          *time.Time
}
