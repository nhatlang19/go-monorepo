package route


import (
	"github.com/nhatlang19/go-monorepo/services/user/controller"
	// "github.com/nhatlang19/go-monorepo/services/user/middleware"
	"github.com/nhatlang19/go-monorepo/services/user/repository"
	"github.com/nhatlang19/go-monorepo/services/user/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)


func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	userService := service.NewUserService(userRepository)

	authController := controller.NewAuthController(userService)

	auth := httpRouter.Group("auth")
	auth.POST("/register", authController.Register)

	httpRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "User", "timestamp": time.Now()})
	})
	httpRouter.Run()
}