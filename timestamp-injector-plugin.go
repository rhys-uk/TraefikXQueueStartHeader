package timestamp_injector

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

// Config holds plugin configuration.
type Config struct {
	HeaderName string `json:"headerName,omitempty"`
}

// CreateConfig creates default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		HeaderName: "X-Queue-Start", // default value optimized for APM queue time
	}
}

type TimestampInjector struct {
	name       string
	next       http.Handler
	headerName string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &TimestampInjector{
		name:       name,
		next:       next,
		headerName: config.HeaderName,
	}, nil
}

func (eh *TimestampInjector) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	msec := time.Now().UnixMilli()
	req.Header.Set(eh.headerName, strconv.FormatInt(msec, 10))
	eh.next.ServeHTTP(rw, req)
}
