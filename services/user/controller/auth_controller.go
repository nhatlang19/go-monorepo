package controller

import  (
	"github.com/gin-gonic/gin"

	"github.com/nhatlang19/go-monorepo/services/user/model"
	"github.com/nhatlang19/go-monorepo/services/user/service"
	"net/http"
	"log"
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

	c.JSON(http.StatusOK, gin.H{"data": user})
}