package domain

import (
	"goweb/internal/di"
	"goweb/pkg/config"
	"testing"
)

func TestMain(m *testing.M) {
	config.Path = config.ConfPath("example-test")
	di.Init()
	m.Run()
}
