package controller

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"github.com/nhatlang19/go-monorepo/services/user/model"
	"github.com/nhatlang19/go-monorepo/services/user/service"
)

type AuthController interface {
	Register(*gin.Context)
}

type authController struct {
	userService service.UserService
}

func NewAuthController(s service.UserService) AuthController {
	return authController{
		userService: s,
	}
}

func (u authController) Register(c *gin.Context) {
	log.Print("[AuthController]...Register")
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userService.Save(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
