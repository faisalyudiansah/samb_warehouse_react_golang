package services

import (
	"context"
	"server/apperrors"
	"server/dtos"
	"server/helpers"
	"server/repositories"
)

type ItemServiceInterface interface {
	CreatePenerimaanBarangService(ctx context.Context, reqBody dtos.PenerimaanBarang) error
	CreatePengeluaranBarangService(ctx context.Context, reqBody dtos.PengeluaranBarang) error
	GetReportResultService(ctx context.Context) ([]dtos.ResponseReportResult, error)
	GetWarehouseService(ctx context.Context) ([]dtos.ResponseWarehouse, error)
	GetProductService(ctx context.Context) ([]dtos.ResponseProduct, error)
	GetSupplierService(ctx context.Context) ([]dtos.ResponseSupplier, error)
}

type ItemServiceImplementation struct {
	ItemRepository       repositories.ItemRepositoryInterface
	TransactorRepository repositories.TransactorRepositoryInterface
}

func NewItemServiceImplementation(
	ItemRepository repositories.ItemRepositoryInterface,
	TransactorRepository repositories.TransactorRepositoryInterface,
) *ItemServiceImplementation {
	return &ItemServiceImplementation{
		ItemRepository:       ItemRepository,
		TransactorRepository: TransactorRepository,
	}
}

func (ir *ItemServiceImplementation) CreatePenerimaanBarangService(ctx context.Context, reqBody dtos.PenerimaanBarang) error {
	err := ir.TransactorRepository.Atomic(ctx, func(cForTx context.Context) error {
		warehouseDB, err := ir.ItemRepository.FindWarehouseById(cForTx, reqBody.WhsIdf)
		if err != nil || warehouseDB.WhsPK == 0 {
			return apperrors.WarehouseIDInvalid
		}
		supplierDB, err := ir.ItemRepository.FindSupplierById(cForTx, reqBody.TrxInSuppIdf)
		if err != nil || supplierDB.SupplierPK == 0 {
			return apperrors.SupplierIDInvalid
		}
		productDB, err := ir.ItemRepository.FindProductById(cForTx, reqBody.TrxInDProductIdf)
		if err != nil || productDB.ProductPK == 0 {
			return apperrors.ProductIDInvalid
		}
		lastTrxInPK, err := ir.ItemRepository.FindLastTrxInPKPenerimaan(cForTx)
		if err != nil {
			return err
		}
		newTrxInNo := helpers.GenerateTrxInNo(lastTrxInPK, "IN")
		res, err := ir.ItemRepository.CreatePenerimaanBarangHeader(cForTx, &reqBody, newTrxInNo)
		if err != nil {
			return err
		}
		err = ir.ItemRepository.CreatePenerimaanBarangDetail(cForTx, &reqBody, *res)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (ir *ItemServiceImplementation) CreatePengeluaranBarangService(ctx context.Context, reqBody dtos.PengeluaranBarang) error {
	err := ir.TransactorRepository.Atomic(ctx, func(cForTx context.Context) error {
		warehouseDB, err := ir.ItemRepository.FindWarehouseById(cForTx, reqBody.WhsIdf)
		if err != nil || warehouseDB.WhsPK == 0 {
			return apperrors.WarehouseIDInvalid
		}
		supplierDB, err := ir.ItemRepository.FindSupplierById(cForTx, reqBody.TrxOutSuppIdf)
		if err != nil || supplierDB.SupplierPK == 0 {
			return apperrors.SupplierIDInvalid
		}
		productDB, err := ir.ItemRepository.FindProductById(cForTx, reqBody.TrxOutDProductIdf)
		if err != nil || productDB.ProductPK == 0 {
			return apperrors.ProductIDInvalid
		}
		lastTrxOutPK, err := ir.ItemRepository.FindLastTrxOutPKPengeluaran(cForTx)
		if err != nil {
			return err
		}
		newTrxOutNo := helpers.GenerateTrxInNo(lastTrxOutPK, "OUT")
		res, err := ir.ItemRepository.CreatePengeluaranBarangHeader(cForTx, &reqBody, newTrxOutNo)
		if err != nil {
			return err
		}
		err = ir.ItemRepository.CreatePengeluaranBarangDetail(cForTx, &reqBody, *res)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (ir *ItemServiceImplementation) GetReportResultService(ctx context.Context) ([]dtos.ResponseReportResult, error) {
	var res []dtos.ResponseReportResult
	err := ir.TransactorRepository.Atomic(ctx, func(cForTx context.Context) error {
		reportData, err := ir.ItemRepository.GetReportData(cForTx)
		if err != nil {
			return err
		}
		res = make([]dtos.ResponseReportResult, len(reportData))
		for i := range reportData {
			res[i] = dtos.ResponseReportResult{
				Warehouse: reportData[i].Warehouse,
				Product:   reportData[i].Product,
				QtyDus:    reportData[i].QtyDus,
				QtyPcs:    reportData[i].QtyPcs,
				Type:      reportData[i].Type,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ir *ItemServiceImplementation) GetWarehouseService(ctx context.Context) ([]dtos.ResponseWarehouse, error) {
	var res []dtos.ResponseWarehouse
	err := ir.TransactorRepository.Atomic(ctx, func(cForTx context.Context) error {
		resDB, err := ir.ItemRepository.GetWarehouseRepository(cForTx)
		if err != nil {
			return err
		}
		res = make([]dtos.ResponseWarehouse, len(resDB))
		for i := range resDB {
			res[i] = dtos.ResponseWarehouse{
				WhsPK:   resDB[i].WhsPK,
				WhsName: resDB[i].WhsName,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ir *ItemServiceImplementation) GetProductService(ctx context.Context) ([]dtos.ResponseProduct, error) {
	var res []dtos.ResponseProduct
	err := ir.TransactorRepository.Atomic(ctx, func(cForTx context.Context) error {
		resDB, err := ir.ItemRepository.GetProductRepository(cForTx)
		if err != nil {
			return err
		}
		res = make([]dtos.ResponseProduct, len(resDB))
		for i := range resDB {
			res[i] = dtos.ResponseProduct{
				ProductPK:   resDB[i].ProductPK,
				ProductName: resDB[i].ProductName,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ir *ItemServiceImplementation) GetSupplierService(ctx context.Context) ([]dtos.ResponseSupplier, error) {
	var res []dtos.ResponseSupplier
	err := ir.TransactorRepository.Atomic(ctx, func(cForTx context.Context) error {
		resDB, err := ir.ItemRepository.GetSupplierRepository(cForTx)
		if err != nil {
			return err
		}
		res = make([]dtos.ResponseSupplier, len(resDB))
		for i := range resDB {
			res[i] = dtos.ResponseSupplier{
				SupplierPK:   resDB[i].SupplierPK,
				SupplierName: resDB[i].SupplierName,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
