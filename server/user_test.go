package server

import (
	"promonitor/monitor/model"
	"testing"
)

func TestUserList(t *testing.T) {
	PromDB.AutoMigrate(&model.User{})
	PromDB.Model(&model.User{}).AddUniqueIndex("name_email", "id", "name", "email")
	PromDB.Create(&model.User{Name: "Tia", Age: 16, Email: "tia@gmail.com"})
	PromDB.Create(&model.User{Name: "Bob", Age: 18, Email: "bob@gmail.com"})
}
