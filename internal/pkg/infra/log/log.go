package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New serve caller to create zap.Logger
func New(logLevel string, output string) (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
	)

	err = level.UnmarshalText([]byte(logLevel))
	if err != nil {
		return nil, err
	}

	cw := zapcore.Lock(os.Stdout)
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(config)
	if output == "json" {
		config = zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.RFC3339NanoTimeEncoder
		enc = zapcore.NewJSONEncoder(config)
	}

	cores := make([]zapcore.Core, 0, 2)
	cores = append(cores, zapcore.NewCore(enc, cw, level))

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)

	return logger, nil
}
