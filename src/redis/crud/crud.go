package crud

import (
	"errors"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	Client *redis.Client
}

func (rc RedisClient) ReadKey(key string) (string, error) {
	return rc.Client.Get(key).Result()
}

func (rc RedisClient) WriteKey(key string, val interface{}) error {
	exist, _ := rc.Client.Exists(key).Result()
	if exist == 1 {
		return errors.New("Duplicate Key")
	}
	return rc.Client.Set(key, val, 0).Err()
}

func (rc RedisClient) DeleteKey(key string) (int64, error) {
	return rc.Client.Del(key).Result()
}
