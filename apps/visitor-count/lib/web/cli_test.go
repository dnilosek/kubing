package web_test

import (
	"testing"

	. "github.com/dnilosek/kubing/apps/visitor-count/lib/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCli(t *testing.T) {
	// Mock server
	runServer := func(*Config) error { return nil }

	cli := Cli(&CliMethods{
		RunServer: runServer,
	})

	assert.Equal(t, "visitor-counter", cli.Name)
	assert.Equal(t, "visitor-counter", cli.HelpName)
	assert.Equal(t, "Run the visitor counter website server", cli.Usage)
	assert.Equal(t, "visitor-counter -e ENV -p PORT --api-path API_PATH --web-dir WEB_DIR --db-url DB_URL", cli.UsageText)
	assert.Equal(t, 0, len(cli.Commands))
	assert.Equal(t, 5, len(cli.Flags))
	assert.Equal(t, "env,e", cli.Flags[0].GetName())
	assert.Equal(t, "port,p", cli.Flags[1].GetName())
	assert.Equal(t, "api-path", cli.Flags[2].GetName())
	assert.Equal(t, "web-dir", cli.Flags[3].GetName())
	assert.Equal(t, "db-url", cli.Flags[4].GetName())
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
	assert.Equal(t, "dev", result.Env)
	assert.Equal(t, 8080, result.Port)
	assert.Equal(t, "/", result.APIPath)
	assert.Equal(t, "./web", result.WebDir)
	assert.Equal(t, "redis://localhost:6379", result.DBURL)

	// Test input
	err = cli.Run([]string{"app", "--env=prod", "--port=80", "--api-path=/api", "--web-dir=/web", "--db-url=redis://localhost:8080"})
	require.Nil(t, err)
	assert.Equal(t, "prod", result.Env)
	assert.Equal(t, 80, result.Port)
	assert.Equal(t, "/api", result.APIPath)
	assert.Equal(t, "/web", result.WebDir)
	assert.Equal(t, "redis://localhost:8080", result.DBURL)
}
