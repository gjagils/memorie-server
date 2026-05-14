package immich

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNew_requiresBaseURLAndAPIKey(t *testing.T) {
	if _, err := New(Config{APIKey: "k"}); err == nil {
		t.Error("expected error when BaseURL is missing")
	}
	if _, err := New(Config{BaseURL: "http://x"}); err == nil {
		t.Error("expected error when APIKey is missing")
	}
	if _, err := New(Config{BaseURL: "not-a-url", APIKey: "k"}); err == nil {
		t.Error("expected error when BaseURL has no scheme/host")
	}
}

func TestHealth_success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/users/me" {
			t.Errorf("unexpected path %q", r.URL.Path)
		}
		if got := r.Header.Get("x-api-key"); got != "test-key" {
			t.Errorf("missing or wrong x-api-key header: %q", got)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"u1","email":"x@y"}`))
	}))
	defer srv.Close()

	c, err := New(Config{BaseURL: srv.URL, APIKey: "test-key"})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Health(context.Background()); err != nil {
		t.Fatalf("Health: %v", err)
	}
}

func TestHealth_trailingSlashBaseURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/users/me" {
			t.Errorf("unexpected path %q (trailing slash should be stripped)", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c, err := New(Config{BaseURL: srv.URL + "/", APIKey: "k"})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Health(context.Background()); err != nil {
		t.Fatalf("Health: %v", err)
	}
}

func TestHealth_unauthorized(t *testing.T) {
	for _, status := range []int{http.StatusUnauthorized, http.StatusForbidden} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(status)
		}))

		c, _ := New(Config{BaseURL: srv.URL, APIKey: "bad"})
		err := c.Health(context.Background())
		if err == nil || !strings.Contains(err.Error(), "API key invalid") {
			t.Errorf("status %d: expected 'API key invalid' error, got %v", status, err)
		}
		srv.Close()
	}
}

func TestHealth_unexpectedStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("oops"))
	}))
	defer srv.Close()

	c, _ := New(Config{BaseURL: srv.URL, APIKey: "k"})
	err := c.Health(context.Background())
	if err == nil || !strings.Contains(err.Error(), "unexpected status 500") {
		t.Fatalf("expected 'unexpected status 500' error, got %v", err)
	}
}

func TestHealth_contextCancelled(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c, _ := New(Config{BaseURL: srv.URL, APIKey: "k"})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := c.Health(ctx); err == nil {
		t.Fatal("expected error from cancelled context")
	}
}
