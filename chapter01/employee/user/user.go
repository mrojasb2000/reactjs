package user

import "time"

// User structure
type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// New - create basic user
func (user *User) New(id int, name string, created time.Time, status bool) {
	user.Id = id
	user.Name = name
	user.CreatedAt = created
	user.Status = status
}
