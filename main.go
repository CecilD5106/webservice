package main

import (
	"net/http"

	"github.com/CecilD5106/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
