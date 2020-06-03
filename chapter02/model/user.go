package model

import "time"

// User basic struct
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// New - Add new user entity
func (user *User) New(id int, name string, created time.Time, status bool) {
	user.ID = id
	user.Name = name
	user.CreatedAt = created
	user.Status = status
}
