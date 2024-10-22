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
