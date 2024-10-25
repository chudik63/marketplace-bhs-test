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
	}
}

func (h *UserHandler) SingUp(c *gin.Context) {
	var inp service.SignInInput

	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read user info",
		})

		return
	}

	if err := h.userService.SignUp(&inp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to sign up",
		})

		return
	}

	c.Status(http.StatusCreated)
}

func (h *UserHandler) SignIn(c *gin.Context) {

}
