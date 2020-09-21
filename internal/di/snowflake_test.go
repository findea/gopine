package di

import (
	"github.com/stretchr/testify/assert"
	"goweb/pkg/log"
	"goweb/pkg/snowflake"
	"testing"
)

func TestSnowfkae(t *testing.T) {
	log.Info(snowflake.GenerateInt64Id())
	log.Info(snowflake.GenerateHex())
	assert.True(t, snowflake.GenerateInt64Id() > 0)
	assert.True(t, snowflake.GenerateInt64Id() != snowflake.GenerateInt64Id())
}