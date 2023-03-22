package traefikplugin_hc_index

import (
	"context"
	"log"
	"net/http"
)

type Config struct {
}

func CreateConfig() *Config {
	return &Config{}
}

type HcIndex struct {
	next http.Handler
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HcIndex{
		next: next,
		name: name,
	}, nil
}

func (a *HcIndex) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Got serve request: %v", req.URL.Path)
	
	a.next.ServeHTTP(rw, req)
}
