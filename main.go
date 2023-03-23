// Package traefikpluginhcindex plugin.
package traefikpluginhcindex

import (
	"context"
	"fmt"
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
	fmt.Printf(">>>>> req: %v  <<<<<<< ", req)

	if strings.HasSuffix(req.URL.Path, "/") {
		req.URL.Path += "index.html"
	} else {
		req.URL.Path += "/index.html"
	}

	if req.URL.RawPath != "" {
		if strings.HasSuffix(req.URL.RawPath, "/") {
			req.URL.RawPath += "index.html"
		} else {
			req.URL.RawPath += "/index.html"
		}
	}
	req.RequestURI = req.URL.RequestURI()
	a.next.ServeHTTP(rw, req)
}
