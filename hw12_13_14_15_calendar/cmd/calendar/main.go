package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	internalconfig "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/config"
	internallogger "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
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
		log.Fatalf("Failed to create logger: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	storage, err := createStorage(ctx, *config)
	if err != nil {
		cancel()
		log.Fatalf("Failed to create storage: %s", err) //nolint:gocritic
	}

	calendar := app.New(logg, storage)
	server := internalhttp.NewServer(logg, calendar, config.HTTP.Host, config.HTTP.Port)

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
		cancel()
		log.Fatalf("Failed to start http server: %s", err)
	}
}

func createStorage(ctx context.Context, config internalconfig.Config) (app.Storage, error) {
	var storage app.Storage
	var err error
	switch config.Storage.Type {
	case internalconfig.InMemory:
		storage = memorystorage.New()
	case internalconfig.SQL:
		storage, err = sqlstorage.New(ctx, config.Storage.Dsn).Connect(ctx)
		if err != nil {
			return nil, err
		}
	default:
		log.Fatalf("Unknown storage type: %s\n", config.Storage.Type)
	}
	return storage, nil
}
