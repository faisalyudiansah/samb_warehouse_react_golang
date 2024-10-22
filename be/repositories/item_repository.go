package repositories

import (
	"context"
	"database/sql"
	"server/dtos"
	helpercontext "server/helpers/helper_context"
	"server/models"
)

type ItemRepositoryInterface interface {
	FindWarehouseById(ctx context.Context, warehouseID int64) (*models.MasterWarehouse, error)
	FindSupplierById(ctx context.Context, supplierID int64) (*models.MasterSupplier, error)
	FindProductById(ctx context.Context, productID int64) (*models.MasterProduct, error)
	FindLastTrxInPKPenerimaan(ctx context.Context) (int64, error)
	FindLastTrxOutPKPengeluaran(ctx context.Context) (int64, error)
	CreatePenerimaanBarangHeader(ctx context.Context, reqBody *dtos.PenerimaanBarang, trxInNo string) (*int64, error)
	CreatePenerimaanBarangDetail(ctx context.Context, reqBody *dtos.PenerimaanBarang, TrxInIDF int64) error
	CreatePengeluaranBarangHeader(ctx context.Context, reqBody *dtos.PengeluaranBarang, trxOutNo string) (*int64, error)
	CreatePengeluaranBarangDetail(ctx context.Context, reqBody *dtos.PengeluaranBarang, TrxOutIDF int64) error
	GetReportData(ctx context.Context) ([]models.ReportResult, error)
}

type ItemRepositoryImplementation struct {
	db *sql.DB
}

func NewItemRepositoryImplementation(db *sql.DB) *ItemRepositoryImplementation {
	return &ItemRepositoryImplementation{
		db: db,
	}
}

func (c *ItemRepositoryImplementation) FindWarehouseById(ctx context.Context, warehouseID int64) (*models.MasterWarehouse, error) {
	query := "SELECT m.WhsPK, m.WhsName FROM masterwarehouse m WHERE m.whspk = $1 AND deleted_at IS NULL"
	tx := helpercontext.GetTx(ctx)
	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, warehouseID)
	} else {
		row = c.db.QueryRowContext(ctx, query, warehouseID)
	}
	var warehouse models.MasterWarehouse
	err := row.Scan(
		&warehouse.WhsPK,
		&warehouse.WhsName,
	)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (c *ItemRepositoryImplementation) FindSupplierById(ctx context.Context, supplierID int64) (*models.MasterSupplier, error) {
	query := "SELECT m.supplierpk, m.suppliername FROM mastersupplier m WHERE m.supplierpk = $1 AND deleted_at IS NULL"
	tx := helpercontext.GetTx(ctx)
	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, supplierID)
	} else {
		row = c.db.QueryRowContext(ctx, query, supplierID)
	}
	var supplier models.MasterSupplier
	err := row.Scan(
		&supplier.SupplierPK,
		&supplier.SupplierName,
	)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (c *ItemRepositoryImplementation) FindProductById(ctx context.Context, productID int64) (*models.MasterProduct, error) {
	query := "SELECT m.productpk, m.productname FROM masterproduct m WHERE m.productpk = $1 AND deleted_at IS NULL"
	tx := helpercontext.GetTx(ctx)
	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, productID)
	} else {
		row = c.db.QueryRowContext(ctx, query, productID)
	}
	var product models.MasterProduct
	err := row.Scan(
		&product.ProductPK,
		&product.ProductName,
	)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (c *ItemRepositoryImplementation) FindLastTrxInPKPenerimaan(ctx context.Context) (int64, error) {
	query := "SELECT TrxInPK FROM TransaksiPenerimaanBarangHeader ORDER BY created_at DESC LIMIT 1"
	tx := helpercontext.GetTx(ctx)
	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query)
	} else {
		row = c.db.QueryRowContext(ctx, query)
	}

	var lastTrxInPK int64
	err := row.Scan(&lastTrxInPK)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return lastTrxInPK, nil
}

func (c *ItemRepositoryImplementation) FindLastTrxOutPKPengeluaran(ctx context.Context) (int64, error) {
	query := "SELECT TrxOutPK FROM TransaksiPengeluaranBarangHeader ORDER BY created_at DESC LIMIT 1"
	tx := helpercontext.GetTx(ctx)
	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query)
	} else {
		row = c.db.QueryRowContext(ctx, query)
	}

	var lastTrxOutPK int64
	err := row.Scan(&lastTrxOutPK)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return lastTrxOutPK, nil
}

func (c *ItemRepositoryImplementation) CreatePenerimaanBarangHeader(ctx context.Context, reqBody *dtos.PenerimaanBarang, trxInNo string) (*int64, error) {
	query := `
        INSERT INTO TransaksiPenerimaanBarangHeader (TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING TrxInPK
    `
	tx := helpercontext.GetTx(ctx)
	var err error
	var trxInHeader models.TransaksiPenerimaanBarangHeader
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, trxInNo, reqBody.WhsIdf, reqBody.TrxInDate, reqBody.TrxInSuppIdf, reqBody.TrxInNotes).Scan(&trxInHeader.TrxInPK)
	} else {
		err = c.db.QueryRowContext(ctx, query, trxInNo, reqBody.WhsIdf, reqBody.TrxInDate, reqBody.TrxInSuppIdf, reqBody.TrxInNotes).Scan(&trxInHeader.TrxInPK)
	}
	if err != nil {
		return nil, err
	}
	return &trxInHeader.TrxInPK, nil
}

func (c *ItemRepositoryImplementation) CreatePenerimaanBarangDetail(ctx context.Context, reqBody *dtos.PenerimaanBarang, TrxInIDF int64) error {
	query := `
        INSERT INTO TransaksiPenerimaanBarangDetail (TrxInIDF, TrxInDProductIdf, TrxInDQtyDus, TrxInDQtyPcs, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
    `
	tx := helpercontext.GetTx(ctx)
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, TrxInIDF, reqBody.TrxInDProductIdf, reqBody.TrxInDQtyDus, reqBody.TrxInDQtyPcs)
	} else {
		_, err = c.db.ExecContext(ctx, query, TrxInIDF, reqBody.TrxInDProductIdf, reqBody.TrxInDQtyDus, reqBody.TrxInDQtyPcs)
	}
	return err
}

func (c *ItemRepositoryImplementation) CreatePengeluaranBarangHeader(ctx context.Context, reqBody *dtos.PengeluaranBarang, trxOutNo string) (*int64, error) {
	query := `
        INSERT INTO TransaksiPengeluaranBarangHeader (TrxOutNo, WhsIdf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING TrxOutPK
    `
	tx := helpercontext.GetTx(ctx)
	var err error
	var trxInHeader models.TransaksiPengeluaranBarangHeader
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, trxOutNo, reqBody.WhsIdf, reqBody.TrxOutDate, reqBody.TrxOutSuppIdf, reqBody.TrxOutNotes).Scan(&trxInHeader.TrxOutPK)
	} else {
		err = c.db.QueryRowContext(ctx, query, trxOutNo, reqBody.WhsIdf, reqBody.TrxOutDate, reqBody.TrxOutSuppIdf, reqBody.TrxOutNotes).Scan(&trxInHeader.TrxOutPK)
	}
	if err != nil {
		return nil, err
	}
	return &trxInHeader.TrxOutPK, nil
}

func (c *ItemRepositoryImplementation) CreatePengeluaranBarangDetail(ctx context.Context, reqBody *dtos.PengeluaranBarang, TrxOutIDF int64) error {
	query := `
        INSERT INTO TransaksiPengeluaranBarangDetail (TrxOutIDF, TrxOutDProductIdf, TrxOutDQtyDus, TrxOutDQtyPcs, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
    `
	tx := helpercontext.GetTx(ctx)
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, TrxOutIDF, reqBody.TrxOutDProductIdf, reqBody.TrxOutDQtyDus, reqBody.TrxOutDQtyPcs)
	} else {
		_, err = c.db.ExecContext(ctx, query, TrxOutIDF, reqBody.TrxOutDProductIdf, reqBody.TrxOutDQtyDus, reqBody.TrxOutDQtyPcs)
	}
	return err
}

func (c *ItemRepositoryImplementation) GetReportData(ctx context.Context) ([]models.ReportResult, error) {
	query := `
        SELECT
			wh.WhsName AS "warehouse",
			p.ProductName AS "product",
			COALESCE(SUM(tin_detail.TrxInDQtyDus), 0) AS "qty_dus",
			COALESCE(SUM(tin_detail.TrxInDQtyPcs), 0) AS "qty_pcs",
			'Penerimaan' AS "type"
		FROM
			TransaksiPenerimaanBarangHeader tin_header
		JOIN
			TransaksiPenerimaanBarangDetail tin_detail ON tin_detail.TrxInIDF = tin_header.TrxInPK
		JOIN
			MasterWarehouse wh ON tin_header.WhsIdf = wh.WhsPK
		JOIN
			MasterProduct p ON tin_detail.TrxInDProductIdf = p.ProductPK
		GROUP BY
			wh.WhsName,
			p.ProductName

		UNION ALL

		SELECT
			wh.WhsName AS "warehouse",
			p.ProductName AS "product",
			COALESCE(SUM(tout_detail.TrxOutDQtyDus), 0) AS "qty_dus",
			COALESCE(SUM(tout_detail.TrxOutDQtyPcs), 0) AS "qty_pcs",
			'Pengeluaran' AS "type"
		FROM
			TransaksiPengeluaranBarangHeader tout_header
		JOIN
			TransaksiPengeluaranBarangDetail tout_detail ON tout_detail.TrxOutIDF = tout_header.TrxOutPK
		JOIN
			MasterWarehouse wh ON tout_header.WhsIdf = wh.WhsPK
		JOIN
			MasterProduct p ON tout_detail.TrxOutDProductIdf = p.ProductPK
		GROUP BY
			wh.WhsName,
			p.ProductName
		ORDER BY
			"warehouse"
    `
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reportResults []models.ReportResult
	for rows.Next() {
		var result models.ReportResult
		err := rows.Scan(&result.Warehouse, &result.Product, &result.QtyDus, &result.QtyPcs, &result.Type)
		if err != nil {
			return nil, err
		}
		reportResults = append(reportResults, result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reportResults, nil
}
