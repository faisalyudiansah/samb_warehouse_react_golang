package models

import "time"

type MasterSupplier struct {
	SupplierPK   int
	SupplierName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     *time.Time
}

type MasterCustomer struct {
	CustomerPK   int
	CustomerName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     *time.Time
}

type MasterProduct struct {
	ProductPK   int
	ProductName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeleteAt    *time.Time
}

type MasterWarehouse struct {
	WhsPK     int
	WhsName   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}

type TransaksiPenerimaanBarangHeader struct {
	TrxInPK      int
	TrxInNo      string
	WhsIdf       int
	TrxInDate    time.Time
	TrxInSuppIdf int
	TrxInNotes   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     *time.Time
}

type TransaksiPenerimaanBarangDetail struct {
	TrxInDPK         int
	TrxInIDF         int
	TrxInDProductIdf int
	TrxInDQtyDus     int
	TrxInDQtyPcs     int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeleteAt         *time.Time
}

type TransaksiPengeluaranBarangHeader struct {
	TrxOutPK      int
	TrxOutNo      string
	WhsIdf        int
	TrxOutDate    time.Time
	TrxOutSuppIdf int
	TrxOutNotes   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeleteAt      *time.Time
}

type TransaksiPengeluaranBarangDetail struct {
	TrxOutDPK         int
	TrxOutIDF         int
	TrxOutDProductIdf int
	TrxOutDQtyDus     int
	TrxOutDQtyPcs     int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeleteAt          *time.Time
}
