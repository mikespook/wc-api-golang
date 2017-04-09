package woocommerce

import (
	"time"
)

type Options struct {
	API             bool
	APIPrefix       string
	Version         string
	Timeout         time.Duration
	VerifySSL       bool
	QueryStringAuth string
	OauthTimestamp  time.Time
}
