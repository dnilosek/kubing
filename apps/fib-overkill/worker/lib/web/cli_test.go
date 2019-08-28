package web_test

import (
	"testing"

	. "github.com/dnilosek/kubing/apps/fib-overkill/worker/lib/web"
	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {

	// Mock worker
	runWorker := func(*Config) error { return nil }

	cli := Cli(&CliMethods{
		RunWorker: runWorker,
	})

	assert.Equal(t, "fib-worker", cli.Name)
	assert.Equal(t, "fib-worker", cli.HelpName)
	assert.Equal(t, "Listen to a redis DB for values to compute fibonacci number on", cli.Usage)
	assert.Equal(t, "fib-worker --db-url DB_URL --in-chan MSG_CHANNEL --out-chan VAL_CHANNEL", cli.UsageText)
	assert.Equal(t, 0, len(cli.Commands))
	assert.Equal(t, 3, len(cli.Flags))
	assert.Equal(t, "db-url", cli.Flags[0].GetName())
	assert.Equal(t, "in-chan", cli.Flags[1].GetName())
	assert.Equal(t, "out-chan", cli.Flags[2].GetName())

}
