package services

import (
	"github.com/zoharngo/golang-microservices/mvc/domain"
	"github.com/zoharngo/golang-microservices/mvc/utills"
)

type usersService struct{}

var (
	// UsersService - Public variable to handle all user service
	// operations
	UsersService *usersService
)

func init() {
	UsersService = &usersService{}
}

// GetUser - Public
func (u *usersService) GetUser(userID uint64) (*domain.User, *utills.ApplicationError) {
	user, appError := domain.UserDao.GetUser(userID)
	if appError != nil {
		return nil, appError
	}

	return user, nil
}
