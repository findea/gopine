package domain

import (
	"context"
	"fmt"
	"goweb/internal/model/errs"
	"goweb/internal/model/redismodel"
	"goweb/pkg/mail"
	"goweb/pkg/redis"
	"goweb/pkg/util/json"
	"goweb/pkg/util/rand"
	"time"
)

type emailCodeDomain struct {
}

var EmailCodeDomain = new(emailCodeDomain)

/**
get emailcode:email from redis
if exist and time < 50s return
set emailcode:email => time, code to redis
send email with code
*/
func (d *emailCodeDomain) SendEmailCode(_ context.Context, email, code string) error {
	val, err := d.emailCodeFromRedis(email)
	if err != nil {
		return err
	}

	remain := val.Time + 50 - time.Now().Unix()
	if remain > 0 {
		return nil
	}

	val.Time = time.Now().Unix()
	val.Code = code
	err = redis.Set(d.emailCodeRedisKey(email), json.ToJsonIgnoreError(val), 30*time.Minute)
	if err != nil {
		return errs.ServerError(err)
	}

	mail.SendMail(email,
		fmt.Sprintf("验证码：%s", val.Code),
		fmt.Sprintf("您的验证码是：%s（请尽快使用，后面会过期）", val.Code),
	)

	return nil
}

func (d *emailCodeDomain) VerifyEmailCode(_ context.Context, email, code string) (bool, error) {
	val, err := d.emailCodeFromRedis(email)
	if err != nil {
		return false, err
	}

	if val.Code == "" {
		return false, nil
	}

	return val.Code == code, nil
}

func (d *emailCodeDomain) Rand4DigistCode(_ context.Context) string {
	return fmt.Sprintf("%d", rand.Int31Range(1000, 9999))
}

func (d *emailCodeDomain) emailCodeFromRedis(email string) (*redismodel.EmailCode, error) {
	redisKey := d.emailCodeRedisKey(email)
	str, err := redis.Get(redisKey)
	if err != nil && err != redis.Nil {
		return nil, errs.ServerError(err)
	}

	var val = new(redismodel.EmailCode)
	if str != "" {
		err = json.FromJson(str, val)
		if err != nil {
			return nil, errs.ServerError(err)
		}
	}

	return val, nil
}

func (d *emailCodeDomain) emailCodeRedisKey(email string) string {
	return fmt.Sprintf("emailcode:%s", email)
}
