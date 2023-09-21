package log

import (
	"fmt"
	"io"
	"os"

	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	conf := FromEnv()
	var level slog.Level
	err := level.UnmarshalText([]byte(conf.Level))
	if err != nil {
		panic(err)
	}
	path := fmt.Sprintf("logs/%s.log", conf.AppName)
	writer := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge, //days
		LocalTime:  true,
		Compress:   false, // disabled by default
	}
	writers := []io.Writer{writer, os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	slog.SetDefault(slog.New(slog.NewJSONHandler(fileAndStdoutWriter, &slog.HandlerOptions{
		Level: level,
	})))
}
