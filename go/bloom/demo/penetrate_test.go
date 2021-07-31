package demo

import (
	"testing"

	"bloom/client/mysql"
	"bloom/client/redis"
	"bloom/common"
	"bloom/model"
)

// TestGetUserInfo 测试缓存穿透
func TestGetUserInfo(t *testing.T) {
	var username string
	username = "dabai2"
	for i := 0; i < 10; i++ {
		GetUserInfo(&username)
	}

}

func TestPenetrateGetUserInfo(t *testing.T) {

	var username string
	username = "dabai2"
	for i := 0; i < 10; i++ {
		PenetrateGetUserInfo(&username)
	}
}

func init() {
	mysql.InitMysql()
	redis.InitRedis()
	model.InitModel()
	common.InitBloom()
	common.InitFunc() //数据预热
}
