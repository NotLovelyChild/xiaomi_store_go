package redis

import (
	"fmt"
	"os"
	"strconv"
	"xiaomi_store/config"

	"github.com/go-redis/redis/v8"
)

var (
	RedisDB *redis.Client
)

func init() {
	db, err := strconv.ParseInt(config.Config.Section("redis").Key("db").String(), 10, 64)
	if err != nil {
		fmt.Printf("Fail to DB ParseInt: %v", err)
		os.Exit(1)
	}
	addr := config.Config.Section("redis").Key("host").String() + ":" + config.Config.Section("redis").Key("port").String()
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Config.Section("redis").Key("password").String(),
		DB:       int(db),
	})
}
