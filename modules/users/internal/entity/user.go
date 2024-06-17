package entity

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SearchUserParams struct {
	Username string
}
