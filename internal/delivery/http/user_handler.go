package http

import (
	"marketplace-bhs-test/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(router *gin.Engine, service service.UserService) {
	handler := &UserHandler{
		userService: service,
	}

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/sign-up", handler.SingUp)
		userRoutes.POST("/sign-in", handler.SignIn)
		userRoutes.POST("/sign-out", handler.SignOut)
	}
}

func (h *UserHandler) SingUp(c *gin.Context) {
	var inp service.SignUpInput

	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read user info",
		})

		return
	}

	if err := h.userService.SignUp(c, &inp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to sign up",
		})

		return
	}

	c.Status(http.StatusCreated)
}

func (h *UserHandler) SignIn(c *gin.Context) {
	var inp service.SignUpInput

	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read user info",
		})

		return
	}

	res, err := h.userService.SignIn(c, &inp)
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

func (h *UserHandler) SignOut(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	c.Status(http.StatusOK)
}
