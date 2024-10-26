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

func NewAssetHandler(router *gin.Engine, service service.AssetService, MW gin.HandlerFunc) {
	handler := &AssetHandler{
		assetService: service,
	}

	assetRoutes := router.Group("/marketplace")
	assetRoutes.Use(MW)
	{
		assetRoutes.POST("/assets", handler.CreateAsset)
		assetRoutes.DELETE("/assets/:id", handler.DeleteAsset)
		assetRoutes.PATCH("/assets/:id", handler.BuyAsset)
	}

}

// @Summary Create an asset
// @Description Create a new asset
// @Tags assets
// @Accept json
// @Produce json
// @Param asset body entity.Asset true "Asset data"
// @Success 201 {object} entity.Asset
// @Failure 400 {object} map[string]string
// @Router /marketplace/assets [post]
func (h *AssetHandler) CreateAsset(c *gin.Context) {
	var asset entity.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read asset info",
		})

		return
	}

	untypedUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization required",
		})
		return
	}
	userID := untypedUserID.(uint64)
	asset.UserID = userID

	if err := h.assetService.CreateAsset(c.Request.Context(), &asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create asset",
		})

		return
	}

	c.JSON(http.StatusCreated, asset)
}

// DeleteAsset godoc
// @Summary Delete an asset
// @Description Delete an asset by ID
// @Tags assets
// @Produce json
// @Param id path int true "Asset ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Router /marketplace/assets/{id} [delete]
func (h *AssetHandler) DeleteAsset(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Asset ID",
		})

		return
	}

	if err := h.assetService.DeleteAsset(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete asset",
		})

		return
	}

	c.Status(http.StatusOK)
}

// BuyAsset godoc
// @Summary Buy an asset
// @Description Purchase an asset by ID
// @Tags assets
// @Produce json
// @Param id path int true "Asset ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Router /marketplace/assets/{id} [patch]
func (h *AssetHandler) BuyAsset(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Asset ID",
		})

		return
	}

	untypedUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization required",
		})
		return
	}
	userID := untypedUserID.(uint64)

	if err := h.assetService.BuyAsset(c.Request.Context(), id, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to buy asset",
		})

		return
	}

	c.Status(http.StatusOK)
}
