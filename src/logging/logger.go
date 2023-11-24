package logging

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

var DefLogger *slog.Logger

func InitLogger() {

	w := os.Stderr

	// create a new logger
	DefLogger = slog.New(tint.NewHandler(w, nil))

	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
}
