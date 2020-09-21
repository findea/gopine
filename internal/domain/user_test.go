package domain

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"goweb/internal/model/dbmodel"
	"goweb/pkg/util/rand"
	"testing"
)

func TestUserAdd(t *testing.T) {
	user := NewDummyUser(t)

	userID, err := UserDomain.Insert(context.TODO(), user)
	assert.Nil(t, err)
	assert.Greater(t, userID, int64(0))

	userInDB, err := UserDomain.SelectOneByEmail(context.TODO(), user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.UserID, userInDB.UserID)
	t.Logf("%+v", userInDB)

	userInDB, err = UserDomain.SelectOneByUserID(context.TODO(), user.UserID)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userInDB.Email)
	t.Logf("%+v", userInDB)
}

func NewDummyUser(t *testing.T) *dbmodel.User {
	nickname := fmt.Sprintf("dummy%d", rand.Int31Range(100_000, 999_999))
	user := &dbmodel.User{
		Nickname:    nickname,
		Email:       fmt.Sprintf("%s@dummy.com", nickname),
		Password:    "dummy123",
		LastLoginIP: "222.222.222.222",
	}

	return user
}
