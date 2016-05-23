package models

import "time"

type User struct {
	//ID        uint `gorm:"primary_key"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt *time.Time `sql:"index"`
	ID uint `gorm:"primary_key"`
	Name string
	Password string
	Money float64
	Age int64
	Gender bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (u User) TableName() string {
	return "users"
}



