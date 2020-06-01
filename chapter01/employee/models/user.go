package models

import "time"

// User object model
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// New - create basic user
func (user *User) New(id int, name string, created time.Time, status bool) {
	user.ID = id
	user.Name = name
	user.CreatedAt = created
	user.Status = status
}
