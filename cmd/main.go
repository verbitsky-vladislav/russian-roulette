package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"russian-roulette/cmd/setup"
	"russian-roulette/internal/config"
	"time"
)

func newLogger() *zap.Logger {
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02 15:04:05"))
	}

	zcfg := zap.NewDevelopmentConfig()
	zcfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:          "msg",
		LevelKey:            "level",
		TimeKey:             "timestamp",
		NameKey:             "logger",
		CallerKey:           "caller",
		FunctionKey:         "",
		StacktraceKey:       "stacktrace",
		SkipLineEnding:      false,
		LineEnding:          "\n\n",
		EncodeLevel:         zapcore.CapitalColorLevelEncoder,
		EncodeTime:          customTimeEncoder,
		EncodeDuration:      zapcore.SecondsDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          func(string, zapcore.PrimitiveArrayEncoder) {},
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "  ",
	}

	encoder := zapcore.NewConsoleEncoder(zcfg.EncoderConfig)

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())

	return logger
}

func main() {
	logger := newLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	setup.Setup(cfg, logger)
}
