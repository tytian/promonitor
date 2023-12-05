package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"promonitor/monitor"
	"promonitor/monitor/model"
	"time"
)

type UserListInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Id   uint   `json:"id"`
}

type UserCreateInput struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

func UserList(ctx *gin.Context) {
	var input UserListInput
	err := ctx.Bind(&input)
	if err != nil {
		log.Fatalf("get params error")
		return
	}
	var users []model.User
	monitor.PromDB.Find(&users)
	ctx.JSON(200, gin.H{
		"data": users,
		"code": 0,
	})
}

func UserGet(ctx *gin.Context) {
	var input UserListInput
	err := ctx.Bind(&input)
	if err != nil {
		log.Fatalf("get params error")
		return
	}
	var user model.User
	monitor.PromDB.First(&user, "name = ?", input.Name)
	ctx.JSON(200, gin.H{
		"data": user,
		"code": 0,
	})
}

func UserCreate(ctx *gin.Context) {
	var input UserCreateInput
	err := ctx.Bind(&input)
	if err != nil {
		log.Fatalf("get params error")
		return
	}
	var user model.User
	user.Name = input.Name
	user.Age = input.Age
	user.Email = input.Email
	user.Birthday = time.Now()
	if err := monitor.PromDB.Model(&model.User{}).Save(&user).Error; err != nil {
		log.Fatalf("save model error")
		ctx.JSON(500, gin.H{
			"error": err.Error(),
			"code":  1,
		})
	}
	ctx.JSON(200, gin.H{
		"id":   user.ID,
		"code": 0,
	})
}
