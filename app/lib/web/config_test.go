package web_test

import (
	"testing"

	. "github.com/dnilosek/kubing/app/lib/web"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig("prod", 80, "/api", "./webdir", "redis://localhost:8080")

	assert.Equal(t, "prod", cfg.Env)
	assert.Equal(t, 80, cfg.Port)
	assert.Equal(t, "/api", cfg.APIPath)
	assert.Equal(t, "./webdir", cfg.WebDir)
	assert.Equal(t, "redis://localhost:8080", cfg.DBURL)
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	assert.Equal(t, "dev", cfg.Env)
	assert.Equal(t, 8080, cfg.Port)
	assert.Equal(t, "/", cfg.APIPath)
	assert.Equal(t, "./web", cfg.WebDir)
	assert.Equal(t, "redis://localhost:6379", cfg.DBURL)
}
