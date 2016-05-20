package models

type User struct {

	ID uint `gorm:"primary_key"`
	Name string
	Password string
	Money float64
	Age int64
	Gender bool
}

func (u User) TableName() string {
	return "users"
}



