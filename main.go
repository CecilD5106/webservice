package main

import (
	"fmt"

	"github.com/CecilD5106/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillian",
	}
	fmt.Println(u)
}
