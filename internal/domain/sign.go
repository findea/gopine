package domain

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"github.com/martinlindhe/base36"
	"goweb/internal/conf"
	"strings"
)

type signDomain struct {
}

var SignDomain = new(signDomain)

func (d *signDomain) Sign(_ context.Context, content string) string {
	h := hmac.New(sha256.New, conf.Conf.JWT.SecretBytes())
	h.Write([]byte(content))
	return strings.ToLower(base36.EncodeBytes(h.Sum(nil)))
}

func (d *signDomain) VerifySign(c context.Context, content, sign string) bool {
	return d.Sign(c, content) == strings.ToLower(sign)
}
