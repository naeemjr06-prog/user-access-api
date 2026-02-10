package domain

import "time"

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Role      Role
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
