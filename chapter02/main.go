package main

import (
	"fmt"
	"time"
)

type user struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Status    bool
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
}
