package base62

import (
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	str := Encode(time.Now().Unix())
	t.Log(str, Decode(str))
}
