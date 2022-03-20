package route

import (
	"log"

	"github.com/nhatlang19/go-monorepo/services/user/controller"
	"github.com/nhatlang19/go-monorepo/services/user/middleware"

	"github.com/nhatlang19/go-monorepo/services/user/repository"
	"github.com/nhatlang19/go-monorepo/services/user/service"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	client "github.com/nhatlang19/go-monorepo/services/user/client"
)

func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}

	mailClient := client.NewMailClient()
	userService := service.NewUserService(userRepository, mailClient)

	authController := controller.NewAuthController(userService)

	tokenService := service.NewTokenService()

	auth := httpRouter.Group("auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
	auth.GET("/me", middleware.AuthUser(tokenService), authController.Me)

	httpRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "User", "timestamp": time.Now()})
	})
	httpRouter.Run()
}
