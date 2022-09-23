package httpclient

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

func newHTTPClient(cfg Config) *http.Client {
	return &http.Client{
		Timeout: cfg.GetHTTPTimeout(),
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   cfg.GetHTTPDialTimeout(),
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			DisableKeepAlives:     false,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          cfg.MaxIdleConns,
			MaxIdleConnsPerHost:   cfg.MaxIdleConnsPerHost,
			MaxConnsPerHost:       cfg.MaxConnsPerHost,
			IdleConnTimeout:       cfg.GetHTTPIdleConnTimeout(),
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

var (
	restyClient *resty.Client
	once        sync.Once
)

func GetRestyClient(cfg Config) *resty.Client {
	once.Do(func() {
		restyClient = resty.NewWithClient(newHTTPClient(cfg))
	})
	return restyClient
}
