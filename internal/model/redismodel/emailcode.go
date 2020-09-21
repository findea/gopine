package redismodel

type EmailCode struct {
	Time int64  `json:"time"`
	Code string `json:"code"`
}
