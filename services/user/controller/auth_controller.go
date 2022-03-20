package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"github.com/nhatlang19/go-monorepo/pkg/helper"
	"github.com/nhatlang19/go-monorepo/pkg/model"
	"github.com/nhatlang19/go-monorepo/services/user/service"

	"github.com/go-playground/validator/v10"
)

type AuthController interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Me(*gin.Context)
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
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

func (u authController) Login(c *gin.Context) {
	log.Print("[AuthController]...Login")
	loginRequest := LoginRequest{}

	v := validator.New()
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := v.Struct(loginRequest)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		errors := make([]string, 0, 1)
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, string(err.Error()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
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

func (u authController) Me(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := fmt.Errorf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
