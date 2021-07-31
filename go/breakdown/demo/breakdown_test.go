package demo

import (
	"testing"
	"time"

	"breakdown/client/mysql"
	"breakdown/client/redis"
	"breakdown/model"
)

func TestBreakdownGetUser(t *testing.T) {
	var username string
	username = "dabai"
	GetUser(&username)
	time.Sleep(time.Second * 10)
}

func TestBreakdownGetBreakdownUser(t *testing.T) {
	var username string
	username = "dabai"
	GetBreakdownUser(&username)
	time.Sleep(time.Second * 10)
}

func init() {
	mysql.InitMysql()
	redis.InitRedis()
	model.InitModel()
}
