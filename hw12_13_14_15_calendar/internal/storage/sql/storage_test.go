package sqlstorage

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	sqlstorage "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"gopkg.in/yaml.v3"
)

const DefaultConfigFile = "configs/config.yaml"

func TestStorage(t *testing.T) {
	if _, err := os.Stat(DefaultConfigFile); errors.Is(err, os.ErrNotExist) {
		t.Skip(DefaultConfigFile + " file does not exists")
	}

	configContent, _ := os.ReadFile(DefaultConfigFile)
	var config struct {
		Storage struct {
			Dsn string
		}
	}

	err := yaml.Unmarshal(configContent, config)
	if err != nil {
		t.Fatal("Failed to unmarshal config", err)
	}

	ctx := context.Background()
	storage := New(ctx, config.Storage.Dsn)
	if err := storage.Connect(ctx); err != nil {
		t.Fatal("Failed to connect to DB server", err)
	}

	t.Run("test SQLStorage CRUDL", func(t *testing.T) {
		tx, err := storage.conn.BeginTx(ctx, pgx.TxOptions{
			IsoLevel:       pgx.Serializable,
			AccessMode:     pgx.ReadWrite,
			DeferrableMode: pgx.NotDeferrable,
		})
		if err != nil {
			t.Fatal("Failed to connect to DB server", err)
		}

		userID := uuid.New()
		startedAt, err := time.Parse("2006-01-02 15:04:05", "2022-03-08 12:00:00")
		if err != nil {
			t.FailNow()
			return
		}
		finishedAt, err := time.Parse("2006-01-02 15:04:05", "2022-03-09 12:00:00")
		if err != nil {
			t.FailNow()
			return
		}
		notifyAt, err := time.Parse("2006-01-02 15:04:05", "2022-03-07 12:00:00")
		if err != nil {
			t.FailNow()
			return
		}

		event := sqlstorage.NewEvent(
			"Event title",
			startedAt,
			finishedAt,
			"Event description",
			userID,
			notifyAt,
		)

		err = storage.Create(*event)
		if err != nil {
			t.FailNow()
			return
		}

		saved, err := storage.FindAll()
		if err != nil {
			t.FailNow()
			return
		}
		require.Len(t, saved, 1)
		require.Equal(t, *event, saved[0])

		event.Title = "New event title"
		event.Description = "New event description"

		saved, err = storage.FindAll()
		if err != nil {
			t.FailNow()
			return
		}
		require.Len(t, saved, 1)
		require.NotEqual(t, *event, saved[0])
		require.NotEqual(t, event.Title, saved[0].Title)
		require.NotEqual(t, event.Description, saved[0].Description)

		err = storage.Update(*event)
		if err != nil {
			t.FailNow()
			return
		}

		saved, err = storage.FindAll()
		if err != nil {
			t.FailNow()
			return
		}
		require.Len(t, saved, 1)
		require.Equal(t, *event, saved[0])

		err = storage.Delete(event.ID)
		if err != nil {
			t.FailNow()
			return
		}

		saved, err = storage.FindAll()
		if err != nil {
			t.FailNow()
			return
		}
		require.Len(t, saved, 0)

		err = tx.Rollback(ctx)
		if err != nil {
			t.Fatal("Failed to rollback tx", err)
		}
	})
}
