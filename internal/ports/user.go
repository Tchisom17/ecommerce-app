package ports

import (
	"ecommerce-app/internal/core/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByPhoneNumber(phone string) (*domain.User, error)
	Create(user *domain.User) error
	FindAllUsers() ([]domain.User, error)
	WithTx(tx *gorm.DB) UserRepository
}
type UserService interface {
	FindByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByPhoneNumber(phone string) (*domain.User, error)
	Create(user *domain.User) error
	FindAllUsers() ([]domain.User, error)
}
type UserHandler interface {
	FindByID(c *gin.Context)
	FindByEmail(c *gin.Context)
	FindByPhoneNumber(c *gin.Context)
	Create(c *gin.Context)
	FindAllUsers(c *gin.Context)
	//Login(c *gin.Context)
}
