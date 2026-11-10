package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yuntuCTO/yuntu-wechat-ipad/internal/config"
)

func TestHealthz(t *testing.T) {
	handler := NewHandler(config.Config{Env: config.EnvDev})

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("期望状态码 %d，实际 %d", http.StatusOK, rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Fatalf("期望 JSON Content-Type，实际 %q", ct)
	}
}
