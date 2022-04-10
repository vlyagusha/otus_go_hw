package factory

import (
	"context"
	"log"

	"github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	internalconfig "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/config"
	memorystorage "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
)

func CreateStorage(ctx context.Context, config internalconfig.StorageConf) (app.Storage, error) {
	var storage app.Storage
	var err error
	switch config.Type {
	case internalconfig.InMemory:
		storage = memorystorage.New()
	case internalconfig.SQL:
		storage, err = sqlstorage.New(ctx, config.Dsn).Connect(ctx)
		if err != nil {
			return nil, err
		}
	default:
		log.Fatalf("Unknown storage type: %s\n", config.Type)
	}
	return storage, nil
}
