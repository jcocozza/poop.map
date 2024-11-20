package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jcocozza/poop.map/backend/internal/config"
)

// create the log file
func logFile(env config.Environment) (*os.File, error) {
	path := fmt.Sprintf("cassidy-log-%d-%s.log", time.Now().Unix(), env)
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

func CreateLogger(level slog.Level, env config.Environment) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}
	f, err := logFile(env)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewTextHandler(f, opts))
	return logger
}
