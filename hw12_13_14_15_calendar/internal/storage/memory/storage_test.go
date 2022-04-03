package memorystorage

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	memorystorage "github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
)

func TestStorage(t *testing.T) {
	storage := New()

	t.Run("common test", func(t *testing.T) {
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

		event := memorystorage.NewEvent(
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
		require.Equal(t, event.Title, saved[0].Title)
		require.Equal(t, event.Description, saved[0].Description)

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
	})

	t.Run("test ", func(t *testing.T) {
		events := []memorystorage.Event{
			{
				ID:        parseUUID("4927aa58-a175-429a-a125-c04765597150"),
				StartedAt: parseDateTime("2022-04-03 11:59:59"),
				Notify:    parseDateTime("2022-04-03 11:59:59"),
			},
			{
				ID:        parseUUID("4927aa58-a175-429a-a125-c04765597151"),
				StartedAt: parseDateTime("2022-04-03 12:00:00"),
				Notify:    parseDateTime("2022-04-03 12:00:00"),
			},
			{
				ID:        parseUUID("4927aa58-a175-429a-a125-c04765597152"),
				StartedAt: parseDateTime("2022-04-04 12:00:00"),
				Notify:    parseDateTime("2022-04-03 12:00:00"),
			},
			{
				ID:        parseUUID("4927aa58-a175-429a-a125-c04765597153"),
				StartedAt: parseDateTime("2022-04-05 12:00:01"),
				Notify:    parseDateTime("2022-04-04 11:59:01"),
			},
		}

		for _, e := range events {
			_ = storage.Create(e)
		}

		readyEvents, err := storage.GetEventsReadyToNotify(parseDateTime("2022-04-03 12:00:00"))
		require.Nil(t, err)

		ids := extractEventIDs(readyEvents)
		idsExpected := []string{
			"4927aa58-a175-429a-a125-c04765597150",
			"4927aa58-a175-429a-a125-c04765597151",
			"4927aa58-a175-429a-a125-c04765597152",
		}
		require.Equal(t, idsExpected, ids)
	})
}

func parseUUID(stringUUID string) uuid.UUID {
	res, _ := uuid.Parse(stringUUID)

	return res
}

func parseDateTime(stringDateTime string) time.Time {
	res, _ := time.Parse("2006-01-02 15:04:05", stringDateTime)

	return res
}

func extractEventIDs(events []memorystorage.Event) []string {
	res := make([]string, 0, len(events))
	for _, e := range events {
		res = append(res, e.ID.String())
	}

	return res
}
