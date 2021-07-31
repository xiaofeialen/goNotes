package common

import (
	"encoding/json"
	"fmt"
	"time"

	"bloom/client/redis"
	"bloom/model"
	"github.com/bits-and-blooms/bloom/v3"
)

var bloomFilter *bloom.BloomFilter

func InitBloom() {
	bloomFilter = bloom.NewWithEstimates(1000000, 0.3)
}

func GetBloom() *bloom.BloomFilter {
	return bloomFilter
}

func Add(key []byte) {
	bloomFilter.Add(key)
}

func CheckKey(key []byte) bool {
	return bloomFilter.Test(key)
}

func InitFunc() {
	addUserFilter()
	addUserRedis()
}

func addUserFilter() {
	users, err := model.UserInfoToFilter()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range users {
		Add([]byte(v.Username))
	}
}

func addUserRedis() {
	users, err := model.UserInfoToFilter()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, v := range users {
		bytes, _ := json.Marshal(v)
		redis.SetExString(v.Username, bytes, time.Duration(10+i))
	}
}
