package http

import (
	"marketplace-bhs-test/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "marketplace-bhs-test/docs"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(router *gin.Engine, service service.UserService, MW gin.HandlerFunc) {
	handler := &UserHandler{
		userService: service,
	}

	router.POST("/sign-up", handler.SingUp)
	router.POST("/sign-in", handler.SignIn)
	router.POST("/sign-out", handler.SignOut)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRoutes := router.Group("/users")
	userRoutes.Use(MW)
	{
		userRoutes.PATCH("/:id/balance/:count", handler.UpdateBalance)
	}

}

// SingUp godoc
// @Summary Sign up a new user
// @Description Register a new user with the provided credentials
// @Tags users
// @Accept json
// @Produce json
// @Param input body service.SignUpInput true "User credentials"
// @Success 201
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sign-up [post]
func (h *UserHandler) SingUp(c *gin.Context) {
	var inp service.SignUpInput

	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read user info",
		})
		return
	}

	if err := h.userService.SignUp(c.Request.Context(), &inp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to sign up",
		})
		return
	}

	c.Status(http.StatusCreated)
}

// SignIn godoc
// @Summary Sign in a user
// @Description Authenticate a user and return access and refresh tokens
// @Tags users
// @Accept json
// @Produce json
// @Param input body service.SignUpInput true "User credentials"
// @Success 202 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sign-in [post]
func (h *UserHandler) SignIn(c *gin.Context) {
	var inp service.SignInInput

	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read user info",
		})
		return
	}

	res, err := h.userService.SignIn(c.Request.Context(), &inp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to sign in",
		})
		return
	}

	c.SetCookie("access_token", res.AccessToken, int(res.AccessTokenTTL), "/", "localhost", false, true)
	c.SetCookie("refresh_token", res.RefreshToken, int(res.RefreshTokenTTL), "/", "localhost", false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"AccessToken":  res.AccessToken,
		"RefreshToken": res.RefreshToken,
	})
}

// SignOut godoc
// @Summary Sign out a user
// @Description Clear the user's access and refresh tokens
// @Tags users
// @Success 200
// @Router /sign-out [post]
func (h *UserHandler) SignOut(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	c.Status(http.StatusOK)
}

// UpdateBalance godoc
// @Summary Update user's balance
// @Description Update the user's balance by a specified count
// @Tags users
// @Param count path float64 true "Balance change amount"
// @Param id path uint64 true "User id"
// @Success 200
// @Failure 400 {object} map[string]string
// @Router /users/{id}/balance/{count} [patch]
func (h *UserHandler) UpdateBalance(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User ID",
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

	if userID != id {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	countStr := c.Param("count")
	count, err := strconv.ParseFloat(countStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid count",
		})
		return
	}

	if err := h.userService.UpdateBalance(c.Request.Context(), userID, count); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update balance",
		})
		return
	}

	c.Status(http.StatusOK)
}
