package domain

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"goweb/pkg/util/rand"
	"testing"
)

func TestSendVerifyEmailCode(t *testing.T) {
	email := fmt.Sprintf("dummy%d@dummy.com", rand.Int31())
	code := EmailCodeDomain.Rand4DigistCode(context.TODO())

	err := EmailCodeDomain.SendEmailCode(context.TODO(), email, code)
	assert.Nil(t, err)

	err = EmailCodeDomain.SendEmailCode(context.TODO(), email, code)
	assert.Nil(t, err)

	result, err := EmailCodeDomain.VerifyEmailCode(context.TODO(), email, code)
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = EmailCodeDomain.VerifyEmailCode(context.TODO(), email, "badcode")
	assert.False(t, result)

	result, err = EmailCodeDomain.VerifyEmailCode(context.TODO(), "bademail@dummy.com", code)
	assert.False(t, result)
}
