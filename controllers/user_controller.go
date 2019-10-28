package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/zoharngo/golang-microservices/mvc/services"
	"github.com/zoharngo/golang-microservices/mvc/utills"
)

type userController struct{}

type userControllerInterface interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}

var (
	logger = log.New(os.Stdout, "controllers => ", log.Ltime)
	// UserController - controller interface for user CRUD operation
	UserController userControllerInterface
)

func init() {
	UserController = &userController{}
}

// GetUser - controller method to retrive users from service
func (u *userController) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		appError := &utills.ApplicationError{
			Message: "user_id must be number",
			Status:  "bad_request",
			Code:    http.StatusBadRequest,
		}

		w.WriteHeader(appError.Code)
		serilaizedJSONErrorgo, _ := appError.Error()
		logger.Println(string(serilaizedJSONErrorgo))

		w.Write(serilaizedJSONErrorgo)

		return
	}

	user, appError := services.UsersService.GetUser(uint64(userID))

	if appError != nil {
		w.WriteHeader(appError.Code)
		serilaizedJSONErrorgo, _ := appError.Error()
		logger.Println(string(serilaizedJSONErrorgo))

		w.Write(serilaizedJSONErrorgo)

		return
	}

	userjson, _ := json.Marshal(user)
	logger.Println(string(userjson))

	w.Write(userjson)
}
