package bootstrap

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"todo_list/global"
)

// 初始化redis

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
		Password: global.App.Config.Redis.Password, // 未设置密码
		DB:       global.App.Config.Redis.DB,       //使用默认数据库
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed,err:", zap.Any("err", err))
	}
	fmt.Println("client", client)
	return client
}
