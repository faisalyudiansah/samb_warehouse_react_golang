package servers

import (
	"net/http"
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupItemRoute(router *gin.Engine, ItemController *controllers.ItemController) {
	u := router.Group("/")
	u.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "server is running!"})
	})

	u.POST("/penerimaan-barang", ItemController.CreatePenerimaanBarang)
	u.POST("/pengeluaran-barang", ItemController.CreatePengeluaranBarang)
}
