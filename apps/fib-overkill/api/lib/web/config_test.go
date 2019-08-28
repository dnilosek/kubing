package web_test

import (
	"testing"

	. "github.com/dnilosek/kubing/apps/fib-overkill/api/lib/web"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig(80, "/api", "./webdir", "redis://localhost:8080", "postgres://localhost:8080")

	assert.Equal(t, 80, cfg.Port)
	assert.Equal(t, "/api", cfg.APIPath)
	assert.Equal(t, "redis://localhost:8080", cfg.RedisURL)
	assert.Equal(t, "postgres://localhost:8080", cfg.PostgresURL)
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	assert.Equal(t, 8080, cfg.Port)
	assert.Equal(t, "/", cfg.APIPath)
	assert.Equal(t, "redis://localhost:6379", cfg.RedisURL)
	assert.Equal(t, "redis://localhost:5432", cfg.PostgresURL)
}
