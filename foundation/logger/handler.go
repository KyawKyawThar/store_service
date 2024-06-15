package logger

import (
	"context"
	"log/slog"
)

// logHandler provides a wrapper around the slog handler to capture which
// log level is being logged for event handling.
type logHandler struct {
	handler slog.Handler
	event   Events
}

func newLogHandler(h slog.Handler, e Events) *logHandler {
	return &logHandler{
		handler: h,
		event:   e,
	}
}

// WithAttrs returns a new JSONHandler whose attributes consists
// of h's attributes followed by attrs.
func (h *logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &logHandler{
		handler: h.handler.WithAttrs(attrs),
		event:   h.event,
	}
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (h *logHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

// Handle looks to see if an event function needs to be executed for a given
// log level and then formats its argument Record.
func (h *logHandler) Handle(ctx context.Context, record slog.Record) error {

	switch record.Level {
	case slog.LevelInfo:
		if h.event.Info != nil {
			h.event.Info(ctx, toRecord(record))
		}
	case slog.LevelWarn:
		if h.event.Warn != nil {
			h.event.Warn(ctx, toRecord(record))
		}
	case slog.LevelError:
		if h.event.Error != nil {
			h.event.Error(ctx, toRecord(record))
		}

	case slog.LevelDebug:
		if h.event.Debug != nil {
			h.event.Debug(ctx, toRecord(record))
		}
	}
	return h.handler.Handle(ctx, record)
}

// WithGroup returns a new Handler with the given group appended to the receiver's
// existing groups. The keys of all subsequent attributes, whether added by With
// or in a Record, should be qualified by the sequence of group names.
func (h *logHandler) WithGroup(name string) slog.Handler {
	return &logHandler{
		handler: h.handler.WithGroup(name),
		event:   h.event,
	}

}
