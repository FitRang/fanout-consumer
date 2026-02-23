package redispub

import "github.com/redis/go-redis/v9"

func NewRedisClient(conn string) (*redis.Client, error) {
opt, err := redis.ParseURL(conn)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}
