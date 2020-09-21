package domain

import (
	"context"
	"github.com/stretchr/testify/assert"
	"goweb/internal/model/auth"
	"goweb/pkg/util/rand"
	"testing"
	"time"
)

var JwtSecretKey = []byte("5954f23d949108fd79e510283246ee39")

func TestCreateToken(t *testing.T) {
	user := makeDummyOrgUserClaims()
	token, _ := JWTDomain.CreateToken(context.TODO(), JwtSecretKey, user, time.Now().Add(time.Hour*72).Unix())
	t.Log(token)

	claims, err := JWTDomain.ParseToken(context.TODO(), JwtSecretKey, token)
	assert.Nil(t, err)
	assert.Equal(t, user, claims)

	t.Log("claims:", claims)
}

func TestTokenExpired(t *testing.T) {
	user := makeDummyOrgUserClaims()
	token, _ := JWTDomain.CreateToken(context.TODO(), JwtSecretKey, user, time.Now().Add(time.Second).Unix())
	t.Log(token)

	time.Sleep(time.Second * time.Duration(2))
	_, err := JWTDomain.ParseToken(context.TODO(), JwtSecretKey, token)
	assert.NotNil(t, err)
	t.Log(err)
}

func TestWrongToken(t *testing.T) {
	token := "wrongtoken"
	_, err := JWTDomain.ParseToken(context.TODO(), JwtSecretKey, token)
	assert.NotNil(t, err)
	t.Log(err)
}

func TestGenerateJWTToken(t *testing.T) {
	claims := &auth.UserClaims{
		UserID: rand.Int63(),
	}

	token, err := JWTDomain.GenerateJWTToken(context.TODO(), claims)
	assert.Nil(t, err)
	t.Log(token)

	result, err := JWTDomain.VerifyJWTToken(context.TODO(), token)
	assert.Equal(t, claims, result)
	t.Log(result.UserID)
}

func makeDummyOrgUserClaims() *auth.UserClaims {
	return &auth.UserClaims{
		UserID: 767134721,
	}
}
