package redis

import (
	"fmt"

	"github.com/Thewalkers2012/BlogBackend/settings"
	"github.com/go-redis/redis"
)

// 声明一个全局的 rdb 变量
var (
	client *redis.Client
	Nil    = redis.Nil
)

// 初始化链接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err = client.Ping().Result()

	return err
}

func Close() {
	_ = client.Close()
}
