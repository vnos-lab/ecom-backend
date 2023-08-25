package cache

import (
	"erp/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Client struct {
	Client *redis.Client
}

func NewCacheClient(config *config.Config, logger *zap.Logger) *Client {
	client := &Client{}
	logger.Info("Connecting to redis...")

	client.InitializeConnection(config.Redis.Host, config.Redis.Password, config.Redis.PoolSize, config.Redis.MinIdleConns, config.Redis.DB)
	return client
}

func (c *Client) InitializeConnection(host string, passWord string, poolSize, minIdleConns, DB int) {

	c.Client = redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     passWord,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
		DB:           0,
	})
}

// Pipeline .
func (c *Client) Pipeline() redis.Pipeliner {
	return c.Client.Pipeline()
}
