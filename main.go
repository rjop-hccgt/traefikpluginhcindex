// Package traefikpluginhcindex plugin.
package traefikpluginhcindex

import (
	"context"
	"log"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
}

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
	log.Printf("Got serve request: %v", req.URL.Path)
	req.URL.Path = req.URL.Path + "/index.html"
	if req.URL.RawPath != "" {
		req.URL.RawPath = req.URL.RawPath + "/index.html"
	}
	req.RequestURI = req.URL.RequestURI()
	a.next.ServeHTTP(rw, req)
}
