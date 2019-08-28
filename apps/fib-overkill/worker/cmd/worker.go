package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dnilosek/kubing/apps/fib-overkill/lib/database"
	"github.com/dnilosek/kubing/apps/fib-overkill/worker/lib/web"
	"github.com/dnilosek/kubing/apps/fib-overkill/worker/lib/work"
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
	db, err := database.OpenRedis(cfg.DBURL)
	if err != nil {
		return err
	}
	listener := work.NewListener(db)

	msgChan := listener.Listen(cfg.InputChannel)
	for msg := range msgChan {
		i, err := strconv.Atoi(msg)
		if err != nil {
			log.Println("Cannot convert message to int:", err)
		}
		outVal := work.FibonacciNumberAtIndex(i)
		log.Printf("Recieved message %v , computed value %v\n", i, outVal)
		err = db.HSet(cfg.OutputChannel, msg, strconv.Itoa(outVal))
		if err != nil {
			return err
		}
	}
	return nil
}
