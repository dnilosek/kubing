package web_test

import (
	"testing"

	. "github.com/dnilosek/kubing/apps/fib-overkill/worker/lib/web"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig("redis://localhost:8080")
	assert.Equal(t, "redis://localhost:8080", cfg.DBURL)
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	assert.Equal(t, "redis://localhost:6379", cfg.DBURL)
}
