package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dnilosek/kubing/apps/fib-overkill/worker/lib/web"
	"github.com/dnilosek/kubing/apps/fib-overkill/worker/lib/work"
	"github.com/dnilosek/kubing/apps/visitor-count/lib/database"
)

func main() {
	app := web.Cli(&web.CliMethods{
		RunWorker: runWorker,
	})

	err := app.Run(os.Args)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func runWorker(cfg *web.Config) error {
	db, err := database.Open(cfg.DBURL)
	if err != nil {
		return err
	}
	listener := work.NewListener(db)

	msgChan := listener.Listen("test")
	for msg := range msgChan {
		fmt.Println(msg)
	}
	return nil
}
