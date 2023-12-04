package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"promonitor/monitor/model"
)

type UserListInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Id   uint   `json:"id"`
}

func UserList(ctx *gin.Context) {
	var input UserListInput
	err := ctx.Bind(&input)
	if err != nil {
		log.Fatalf("get params error")
		return
	}
	var users []model.User
	PromDB.Find(&users)

}

func UserGet(ctx *gin.Context) {
	var input UserListInput
	err := ctx.Bind(&input)
	if err != nil {
		log.Fatalf("get params error")
		return
	}
	var user model.User
	PromDB.First(&user, "name = ?", input.Name)
}
