package models

import (
	"context"
	"time"
	"xiaomi_store/redis"
)

const CAPTCHA = "captcha:"

type RedisStore struct {
}

func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	return redis.RedisDB.Set(context.Background(), key, value, time.Minute*60*2).Err()
}

func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := redis.RedisDB.Get(context.Background(), key).Result()
	if err != nil {
		return ""
	}
	if clear {
		// 清除验证码
		redis.RedisDB.Del(context.Background(), key)
	}
	return val
}

func (r RedisStore) Verify(id string, value string, clear bool) bool {
	val := r.Get(id, clear)
	return val == value
}
