package di

import (
	"testing"
)

func TestMain(m *testing.M) {
	err := Init()
	if err != nil {
		panic(err)
	}

	m.Run()
}
