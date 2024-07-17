package env

import (
	"github.com/pleimer/ticketer/server/lib/once"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerConfig struct {
	logger *zap.Logger
	Logger func() *zap.Logger
}

func (l *loggerConfig) init() {
	l.Logger = func() *zap.Logger {
		once.Once(func() {
			config := zap.Config{
				Encoding:         "json",
				Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
				OutputPaths:      []string{"stdout"},
				ErrorOutputPaths: []string{"stderr"},
				EncoderConfig: zapcore.EncoderConfig{
					MessageKey:   "message",
					LevelKey:     "level",
					TimeKey:      "time",
					CallerKey:    "caller",
					EncodeCaller: zapcore.ShortCallerEncoder,
					EncodeLevel:  zapcore.CapitalLevelEncoder,
					EncodeTime:   zapcore.ISO8601TimeEncoder,
				},
			}

			var err error

			l.logger, err = config.Build()
			if err != nil {
				panic(err)
			}
		})

		return l.logger
	}
}
