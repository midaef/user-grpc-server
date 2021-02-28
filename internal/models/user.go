package models

import "time"

// User ...
type User struct {
	User         string `db:"username"`
	Name         string `db:"name"`
	Password     string `db:"password"`
	RegisteredAt time.Time `db:"registered"`
}
