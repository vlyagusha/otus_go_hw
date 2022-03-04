package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	internalconfig "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/config"
	internallogger "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/storage/memory"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := internalconfig.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	logg, err := internallogger.New(config.Logger)
	if err != nil {
		logg.Error("failed to create logger: " + err.Error())
		os.Exit(1)
	}

	storage := memorystorage.New()
	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(logg, calendar, config.HTTP.Host, config.HTTP.Port)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
