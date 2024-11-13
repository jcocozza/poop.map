package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

// create the log file
func logFile() (*os.File, error) {
	path := fmt.Sprintf("cassidy-log-%d.log", time.Now().Unix())
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

func CreateLogger(level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}
	f, err := logFile()
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewTextHandler(f, opts))
	return logger
}
