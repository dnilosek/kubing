package web

import (
	"log"

	"github.com/urfave/cli"
)

const (
	databaseUrlArg = "db-url"
)

// Define operations for CLI
type RunWorker func(cfg *Config) error
type CliMethods struct {
	RunWorker RunWorker
}

// Create the CLI app
func Cli(methods *CliMethods) *cli.App {

	// Define CLI input
	app := cli.NewApp()
	app.Name = "fib-worker"
	app.HelpName = "fib-worker"
	app.Usage = "Listen to a redis DB for values to compute fibonacci number on"
	app.UsageText = "fib-worker --db-url DB_URL"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   databaseUrlArg,
			Value:  defaultDBURL,
			Usage:  "Connection URL to redis server",
			EnvVar: "DB_URL",
		},
	}

	// Create the action for the app
	app.Action = func(c *cli.Context) error {

		// Startup logging
		cfg := getConfig(c)
		log.Printf("DB URL:	%s", cfg.DBURL)

		return methods.RunWorker(cfg)
	}
	return app
}

// Function go create our config
func getConfig(c *cli.Context) *Config {
	dbURL := c.String(databaseUrlArg)
	return NewConfig(dbURL)
}
