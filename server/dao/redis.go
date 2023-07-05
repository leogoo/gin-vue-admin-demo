package dao

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var MyRedis *redis.Client

func InitRedis() {
	config := viper.New()
	config.SetConfigFile("config.yaml") // 指定配置文件路径

	err := config.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	url := config.Get("redis.url")
	password := config.Get("redis.password")
	MyRedis = redis.NewClient(&redis.Options{
		Addr:       url.(string),
		Password:   password.(string),
		DB:         0,
		MaxRetries: 3,
	})

	// 测试连接
	pong, err := MyRedis.Ping(context.Background()).Result()

	if err == redis.Nil {
		MyRedis.Close()
	} else if err != nil {
		fmt.Println("redis Error", err)
		MyRedis.Close()
	} else {
		fmt.Println("redis Ping", pong)
	}
}
