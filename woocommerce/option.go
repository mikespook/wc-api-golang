package woocommerce

import (
	"time"
)

type Option struct {
	API             bool
	APIPrefix       string
	Version         string
	Timeout         time.Duration
	VerifySSL       bool
	QueryStringAuth string
	OauthTimestamp  time.Time
}
