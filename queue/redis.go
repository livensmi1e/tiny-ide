package queue

import (
	"context"

	"github.com/livensmi1e/tiny-ide/pkg/config"
	"github.com/livensmi1e/tiny-ide/pkg/domain"
	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
	key    string
}

func New(cfg *config.Config, key string) *RedisQueue {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Queue.Addr,
		Password: cfg.Queue.Password,
		DB:       0,
	})
	return &RedisQueue{client: rdb, key: key}
}

func (r *RedisQueue) Push(submission *domain.Submission) error {
	data, err := submission.Serialize()
	if err != nil {
		return err
	}
	return r.client.LPush(context.Background(), r.key, data).Err()
}

func (r *RedisQueue) Pop() (*domain.Submission, error) {
	data, err := r.client.RPop(context.Background(), r.key).Result()
	if err != nil {
		return nil, err
	}
	return domain.Deserialize(data)
}
