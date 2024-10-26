package http

import (
	"marketplace-bhs-test/internal/entity"
	"marketplace-bhs-test/internal/service"
	"net/http"
	"strconv"

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
		assetRoutes.DELETE("/assets/:id", handler.DeleteAsset)
		assetRoutes.PATCH("/assets/:id", handler.BuyAsset)
	}
}

func (h *AssetHandler) CreateAsset(c *gin.Context) {
	var asset entity.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read asset info",
		})

		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization required",
		})

		return
	}
	asset.UserID = uint(userID.(float64))

	if err := h.assetService.CreateAsset(c.Request.Context(), &asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create asset",
		})

		return
	}

	c.JSON(http.StatusCreated, asset)
}

func (h *AssetHandler) DeleteAsset(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Asset ID",
		})

		return
	}

	if err := h.assetService.DeleteAsset(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete asset",
		})

		return
	}

	c.Status(http.StatusOK)
}

func (h *AssetHandler) BuyAsset(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Asset ID",
		})

		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization required",
		})

		return
	}

	if err := h.assetService.BuyAsset(c.Request.Context(), uint(id), uint(userID.(float64))); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to buy asset",
		})

		return
	}

	c.Status(http.StatusOK)
}
