package rand

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/martinlindhe/base36"
	"goweb/pkg/util/base62"
	"strings"
	"time"
)

func BytesAsBase64String(bytes int) (string, error) {
	b, err := Bytes(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), err
}

func BytesAsBase36String(bytes int) (string, error) {
	b, err := Bytes(bytes)
	if err != nil {
		return "", err
	}
	return strings.ToLower(base36.EncodeBytes(b)), err
}

func BytesAsBase16String(bytes int) (string, error) {
	b, err := Bytes(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func KeyAsBase62() string {
	timePart := base62.Encode(time.Now().Unix())
	randPart := fillStringWithZero(base62.Encode(int64(Int31())), 3)
	return fmt.Sprintf("%s%s", timePart, randPart)
}

func KeyWithIDAsBase62(id int64) string {
	idPart := base62.Encode(id)
	timePart := base62.Encode(time.Now().Unix())
	randPart := fillStringWithZero(base62.Encode(int64(Int31())), 3)
	return fmt.Sprintf("%s%s%s", idPart, timePart, randPart)
}

func fillStringWithZero(str string, width int) string {
	l := len(str)
	if l >= width {
		start := l - width
		return str[start:]
	}

	var buf bytes.Buffer
	buf.WriteString(str)

	for i := 0; i < width-l; i++ {
		buf.WriteString("0")
	}

	return buf.String()
}
