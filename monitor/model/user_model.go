package model

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	log.Debugln("BeforeCreate....")
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
	//monitor.MetricMonitor.ClientHandleRequestSeconds(middleware.TypeMysql, u.TableName(), middleware.OpCreate, middleware.Database)
	log.Debugln("AfterCreate....")
	if tx.Error != nil {
		log.Errorf("sql exec failed: %s", tx.Error.Error())
	}
	return nil
}
