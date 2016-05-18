package models

type Temp struct {

	ID uint `gorm:"primary_key"`
	Name string
	Password string
	Money	float32
	IsOk bool
}