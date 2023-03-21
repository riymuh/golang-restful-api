package user

import "time"

type User struct {
	ID          int64
	Name        string
	Description string
	Phone       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
