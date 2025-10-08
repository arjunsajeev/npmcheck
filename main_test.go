package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheck_Available(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	r := check(&http.Client{}, server.URL+"/", "test-package")

	if !r.available {
		t.Error("expected package to be available")
	}
	if r.err != nil {
		t.Errorf("unexpected error: %v", r.err)
	}
}

func TestCheck_NotAvailable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	r := check(&http.Client{}, server.URL+"/", "react")

	if r.available {
		t.Error("expected package to be unavailable")
	}
	if r.err != nil {
		t.Errorf("unexpected error: %v", r.err)
	}
}

func TestCheck_Error(t *testing.T) {
	r := check(&http.Client{}, "http://invalid/", "test")

	if r.err == nil {
		t.Error("expected error")
	}
}
