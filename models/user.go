package models

type User struct {
	User_id  uint //`gorm:"primaryKey; autoIncrement"`
	Username string
	Password string
}
