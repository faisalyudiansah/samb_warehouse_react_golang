package controllers

import (
	"server/dtos"
	"server/helpers/ginutils"
	"server/services"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	ItemService services.ItemServiceInterface
}

func NewItemController(
	ItemService services.ItemServiceInterface,
) *ItemController {
	return &ItemController{
		ItemService: ItemService,
	}
}

func (uc *ItemController) CreatePenerimaanBarang(ctx *gin.Context) {
	req := new(dtos.PenerimaanBarang)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}
	err := uc.ItemService.CreatePenerimaanBarangService(ctx, *req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOKPlain(ctx)
}

func (uc *ItemController) CreatePengeluaranBarang(ctx *gin.Context) {
	req := new(dtos.PengeluaranBarang)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}
	err := uc.ItemService.CreatePengeluaranBarangService(ctx, *req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOKPlain(ctx)
}

func (uc *ItemController) GetReportResult(ctx *gin.Context) {
	res, err := uc.ItemService.GetReportResultService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOK(ctx, res)
}

func (uc *ItemController) GetWarehouse(ctx *gin.Context) {
	res, err := uc.ItemService.GetWarehouseService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOK(ctx, res)
}

func (uc *ItemController) GetProduct(ctx *gin.Context) {
	res, err := uc.ItemService.GetProductService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOK(ctx, res)
}

func (uc *ItemController) GetSupplier(ctx *gin.Context) {
	res, err := uc.ItemService.GetSupplierService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOK(ctx, res)
}
