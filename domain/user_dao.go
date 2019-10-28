package domain

import (
	"fmt"
	"net/http"

	"github.com/zoharngo/golang-microservices/mvc/utills"
)

var (
	users = map[uint64]*User{
		123: {ID: 123, FirstName: "Fede", LastName: "Leon", Email: "zoharngo@gmail.com"},
	}

	// UserDao - Interface for data acceses object for User domain
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

// User - User entity
type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type userDaoInterface interface {
	GetUser(userID uint64) (*User, *utills.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userID uint64) (*User, *utills.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utills.ApplicationError{
		Message: fmt.Sprintf("user %v not exists", userID),
		Status:  "not_found",
		Code:    http.StatusNotFound,
	}
}
