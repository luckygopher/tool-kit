package httpclient

import "time"

const (
	DefaultTimeout         = time.Second * 30
	DefaultDialTimeout     = time.Second * 30
	DefaultIDleConnTimeout = time.Second * 90
)
