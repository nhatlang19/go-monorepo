package controller

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"github.com/nhatlang19/go-monorepo/pkg/helper"
	"github.com/nhatlang19/go-monorepo/services/user/model"
	"github.com/nhatlang19/go-monorepo/services/user/service"
)

type AuthController interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type authController struct {
	userService service.UserService
}

func NewAuthController(s service.UserService) AuthController {
	return authController{
		userService: s,
	}
}

type Auth struct {
	Jwt string `json:"jwt"`
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

	jwtToken := helper.NewJwtToken()
	token, err := jwtToken.CreateToken(uint64(user.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Auth{Jwt: token}})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u authController) Login(c *gin.Context) {
	log.Print("[AuthController]...Login")
	loginRequest := LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err, httpCode := u.userService.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		c.JSON(httpCode, gin.H{"error": err.Error()})
		return
	}

	jwtToken := helper.NewJwtToken()
	token, err := jwtToken.CreateToken(uint64(user.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Auth{Jwt: token}})
}
