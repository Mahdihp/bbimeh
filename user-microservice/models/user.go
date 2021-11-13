package models

import "time"

type User struct {
	Id        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Username  string
	Password  string

	CreatedAt time.Time
	UpdatedAt time.Time
}
