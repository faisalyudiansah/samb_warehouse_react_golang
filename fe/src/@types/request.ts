export interface PenerimaanBarang {
    whs_idf: number;
    trx_in_date: string;
    trx_in_supp_idf: number;
    trx_in_notes: string;
    trx_in_dproduct_idf: number;
    trx_in_dqty_dus: number;
    trx_in_dqty_pcs: number;
}

export interface PengeluaranBarang {
    whs_idf: number;
    trx_out_date: string;
    trx_out_supp_idf: number;
    trx_out_notes: string;
    trx_out_dproduct_idf: number;
    trx_out_dqty_dus: number;
    trx_out_dqty_pcs: number;
}
