package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller"))
}

//NewUserController is a constructor for user controllers
func NewUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile('^/users/(\d+)/?'),
	}
}
