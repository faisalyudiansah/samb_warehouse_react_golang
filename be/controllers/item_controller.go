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

func (uc *ItemController) CreatePenerimaanBarang(c *gin.Context) {
	req := new(dtos.PenerimaanBarang)
	if err := c.ShouldBindJSON(req); err != nil {
		c.Error(err)
		return
	}

	ginutils.ResponseOKPlain(c)
}
