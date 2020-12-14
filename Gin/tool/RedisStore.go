package tool

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"log"
	"time"
)

type RedisStore struct {
	Client *redis.Client
}

// 初始化RedisStore
func InitRedisStore() *RedisStore {
	redisConfig := GetConfig().RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
	store := RedisStore{client}
	base64Captcha.SetCustomStore(&store)
	return &store
}

/**
这里实现的是Set方法
*/
func (s *RedisStore) Set(id string, value string) {
	ctx := context.Background()
	err := s.Client.Set(ctx, id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

/**
这个库是直接集成了Redis,这里实现的是GET方法
*/
func (s *RedisStore) Get(id string, clear bool) (value string) {
	ctx := context.Background()
	val, err := s.Client.Get(ctx, id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := s.Client.Del(ctx, id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}
