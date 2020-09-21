package domain

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"goweb/pkg/util/rand"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	key := fmt.Sprintf("ip:192.168.%d.%d", rand.Int31Range(0, 255), rand.Int31Range(0, 255))

	for i := 1; i <= 10; i++ {
		count, err := CounterDomain.Count(context.TODO(), key, "/api", time.Minute*10)
		assert.Nil(t, err)
		assert.Equal(t, count, int64(i))
	}
}
