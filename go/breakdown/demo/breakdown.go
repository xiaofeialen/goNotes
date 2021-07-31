package demo

import (
	"encoding/json"
	"fmt"
	"time"

	"breakdown/client/redis"
	"breakdown/model"
)

// 缓存击穿
func GetUser(username *string) {
	for i := 0; i < 20; i++ {
		go getUserInfo(username, i)
	}

}

func getUserInfo(username *string, index int) error {
	var userInfo model.User
	user, err := redis.GetString(*username)
	if err != nil {
		fmt.Printf("redis get username is error:%v", err)
	}
	if user != "" {
		json.Unmarshal([]byte(user), &userInfo)
	} else {
		userInfo, err = model.GetUserByName(*username)
		fmt.Println("reader mysql...", index)
		if err != nil {
			fmt.Printf("get mysql is error:%v", err)
			return err
		}
		bytes, _ := json.Marshal(userInfo)
		err = redis.SetString(*username, bytes)
		fmt.Println(err)
	}
	fmt.Println(userInfo)
	return nil
}

// 防止缓存击穿
func GetBreakdownUser(username *string) {
	for i := 0; i < 20; i++ {
		go getBreakdownUserInfo(username, i)
	}
}

func getBreakdownUserInfo(username *string, index int) error {
	var userInfo model.User
REDIS:
	user, err := redis.GetString(*username)
	if err != nil {
		fmt.Printf("redis get username is error:%v", err)
	}
	userKey := fmt.Sprintf("%s_nx", *username)

	if user != "" {
		json.Unmarshal([]byte(user), &userInfo)
	} else {
		if !redis.SetNx(userKey) {
			time.Sleep(time.Second * 1)
			goto REDIS
		}
		userInfo, err = model.GetUserByName(*username)
		fmt.Println("reader mysql...", index)
		if err != nil {
			fmt.Printf("get mysql is error:%v", err)
			return err
		}
		bytes, _ := json.Marshal(userInfo)
		err = redis.SetString(*username, bytes)
		fmt.Println(err)
	}

	fmt.Println(userInfo)
	return nil
}
