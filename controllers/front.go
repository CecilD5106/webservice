package controllers

import "net/http"

//RegisterControllers will register the controller
func RegisterControllers() {
	uc := NewUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
