package sqlstorage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	ctx  context.Context
	conn *pgx.Conn
	dsn  string
}

func New(ctx context.Context, dsn string) *Storage {
	return &Storage{
		ctx: ctx,
		dsn: dsn,
	}
}

func (s *Storage) Connect(ctx context.Context) app.Storage {
	conn, err := pgx.Connect(ctx, s.dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect database: %v\n", err)
		os.Exit(1)
	}
	s.conn = conn
	return s
}

func (s *Storage) Close(ctx context.Context) error {
	return s.conn.Close(ctx)
}

func (s *Storage) Create(e storage.Event) error {
	sql := `
insert into events (
	id,
    title,
    started_at,
    finished_at,
    description,
    user_id,
    notify
) values ($1, $2, $3, $4, $5, $6, $7)
`
	_, err := s.conn.Exec(
		s.ctx,
		sql,
		e.ID.String(),
		e.Title,
		e.StartedAt.Format(time.RFC3339),
		e.FinishedAt.Format(time.RFC3339),
		e.Description,
		e.UserID,
		e.Notify.Format(time.RFC3339),
	)

	return err
}

func (s *Storage) Update(e storage.Event) error {
	sql := `
update events 
set
    title = $2,
    started_at = $3,
    finished_at = $4,
    description = $5,
    user_id = $6,
    notify = $7
where
	id = $1
`
	_, err := s.conn.Exec(
		s.ctx,
		sql,
		e.ID.String(),
		e.Title,
		e.StartedAt.Format(time.RFC3339),
		e.FinishedAt.Format(time.RFC3339),
		e.Description,
		e.UserID,
		e.Notify.Format(time.RFC3339),
	)

	return err
}

func (s *Storage) Delete(id uuid.UUID) error {
	sql := "DELETE FROM events WHERE id = $1"
	_, err := s.conn.Exec(s.ctx, sql, id)

	return err
}

func (s *Storage) Find(id uuid.UUID) (*storage.Event, error) {
	var e storage.Event

	sql := `
select id, title, started_at, finished_at, description, user_id, notify 
from events
where id = $1
`
	err := s.conn.QueryRow(s.ctx, sql, id).Scan(
		&e.ID,
		&e.Title,
		&e.StartedAt,
		&e.FinishedAt,
		&e.Description,
		&e.UserID,
		&e.Notify,
	)

	if err == nil {
		return &e, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return nil, fmt.Errorf("failed to scan SQL result into struct: %w", err)
}

func (s *Storage) FindAll() ([]storage.Event, error) {
	events := make([]storage.Event, 0)

	sql := `
select id, title, started_at, finished_at, description, user_id, notify 
from events
order by date
`
	rows, err := s.conn.Query(s.ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e storage.Event
		if err := rows.Scan(
			&e.ID,
			&e.Title,
			&e.StartedAt,
			&e.FinishedAt,
			&e.Description,
			&e.UserID,
			&e.Notify,
		); err != nil {
			return nil, fmt.Errorf("unable to transform array result into struct: %w", err)
		}

		events = append(events, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
