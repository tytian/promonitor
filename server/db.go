package server

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"time"
)

var PromDB *gorm.DB

func InitDB() {
	time.Sleep(5 * time.Second)
	var sqlDb *sql.DB
	db, err := gorm.Open("mysql", "root:mariadb@tcp(mydb:3306)/monitor?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Failed to open database:%s, error: %v", "monitor", err)
	}
	db.Dialect().SetDB(sqlDb)
	db.LogMode(true)
	db.SetLogger(&gormLogger{})
	// 取消表名的复数形式，使得表名和结构体名称一致
	db.SingularTable(true)
	PromDB = db
}

func CloseDB() {
	if PromDB != nil {
		PromDB.Close()
	}
}

type gormLogger struct{}

func (g *gormLogger) Print(v ...any) {
	log.Print(v)
}

func (g *gormLogger) Println(v ...any) {
	log.Println(v)
}
