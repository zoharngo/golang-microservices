package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zoharngo/golang-microservices/controllers"
)

var (
	router *gin.Engine
)

// StartUp - start application http server
func StartUp() error {

	http.HandleFunc("/users", controllers.UserController.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("ListenAndServe error: [%w]", err)
	}

	return nil
}
