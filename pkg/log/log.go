package log

import (
	"fmt"
	"io"
	"os"

	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *slog.Logger

func Init() {
	conf := FromEnv()
	var level slog.Level
	err := level.UnmarshalText([]byte(conf.Level))
	if err != nil {
		panic(err)
	}
	path := conf.Path
	if path == "" {
		path = fmt.Sprintf("logs/%s.log", conf.AppName)
	}
	writer := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge, //days
		LocalTime:  true,
		Compress:   false, // disabled by default
	}
	writers := []io.Writer{writer}
	if conf.Level == LevelDebug {
		writers = append(writers, os.Stdout)
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	Logger = slog.New(slog.NewJSONHandler(fileAndStdoutWriter, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(Logger)
}
