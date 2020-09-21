package di

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"goweb/internal/conf"
	"goweb/pkg/redis"
	"goweb/pkg/util/rand"
	"goweb/pkg/util/strs"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	if conf.Conf.Redis == nil {
		return
	}

	err := redis.Set("score", 100, time.Hour)
	assert.Nil(t, err)

	val, err := redis.Get("score")
	assert.Nil(t, err)
	assert.Equal(t, 100, strs.StrToIntWithDefaultZero(val))

	val, err = redis.Get(fmt.Sprintf("%d", rand.Int63()))
	assert.Equal(t, err, redis.Nil)
	assert.Equal(t, val, "")
}
