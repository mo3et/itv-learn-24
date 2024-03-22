package data

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

// var (
// 	service *redisService
// 	once    sync.Once
// )

type RedisService struct {
	Client *redis.Client
	Logger *log.Helper
	// *log.Helper
}

var (
	instance *RedisService
	once     sync.Once
)

func NewRedisService(logs log.Logger) (*RedisService, error) {
	once.Do(func() {
		// create Redis Client connect
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:26379", // Redis 服务器地址
			Password: "",                // Redis 服务器密码
			DB:       0,                 // 使用默认数据库
		})

		// Test connect heartbeat
		pong, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Error(err)
			return
		} else {
			log.Info("Success ping ", pong)
		}

		// 初始化 RedisService 实例
		instance = &RedisService{
			Client: client,
			Logger: log.NewHelper(logs),
		}
	})
	return instance, nil
}

// func Getclient(logs log.Logger) *redis.Client {
// 	rds, err := NewRedisService(logs)
// 	fmt.Println("Get Client failed.", err)
// 	return rds.client
// }

func (rs *RedisService) Close() error {
	// close Redis client connect
	return rs.Client.Close()
}

func (rs *RedisService) SetPlayerUsername(playerId int, username string) error {
	// 构建key name
	key := fmt.Sprintf("player:%d:username", playerId)

	exists, err := rs.Client.Exists(context.Background(), key).Result()
	if err != nil {
		rs.Logger.Errorf("Failed to check if key exists: %v", err)
		return err
	}
	// set player data
	// 命名空间用:来分隔, 会进行分组(或者理解为分成表), 例如 Player:id:name
	if exists == 0 {
		// err = rs.Client.Set(context.Background(), fmt.Sprintf("player:%d:username", playerId), username, 10*time.Hour).Err()
		// nx 如果exist 则不执行set 并返回错误
		err = rs.Client.SetNX(context.Background(), fmt.Sprintf("player:%d:username", playerId), username, 10*time.Hour).Err()
		if err != nil {
			rs.Logger.Errorf("Failed to set player username: %v", err)
			return err
		}
	} else {
		rs.Logger.Info("ID:", key, username, " is Exists!\n")
	}
	return nil
}

func (rs *RedisService) GetPlayerUsername(playerId int) (string, error) {
	// get player data
	username, err := rs.Client.Get(context.Background(), fmt.Sprintf("player:%d:username", playerId)).Result()
	if err != nil {
		rs.Logger.Errorf("Failed to get player username: %v", err)
	}
	return username, err
}
