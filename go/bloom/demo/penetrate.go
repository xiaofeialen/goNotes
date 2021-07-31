package demo

import (
	"bloom/client/redis"
	"bloom/common"
	"bloom/model"
	"encoding/json"
	"fmt"
)

func GetUserInfo(username *string) (err error) {
	if username == nil {
		return
	}
	var userInfo model.User
	user, err := redis.GetString(*username)
	if err != nil {
		fmt.Printf("redis get username is error:%v", err)
	}
	if user != "" {
		json.Unmarshal([]byte(user), &userInfo)
	} else {
		userInfo, err = model.GetUserByName(*username)
		fmt.Println("reader mysql...")
		if err != nil {
			fmt.Printf("get mysql is error:%v", err)
			return err
		}
		bytes, _ := json.Marshal(userInfo)
		err = redis.SetString(*username, bytes)
		fmt.Println(err)
	}
	fmt.Println(userInfo)
	return
}

func PenetrateGetUserInfo(username *string) (err error) {
	if username == nil {
		return
	}
	if !common.CheckKey([]byte(*username)) {
		fmt.Println("username is invalid")
		return
	}
	var userInfo model.User
	user, err := redis.GetString(*username)
	if err != nil {
		fmt.Printf("redis get username is error:%v", err)
	}
	if user != "" {
		json.Unmarshal([]byte(user), &userInfo)
	} else {
		userInfo, err = model.GetUserByName(*username)
		fmt.Println("reader mysql...")
		if err != nil {
			fmt.Printf("get mysql is error:%v", err)
			return err
		}
		bytes, _ := json.Marshal(userInfo)
		err = redis.SetString(*username, bytes)
		fmt.Println(err)
	}
	fmt.Println(userInfo)
	return
}
