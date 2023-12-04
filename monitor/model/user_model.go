package model

import "time"

type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
}
