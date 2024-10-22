package controllers

import (
	"server/dtos"
	"server/helpers"
	"server/services"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	ItemService       services.ItemServiceInterface
	ValidationReqBody helpers.ValidationReqBodyInterface
}

func NewItemController(
	ItemService services.ItemServiceInterface,
	ValidationReqBody helpers.ValidationReqBodyInterface,
) *ItemController {
	return &ItemController{
		ItemService:       ItemService,
		ValidationReqBody: ValidationReqBody,
	}
}

func (uc *ItemController) PostRegisterUserController(c *gin.Context) {
	var reqBody dtos.PenerimaanBarang
	if err := uc.ValidationReqBody.ValidationReqBody(c, &reqBody); err != nil {
		c.Error(err)
		return
	}
}
