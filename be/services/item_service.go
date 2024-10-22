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
		lastTrxInPK, err := ir.ItemRepository.FindLastTrxInPK(cForTx)
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
