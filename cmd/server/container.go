package server

import (
	"ecommerce-app/internal/adapters/http"
	"ecommerce-app/internal/adapters/repository"
	"ecommerce-app/internal/core/services"
	"ecommerce-app/pkg/config"
	"ecommerce-app/pkg/logger"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Injection inject all dependencies
func Injection() {
	var logging *log.Logger

	if config.Instance.Env == "development" {
		logging = logger.NewLogger(log.New()).MakeLogger("logs/info", true)
		logging.Info("Log setup complete")
	} else {
		logging = logger.NewLogger(log.New()).Hook()
	}

	var (
		ginRoutes      = NewGinRouter(gin.Default())
		userRepository = repository.NewUserRepository(DBConnection)
		userService    = services.NewUserService(userRepository, logging)
		userHandler    = http.NewUserHandler(userService, logging, "User")
	)

	v1 := ginRoutes.GROUP("v1")
	user := v1.Group("/user")
	user.GET("/:id", userHandler.FindByID)
	user.GET("/", userHandler.FindAllUsers)
	//company.POST("/", companyHandler.CreateCompany)
	//company.DELETE("/:id", companyHandler.DeleteCompany)
	//company.PATCH("/:id", companyHandler.UpdateCompany)

	err := ginRoutes.SERVE()

	if err != nil {
		return
	}

}
