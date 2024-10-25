package http

import (
	"marketplace-bhs-test/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	SignUp(input service.SignInInput)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(router *gin.Engine, service *service.UserService) {
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

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read user info",
		})
	}

	h.userService.SignUp(inp)

	c.Status(http.StatusCreated)
}

func (h *UserHandler) SignIn(c *gin.Context) {

}
