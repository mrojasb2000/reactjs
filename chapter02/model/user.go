package model

import "time"

// User basic struct
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Status    bool
}
