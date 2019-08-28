package web_test

import (
	"testing"

	. "github.com/dnilosek/kubing/apps/fib-overkill/api/lib/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCli(t *testing.T) {
	// Mock server
	runServer := func(*Config) error { return nil }

	cli := Cli(&CliMethods{
		RunServer: runServer,
	})

	assert.Equal(t, "fib-api", cli.Name)
	assert.Equal(t, "fib-api", cli.HelpName)
	assert.Equal(t, "Run the fibonacci calculator api", cli.Usage)
	assert.Equal(t, "fib-api -p PORT --api-path API_PATH --redis-url REDIS_URL --postgres-url POSTGRES_URL", cli.UsageText)
	assert.Equal(t, 0, len(cli.Commands))
	assert.Equal(t, 4, len(cli.Flags))
	assert.Equal(t, "port,p", cli.Flags[0].GetName())
	assert.Equal(t, "api-path", cli.Flags[1].GetName())
	assert.Equal(t, "redis-url", cli.Flags[2].GetName())
	assert.Equal(t, "postgres-url", cli.Flags[3].GetName())
}

func TestCliAction(t *testing.T) {

	var result *Config
	runServer := func(cfg *Config) error { result = cfg; return nil }

	cli := Cli(&CliMethods{
		RunServer: runServer,
	})

	// Test defaults
	err := cli.Run([]string{"app"})
	require.Nil(t, err)
	assert.Equal(t, 8080, result.Port)
	assert.Equal(t, "/", result.APIPath)
	assert.Equal(t, "redis://localhost:6379", result.RedisURL)
	assert.Equal(t, "postgres://localhost:5432", result.PostgresURL)

	// Test input
	err = cli.Run([]string{"app", "--port=80", "--api-path=/api", "--redis-url=redis://localhost:8080", "--postgres-url=postgres://localhost:8080"})
	require.Nil(t, err)
	assert.Equal(t, 80, result.Port)
	assert.Equal(t, "/api", result.APIPath)
	assert.Equal(t, "redis://localhost:8080", result.RedisURL)
	assert.Equal(t, "postgres://localhost:8080", result.PostgresURL)
}
