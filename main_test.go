package traefikplugin_hc_index_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	traefikplugin_hc_index "github.com/rjop-hccgt/traefik-plugin"
)

func TestHcIndex(t *testing.T) {
	cfg := traefikplugin_hc_index.CreateConfig()
	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
	handler, err := traefikplugin_hc_index.New(ctx, next, cfg, "hc-index-test")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, req)

	log.Printf("Got header: %v", req.URL.Path)

}
