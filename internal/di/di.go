package di

import (
	"goweb/internal/conf"
	"goweb/pkg/config"
	"goweb/pkg/lighttracer"
	"goweb/pkg/log"
	"goweb/pkg/mail"
	"goweb/pkg/mysql"
	"goweb/pkg/redis"
	"goweb/pkg/snowflake"
)

func Init() error {
	err := config.Init(conf.Conf)
	if err != nil {
		return err
	}

	err = InitWithConf(conf.Conf)
	if err != nil {
		return err
	}

	log.Infof("%s loaded", config.Path)
	return nil
}

func InitWithConf(c *conf.Config) error {
	if c.Log != nil {
		err := log.Init(c.Log)
		if err != nil {
			return err
		}
	}

	if c.Redis != nil {
		err := redis.Init(c.Redis)
		if err != nil {
			return err
		}
	}

	for key, conf := range c.MySQL {
		err := mysql.InitDb(key, conf)
		if err != nil {
			return err
		}
	}

	if c.Mail != nil {
		mail.Init(c.Mail)
	}

	if c.Trace != nil {
		err := lighttracer.Init(c.Trace)
		if err != nil {
			return err
		}
	}

	err := snowflake.Init(c.Snowflake)
	if err != nil {
		return err
	}

	return nil
}
