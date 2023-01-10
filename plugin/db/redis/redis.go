package redis

import (
	"github.com/ct-llper/go-pkg/utils"
	"github.com/go-redis/redis/v8"
	"strconv"
	"sync"
)

var (
	config map[string]map[string]string
	once   sync.Once
)

type Option struct {
	dialect string
}

// Config 初始化配置
func Config(c map[string]map[string]string) {
	once.Do(func() {
		//初始化配置
		config = c
	})
}

// client 建立连接
func client(args []string) (client *redis.Client) {
	var (
		db   int    = 0
		node string = "0"
	)
	if len(args) > 0 && !utils.Empty(args[0]) {
		node = args[0]
	}
	if len(args) > 1 && !utils.Empty(args[1]) {
		db, _ = strconv.Atoi(args[1])
	}
	return redis.NewClient(&redis.Options{
		Addr:     config[node]["addr"],
		Password: config[node]["password"],
		DB:       db,
	})
}

// GetRedis 获取client
func GetRedis(opts ...string) *redis.Client {
	args := make([]string, 0)
	for _, opt := range opts {
		args = append(args, opt)
	}
	return client(args)
}
