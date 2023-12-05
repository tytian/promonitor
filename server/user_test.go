package server

import (
	"promonitor/monitor"
	"promonitor/monitor/model"
	"testing"
)

func TestUserList(t *testing.T) {
	monitor.PromDB.AutoMigrate(&model.User{})
	monitor.PromDB.Model(&model.User{}).AddUniqueIndex("name_email", "id", "name", "email")
	monitor.PromDB.Create(&model.User{Name: "Tia", Age: 16, Email: "tia@gmail.com"})
	monitor.PromDB.Create(&model.User{Name: "Bob", Age: 18, Email: "bob@gmail.com"})
}
