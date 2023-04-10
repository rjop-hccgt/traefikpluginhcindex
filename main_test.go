package traefikpluginhcindex_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	traefikplugin_hc_index "github.com/rjop-hccgt/traefikpluginhcindex"
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, req)
	assertIndex(t, req)

	req, err = http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/folder1", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, req)
	assertIndex(t, req)
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/folder1/style.css", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, req)
	assertNonIndex(t, req)
}

func assertIndex(t *testing.T, req *http.Request) {
	t.Helper()
	log.Printf("Validating %v", req.URL.Path)
	if !strings.HasSuffix(req.URL.Path, "/index.html") {
		t.Errorf("invalid path value: %s", req.URL.Path)
	}
}

func assertNonIndex(t *testing.T, req *http.Request) {
	t.Helper()
	log.Printf("Validating %v", req.URL.Path)
	if strings.HasSuffix(req.URL.Path, "/index.html") {
		t.Errorf("invalid path value: %s", req.URL.Path)
	}
}
