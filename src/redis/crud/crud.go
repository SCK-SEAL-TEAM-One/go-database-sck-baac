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

func (rc RedisClient) ReadAllKey() ([]map[string]string, error) {
	list := []map[string]string{}
	keys, err := rc.Client.Keys("*").Result()
	if err != nil {
		return list, err
	}

	for _, key := range keys {
		value, _ := rc.ReadKey(key)
		list = append(list, map[string]string{
			"id":          key,
			"description": value,
		})
	}
	return list, nil
}

func (rc RedisClient) ReadById(id string) (map[string]string, error) {
	output := map[string]string{}
	value, err := rc.ReadKey(id)
	if err != nil {
		return output, err
	}
	output["id"] = id
	output["description"] = value
	return output, nil
}
