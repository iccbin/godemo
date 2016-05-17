package models

type User struct {

	ID uint `gorm:"primary_key"`
	Name string
	Password string

}

func (u User) TableName() string {
	return "users"
}



