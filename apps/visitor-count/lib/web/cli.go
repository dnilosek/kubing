package web

import (
	"log"
	"strings"

	"github.com/urfave/cli"
)

const (
	envArg         = "env"
	portArg        = "port"
	apiPathArg     = "api-path"
	webDirArg      = "web-dir"
	databaseUrlArg = "db-url"
)

// Define operations that CLI impliments
type RunServer func(cfg *Config) error
type CliMethods struct {
	RunServer RunServer
}

// Create the CLI app
func Cli(methods *CliMethods) *cli.App {
	app := cli.NewApp()

	// Define the app parameters and flags
	app.Name = "visitor-counter"
	app.HelpName = "visitor-counter"
	app.Usage = "Run the visitor counter website server"
	app.UsageText = "visitor-counter -e ENV -p PORT --api-path API_PATH --web-dir WEB_DIR --db-url DB_URL"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   strings.Join([]string{envArg, "e"}, ","),
			Value:  defaultEnv,
			Usage:  "environment (dev | stage | prod)",
			EnvVar: "ENV",
		},
		cli.IntFlag{
			Name:   strings.Join([]string{portArg, "p"}, ","),
			Value:  defaultPort,
			Usage:  "port to listen on",
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   apiPathArg,
			Value:  defaultAPIPath,
			Usage:  "url path prefix for mounting API router",
			EnvVar: "API_PATH",
		},
		cli.StringFlag{
			Name:   webDirArg,
			Value:  defaultWebDir,
			Usage:  "path to local web assets (templates)",
			EnvVar: "WEB_DIR",
		},
		cli.StringFlag{
			Name:   databaseUrlArg,
			Value:  defaultDBURL,
			Usage:  "Connection URL to redis server",
			EnvVar: "DB_URL",
		},
	}

	// Create the action for the app
	app.Action = func(c *cli.Context) error {

		cfg := getConfig(c)

		log.Printf("ENV:	%s", cfg.Env)
		log.Printf("API PATH: 	%s", cfg.APIPath)
		log.Printf("WEB PATH:	%s", cfg.WebDir)
		log.Printf("DB URL:	%s", cfg.DBURL)

		return methods.RunServer(cfg)
	}
	return app
}

func getConfig(c *cli.Context) *Config {
	env := c.String(envArg)
	port := c.Int(portArg)
	apiPath := c.String(apiPathArg)
	webDir := c.String(webDirArg)
	dbURL := c.String(databaseUrlArg)
	return NewConfig(env, port, apiPath, webDir, dbURL)
}
