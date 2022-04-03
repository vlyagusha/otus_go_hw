package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	internalconfig "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/config"
	internallogger "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/mq"
	transport "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/transport/log"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/sender_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	config, err := internalconfig.LoadSenderConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	logg, err := internallogger.New(config.Logger)
	if err != nil {
		log.Fatalf("Failed to create logger: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	notificationSource, err := mq.NewRabbit(ctx, config.Rabbit.Dsn, config.Rabbit.Exchange, config.Rabbit.Queue, logg)
	if err != nil {
		cancel()
		log.Fatalf("Failed to create NotificationSource (rabbit): %s", err) //nolint:gocritic
	}

	transports := []app.NotificationTransport{
		transport.NewLogNotificationTransport(logg),
	}

	sender := app.NewNotificationSender(notificationSource, logg, transports)
	sender.Run()

	<-ctx.Done()
}
