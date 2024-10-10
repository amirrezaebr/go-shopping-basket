package models

import (
	"time"
)

type Basket struct {
	Id         uint  `gorm:"primaryKey"`
	Created_at time.Time
	Updated_at time.Time
	Data       string  
	State      string   
	UserId     uint
}
