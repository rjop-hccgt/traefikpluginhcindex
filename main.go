// Package traefikpluginhcindex plugin.
package traefikpluginhcindex

import (
	"context"
	"net/http"
	"strings"
)

// Config the plugin configuration.
type Config struct{}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// HcIndex plugin.
type HcIndex struct {
	next http.Handler
	name string
}

// New created a new HcIndex plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HcIndex{
		next: next,
		name: name,
	}, nil
}

func (a *HcIndex) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	pos := strings.LastIndex(req.URL.Path, ".")
	if pos == -1 {
		if strings.HasSuffix(req.URL.Path, "/") {
			req.URL.Path += "index.html"
		} else {
			req.URL.Path += "/index.html"
		}

		if req.URL.RawPath != "" {
			req.URL.RawPath = req.URL.Path
		}
	}
	req.RequestURI = req.URL.RequestURI()
	a.next.ServeHTTP(rw, req)
}
