package dtos

type RequestLoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type PenerimaanBarang struct {
	WhsIdf           int64  `json:"whs_idf" binding:"required,numeric,gt=0"`
	TrxInDate        string `json:"trx_in_date" binding:"required,date"`
	TrxInSuppIdf     int64  `json:"trx_in_supp_idf" binding:"required,numeric,gt=0"`
	TrxInNotes       string `json:"trx_in_notes" binding:"required"`
	TrxInDProductIdf int64  `json:"trx_in_dproduct_idf" binding:"required,numeric,gt=0"`
	TrxInDQtyDus     int    `json:"trx_in_dqty_dus" binding:"required,numeric,gt=0"`
	TrxInDQtyPcs     int    `json:"trx_in_dqty_pcs" binding:"required,numeric,gt=0"`
}
