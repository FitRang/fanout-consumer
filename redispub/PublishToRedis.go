package redispub

import (
	"context"
	"encoding/json"

	"github.com/FitRang/fanout-consumer/model"
	"github.com/redis/go-redis/v9"
)

func PublishToRedis(rdb *redis.Client, msg []byte) error {
	var m model.BMessage
	if err := json.Unmarshal(msg, &m); err != nil {
		return err
	}

	channel := "user:" + m.Receiver.Email

	return rdb.Publish(
		context.Background(),
		channel,
		msg,
	).Err()
}
