package model

import (
	"bloom/client/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitModel() {
	db = mysql.GetMysqlDb()
}
