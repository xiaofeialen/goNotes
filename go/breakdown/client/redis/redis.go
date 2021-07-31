package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisPool redis.UniversalClient
var ctx = context.Background()

func InitRedis() {
	redisPool = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{"127.0.0.1:6379"},
		Password: "",
		//DB:           ,

		MinIdleConns: 100,
		MaxRetries:   3,
	})
	if err := IsActive(); err != nil {
		panic(err)
	}
}

func IsActive() error {
	return redisPool.Ping(ctx).Err()

}

func GetString(key string) (str string, err error) {
	return redisPool.Get(ctx, key).Result()
}
func SetString(key string, value interface{}) (err error) {
	return redisPool.Set(ctx, key, value, time.Duration(time.Second*100)).Err()
}
func SetExString(key string, value interface{}, ex time.Duration) (err error) {
	return redisPool.Set(ctx, key, value, ex).Err()
}
func SetNx(key string) bool {
	return redisPool.SetNX(ctx, key, "1", time.Duration(time.Second*90)).Val()
}
func DelString(key string) (err error) {
	return redisPool.Del(ctx, key).Err()
}
