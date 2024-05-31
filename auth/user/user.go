package auth

import (
	"errors"
	"time"
)

type User struct {
	Id        int
	Username  string
	Email     string
	Birthday  *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

var timeUTC = time.UTC
var id1Birthday = time.Date(1999, 3, 12, 0, 0, 0, 0, timeUTC)
var id2Birthday = time.Date(1973, 11, 10, 0, 0, 0, 0, timeUTC)
var id3Birthday = time.Date(1984, 7, 30, 0, 0, 0, 0, timeUTC)

var userLits = []User{
	{
		Id:       1,
		Username: "bangkoker",
		Email:    "bangkoker@gmail.com",
		Birthday: &id1Birthday,
	},
	{
		Id:       2,
		Username: "chiangmaiker",
		Email:    "chiangmaiker@gmail.com",
		Birthday: &id2Birthday,
	},
	{
		Id:       3,
		Username: "phuketker",
		Email:    "phuketker@gmail.com",
		Birthday: &id3Birthday,
	},
}

func (u User) Get(id int) (User, error) {
	var foundUser = User{}
	for _, user := range userLits {
		if user.Id == id {
			foundUser = user
		}
	}
	if foundUser.Id == 0 {
		return foundUser, errors.New("not found user")
	}

	return foundUser, nil
}
