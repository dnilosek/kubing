package web

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartAndStop(t *testing.T) {
	cfg := DefaultConfig()

	// Change defaults to not conflict with other processes
	// and use correct template/public
	cfg.Port = 9999
	cfg.WebDir = "../../web"

	server := NewServer(cfg)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Start()
		require.NotNil(t, err)
		assert.Equal(t, "http: Server closed", err.Error())
	}()

	time.Sleep(1 * time.Second)

	err := server.Stop(context.Background())
	require.Nil(t, err)
	wg.Wait()
}

func TestIndex(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Change defaults to not conflict with other processes
	// and use correct template/public
	cfg := DefaultConfig()
	cfg.Port = 9999
	cfg.WebDir = "../../web"

	server := NewServer(cfg)

	server.router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}