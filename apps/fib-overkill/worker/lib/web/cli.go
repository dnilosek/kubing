package web

import (
	"log"

	"github.com/urfave/cli"
)

const (
	databaseUrlArg   = "db-url"
	outputChannelArg = "out-chan"
	inputChannelArg  = "in-chan"
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
	app.UsageText = "fib-worker --db-url DB_URL --in-chan MSG_CHANNEL --out-chan VAL_CHANNEL"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   databaseUrlArg,
			Value:  defaultDBURL,
			Usage:  "Connection URL to redis server",
			EnvVar: "DB_URL",
		},
		cli.StringFlag{
			Name:   inputChannelArg,
			Value:  defaultInputChannel,
			Usage:  "Input channel to listen on for messages",
			EnvVar: "MSG_CHANNEL",
		},
		cli.StringFlag{
			Name:   outputChannelArg,
			Value:  defaultOutputChannel,
			Usage:  "Output channel to put pass values back on",
			EnvVar: "VAL_CHANNEL",
		},
	}

	// Create the action for the app
	app.Action = func(c *cli.Context) error {

		// Startup logging
		cfg := getConfig(c)
		log.Printf("DB URL:		%s", cfg.DBURL)
		log.Printf("MSG CHANNEL:	%s", cfg.InputChannel)
		log.Printf("VAL CHANNEL:	%s", cfg.OutputChannel)

		return methods.RunWorker(cfg)
	}
	return app
}

// Function go create our config
func getConfig(c *cli.Context) *Config {
	dbURL := c.String(databaseUrlArg)
	inChan := c.String(inputChannelArg)
	outChan := c.String(outputChannelArg)
	return NewConfig(dbURL, inChan, outChan)
}
