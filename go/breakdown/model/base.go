package model

import (
	"breakdown/client/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitModel() {
	db = mysql.GetMysqlDb()
}
