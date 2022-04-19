package repository

import (
	"ecommerce-app/internal/core/domain"
	"ecommerce-app/internal/ports"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of user repository
func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) FindByPhoneNumber(phone string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(&user).Error
}
func (r *userRepository) FindAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *userRepository) WithTx(tx *gorm.DB) ports.UserRepository {
	return NewUserRepository(tx)
}
