package services

import (
	"ecommerce-app/internal/core/domain"
	"ecommerce-app/internal/ports"
	log "github.com/sirupsen/logrus"
)

type userService struct {
	UserRepository ports.UserRepository
	logger         *log.Logger
}

func (u *userService) FindByID(id string) (*domain.User, error) {
	user, err := u.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userService) FindByEmail(email string) (*domain.User, error) {
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userService) FindByPhoneNumber(phone string) (*domain.User, error) {
	user, err := u.UserRepository.FindByPhoneNumber(phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userService) Create(user *domain.User) error {
	err := u.UserRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u userService) FindAllUsers() ([]domain.User, error) {
	users, err := u.UserRepository.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// NewUserService function create a new instance for service
func NewUserService(cr ports.UserRepository, l *log.Logger) ports.UserService {
	return &userService{
		UserRepository: cr,
		logger:         l,
	}
}
