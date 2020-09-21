package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"goweb/internal/domain"
	"goweb/internal/model/auth"
	"goweb/internal/model/dbmodel"
	"goweb/internal/model/request"
	"goweb/pkg/util/rand"
	"testing"
)

func TestUserRegister(t *testing.T) {
	// 正常注册
	DummyUserRegister(t)

	// 参数格式错误
	registerReq := new(request.UserRegisterReq)
	registerReq.Nickname = ""
	registerReq.Code = "1234"
	registerReq.Email = "test"
	registerReq.Password = "123"
	registerReq.LastLoginIP = "222.222.222.222"

	_, err := UserService.Register(context.TODO(), registerReq)
	assert.NotNil(t, err)
	t.Log(err)
}

func TestUserLogin(t *testing.T) {
	// 正常注册
	claims, user := DummyUserRegister(t)

	// 正常登录
	userLoginReq := &request.UserLoginReq{
		Email:    claims.Email,
		Password: user.Password,
	}
	_, err := UserService.Login(context.TODO(), userLoginReq)
	assert.Nil(t, err)

	//密码不正确
	userLoginReq.Password = "123123123"
	_, err = UserService.Login(context.TODO(), userLoginReq)
	assert.NotNil(t, err)
	t.Log(err)

	//邮箱未注册
	userLoginReq.Email = fmt.Sprintf("%d@dummy.com", rand.Int63())
	_, err = UserService.Login(context.TODO(), userLoginReq)
	assert.NotNil(t, err)
	t.Log(err)
}

func DummyUserRegister(t *testing.T) (claims *auth.UserClaims, user *dbmodel.User) {
	password, err := rand.BytesAsBase36String(10)
	assert.Nil(t, err)

	registerReq := new(request.UserRegisterReq)
	registerReq.Nickname = fmt.Sprintf("dummy%d", rand.Int31Range(100_000, 999_999))
	registerReq.Code = domain.EmailCodeDomain.Rand4DigistCode(context.TODO())
	registerReq.Email = fmt.Sprintf("%s@dummy.com", registerReq.Nickname)
	registerReq.Password = password
	registerReq.LastLoginIP = "222.222.222.222"

	err = domain.EmailCodeDomain.SendEmailCode(context.TODO(), registerReq.Email, registerReq.Code)
	assert.Nil(t, err)

	_, err = UserService.Register(context.TODO(), registerReq)
	assert.Nil(t, err)

	user, err = domain.UserDomain.SelectOneByEmail(context.TODO(), registerReq.Email)
	assert.Nil(t, err)

	claims = &auth.UserClaims{
		UserID: user.UserID,
		Email:  user.Email,
	}

	return
}
