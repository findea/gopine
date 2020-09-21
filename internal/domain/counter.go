package domain

import (
	"context"
	"goweb/pkg/redis"
	"time"
)

type counterDomain struct {
}

var CounterDomain = new(counterDomain)

func (d *counterDomain) Count(_ context.Context, key, field string, duration time.Duration) (int64, error) {
	_, err := redis.HGet(key, field)
	if err != nil && err != redis.Nil {
		return 0, err
	}

	if err == redis.Nil {
		err = redis.HSetTTL(key, field, 1, duration)
		if err != nil {
			return 0, err
		}
		return 1, nil
	}

	return redis.HIncrBy(key, field, 1)
}
