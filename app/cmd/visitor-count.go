package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dnilosek/kubing/app/lib/database"
	"github.com/dnilosek/kubing/app/lib/web"
)

func main() {
	app := web.Cli(&web.CliMethods{
		RunServer: runServer,
	})

	err := app.Run(os.Args)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func runServer(cfg *web.Config) error {

	db, err := database.Open(cfg.DBURL)
	if err != nil {
		return err
	}
	server := web.NewServer(cfg, db)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		handleInterrupt(server)
	}()

	err = server.Start()
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func handleInterrupt(server *web.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Interrupted...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	server.Stop(ctx)
}
