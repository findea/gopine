package conf

import (
	"goweb/pkg/config"
	"goweb/pkg/lighttracer"
	"goweb/pkg/log"
	"goweb/pkg/mail"
	"goweb/pkg/mysql"
	"goweb/pkg/redis"
	"goweb/pkg/snowflake"
	"sync"
)

type Config struct {
	Log       *log.Config
	Redis     *redis.Config
	MySQL     map[string]*mysql.Config
	Mail      []*mail.Config
	JWT       *JWT
	Trace     *lighttracer.Config
	Server    *Server
	Snowflake *snowflake.Config
}

type JWT struct {
	sync.Mutex
	SecretKey   string
	secretBytes []byte
	ExpiresAt   config.Duration
}

func (j *JWT) SecretBytes() []byte {
	j.Lock()
	defer j.Unlock()
	if j.secretBytes == nil {
		j.secretBytes = []byte(j.SecretKey)
	}
	return j.secretBytes
}

type Server struct {
	WebAddress string
}
