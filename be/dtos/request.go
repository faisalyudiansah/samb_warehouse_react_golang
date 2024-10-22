package dtos

type RequestLoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type PenerimaanBarang struct {
	WhsIdf           int    `json:"whs_idf" binding:"required"`
	TrxInDate        string `json:"trx_in_date" binding:"required"`
	TrxInSuppIdf     int    `json:"trx_in_supp_idf" binding:"required"`
	TrxInNotes       string `json:"trx_in_notes" binding:"required"`
	TrxInDProductIdf string `json:"trx_in_dproduct_idf" binding:"required"`
	TrxInDQtyDus     string `json:"trx_in_dqty_dus" binding:"required"`
	TrxInDQtyPcs     string `json:"trx_in_dqty_pcs" binding:"required"`
}
