package logger

import (
	"context"
	"log/slog"
	"time"
)

// Level represents different logging levels.
type Level slog.Level

// A set of possible logging levels.
const (
	LevelInfo  = Level(slog.LevelInfo)
	LevelWarn  = Level(slog.LevelWarn)
	LevelError = Level(slog.LevelError)
	LevelDebug = Level(slog.LevelDebug)
)

// Record represents the data that is being logged.
type Record struct {
	Time    time.Time
	Level   Level
	Message string
	//Attribute[T any] map[string]T

	Attribute map[string]any
}

func toRecord(r slog.Record) Record {
	atts := make(map[string]any, r.NumAttrs())

	f := func(attr slog.Attr) bool {
		atts[attr.Key] = attr.Value.Any()
		return true
	}

	r.Attrs(f)

	return Record{
		Time:      r.Time,
		Message:   r.Message,
		Level:     Level(r.Level),
		Attribute: atts,
	}
}

// EventFn is a function to be executed when configured against a log level.
type EventFn func(ctx context.Context, rec Record)

// Events contains an assignment of an event function to a log level.
type Events struct {
	Info  EventFn
	Debug EventFn
	Warn  EventFn
	Error EventFn
}
