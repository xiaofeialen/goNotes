package mysql

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var once sync.Once

func InitMysql() {
	initMysql()
}

func initMysql() {
	once.Do(func() {
		user := "xx"
		password := "xxx"
		host := "127.0.0.1"
		port := "3306"
		database := "test"
		mysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, port, database)
		var err error
		db, err = gorm.Open("mysql", mysql)
		if err != nil {
			os.Exit(1)
		}
		db.LogMode(true)
		db.DB().SetMaxOpenConns(10)
		db.DB().SetConnMaxLifetime(60 * time.Second)
	})
}

// GetMysqlDb get mysql client
func GetMysqlDb() *gorm.DB {
	return db
}

func IsNotFound(err error) bool {
	if err != nil && err.Error() == "record not found" {
		return true
	}
	return false
}
