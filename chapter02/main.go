package main

import (
	"fmt"
	"time"

	"github.com/mrojasb2000/udemy/reactjs/chapter02/model"
)

type user struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

type pepe struct {
	model.User // Herencia en formato golang
}

func main() {
	var usr = new(user)
	usr.ID = 1
	usr.Name = "Mavro"
	usr.CreatedAt = time.Now()
	usr.Status = true

	fmt.Printf("\nID     : %d", usr.ID)
	fmt.Printf("\nName   : %s", usr.Name)
	fmt.Printf("\nCreated: %s", usr.CreatedAt)
	fmt.Printf("\nStatus : %t\n", usr.Status)

	var newUser = new(model.User)
	fmt.Printf("\nID     : %d", newUser.ID)
	fmt.Printf("\nName   : %s", newUser.Name)
	fmt.Printf("\nCreated: %s", newUser.CreatedAt)
	fmt.Printf("\nStatus : %t\n", newUser.Status)

	threeUser := new(pepe)
	threeUser.New(3, "username", time.Now(), false)
	fmt.Printf("\nID     : %d", threeUser.ID)
	fmt.Printf("\nName   : %s", threeUser.Name)
	fmt.Printf("\nCreated: %s", threeUser.CreatedAt)
	fmt.Printf("\nStatus : %t\n", threeUser.Status)
}
