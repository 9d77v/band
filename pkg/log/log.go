package log

import (
	"fmt"
	"os"
	"time"

	"log/slog"

	"github.com/lmittmann/tint"
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
	// 文件日志使用 JSON 格式
	fileHandler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: level,
	})

	if conf.Level == LevelDebug {
		// 控制台使用带颜色的 Text 格式
		consoleHandler := tint.NewHandler(os.Stdout, &tint.Options{
			Level:      level,
			TimeFormat: time.TimeOnly,
			AddSource:  false,
		})
		Logger = slog.New(slog.NewMultiHandler(consoleHandler, fileHandler))
	} else {
		Logger = slog.New(fileHandler)
	}
	slog.SetDefault(Logger)
	slog.Info("log init success")
}
