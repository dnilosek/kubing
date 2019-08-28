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
	redisUrlArg    = "redis-url"
	postgresUrlArg = "postgres-url"
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
	app.Name = "fib-api"
	app.HelpName = "fib-api"
	app.Usage = "Run the fibonacci calculator api"
	app.UsageText = "fib-api -p PORT --api-path API_PATH --redis-url REDIS_URL --postgres-url POSTGRES_URL"
	app.Flags = []cli.Flag{
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
			Name:   redisUrlArg,
			Value:  defaultRedisURL,
			Usage:  "Connection URL to redis server",
			EnvVar: "REDIS_URL",
		},
		cli.StringFlag{
			Name:   postgresUrlArg,
			Value:  defaultPostgresURL,
			Usage:  "Connection URL to postgres server",
			EnvVar: "POSTGRES_URL",
		},
	}

	// Create the action for the app
	app.Action = func(c *cli.Context) error {

		cfg := getConfig(c)

		log.Printf("API PATH: 		%s", cfg.APIPath)
		log.Printf("REDIS URL:		%s", cfg.RedisURL)
		log.Printf("POSTGRES URL:	%s", cfg.PostgresURL)

		return methods.RunServer(cfg)
	}
	return app
}

func getConfig(c *cli.Context) *Config {
	port := c.Int(portArg)
	apiPath := c.String(apiPathArg)
	redisURL := c.String(redisUrlArg)
	postgresURL := c.String(postgresUrlArg)

	return NewConfig(port, apiPath, redisURL, postgresURL)
}
