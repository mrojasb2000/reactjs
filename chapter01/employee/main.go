package main

import (
	"fmt"
	"time"

	"github.com/mrojasb2000/udemy/reactjs/chapter01/employee/models"
)

func main() {
	var user models.User
	user.New(1, "Mavro", time.Now(), true)
	fmt.Println(user)
}
