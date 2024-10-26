package http

import (
	"marketplace-bhs-test/internal/service"

	"github.com/gin-gonic/gin"
)

type AssetHandler struct {
	assetService service.AssetService
}

func NewAssetHandler(router *gin.Engine, service service.AssetService) {
	handler := &AssetHandler{
		assetService: service,
	}

	assetRoutes := router.Group("/marketplace")
	{
		assetRoutes.POST("/assets", handler.CreateAsset)
		assetRoutes.DELETE("/assets", handler.DeleteAsset)
	}
}

func (h *AssetHandler) CreateAsset(c *gin.Context) {

}

func (h *AssetHandler) DeleteAsset(c *gin.Context) {

}
