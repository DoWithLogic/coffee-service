package observability

import (
	"context"
	"io"
	"log"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type (
	Logger struct {
		standard *log.Logger
		zerolog  *zerolog.Logger
	}
)

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	// ctx := e.GetCtx()
	// span := trace.SpanContextFromContext(ctx)

	// e.Str("span_id", span.SpanID().String()).Str("trace_id", span.TraceID().String())
}

func NewZeroLogHook() *Logger {
	z := zerolog.New(os.Stdout).Hook(TracingHook{}).With().Timestamp().Stack().Logger()

	return &Logger{log.New(z, "", 0), &z}
}

func NewZeroLog(ctx context.Context, c ...io.Writer) *Logger {
	z := zerolog.New(os.Stdout).With().Timestamp().
		Stack().Logger()

	return &Logger{log.New(z, "", 0), &z}
}

func (x *Logger) S() *log.Logger     { return x.standard }
func (x *Logger) Z() *zerolog.Logger { return x.zerolog }
func (x *Logger) Level(level string) *Logger {
	lv, err := zerolog.ParseLevel(strings.ToLower(level))
	if err == nil {
		*x.zerolog = x.zerolog.Level(lv)
		x.standard.SetOutput(x.zerolog)
	}

	return x
}
