package auth

import "time"

type User struct {
	Id        int
	Username  string
	Email     string
	Birthday  *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
