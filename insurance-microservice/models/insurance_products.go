package models

import "time"

type InsuranceProducts struct {
	Id    uint `gorm:"primaryKey"`
	Title string

	CreatedAt time.Time
	UpdatedAt time.Time
}
