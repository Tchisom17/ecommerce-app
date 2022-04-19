package http

import (
	"ecommerce-app/internal/common/types"
	"ecommerce-app/internal/core/domain"
	"ecommerce-app/internal/ports"
	"ecommerce-app/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type userHandler struct {
	UserService ports.UserService
	logger      *log.Logger
	handlerName string
}

var (
	result  utils.Result
	message types.Messages
)

// NewUserHandler function creates a new instance for user handler
func NewUserHandler(cs ports.UserService, l *log.Logger, n string) ports.UserHandler {
	return userHandler{
		UserService: cs,
		logger:      l,
		handlerName: n,
	}
}

func (u userHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	user, err := u.UserService.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		u.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(user, message.GetResponseMessage(u.handlerName, types.OKAY)))
}

func (u userHandler) FindByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := u.UserService.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		u.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(user, message.GetResponseMessage(u.handlerName, types.OKAY)))
}

func (u userHandler) FindByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Param("phone")
	user, err := u.UserService.FindByPhoneNumber(phoneNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		u.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(user, message.GetResponseMessage(u.handlerName, types.OKAY)))
}

func (u userHandler) Create(c *gin.Context) {
	body := &domain.User{}
	if err := c.ShouldBindJSON(&body); err != nil {
		u.logger.Error(err, "###")
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	_, err := u.UserService.FindByEmail(body.Email)
	if err == nil {
		u.logger.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "email already exists"})
		return
	}
	_, err = u.UserService.FindByPhoneNumber(body.Phone1)
	if err == nil {
		u.logger.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "phone number already exists"})
		return
	}
	user := &domain.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	user.Password = string(hash)
	err = u.UserService.Create(user)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, result.ReturnSuccessResult(user, message.GetResponseMessage(u.handlerName, types.CREATED)))
}

func (u userHandler) FindAllUsers(c *gin.Context) {
	users, err := u.UserService.FindAllUsers()
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(users, message.GetResponseMessage(u.handlerName, types.OKAY)))
}

//func (u userHandler) Login(c *gin.Context) {
//	user := &domain.User{}
//	request := &struct {
//		Email    string `json:"email" binding:"required"`
//		Password string `json:"password" binding:"required"`
//	}{}
//	_, err := u.UserService.FindByEmail(request.Email)
//	if err != nil {
//		if inactiveErr, ok := err.(helpers.InActiveUserError); ok {
//			c.JSON(http.StatusBadRequest, result.ReturnErrorResult(inactiveErr.Error()))
//			return
//		}
//		u.logger.Error(err)
//		c.JSON(http.StatusUnauthorized, result.ReturnErrorResult("user not found"))
//		return
//	}
//	if user.IsActive == false {
//		c.JSON(http.StatusUnauthorized, result.ReturnErrorResult("email not verified"))
//		return
//	}
//
//	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
//	if err != nil {
//		u.logger.Error(err)
//		c.JSON(http.StatusUnauthorized, result.ReturnErrorResult("invalid password"))
//		return
//	}
//	token, err := utils.CreateToken(user)
//	if err != nil {
//		u.logger.Error(err.Error())
//		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult("internal server error"))
//		return
//	}
//	c.JSON(http.StatusOK, result.ReturnAuthResult(user, token))
//}
