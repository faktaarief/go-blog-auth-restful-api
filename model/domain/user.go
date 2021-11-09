package domain

import "time"

type User struct {
	Id        int
	Name      string `gorm:"size:100"`
	Email     string `gorm:"size:100"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
