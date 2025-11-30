package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger() error {
	// Create logs directory if it doesn't exist
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	// Define log file paths
	logFile := filepath.Join(logDir, "app.log")
	errorLogFile := filepath.Join(logDir, "error.log")

	// Create log files if they don't exist
	_, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = os.OpenFile(errorLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Configure encoder
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	// Create file writers
	logFileWriter, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	errorLogFileWriter, err := os.OpenFile(errorLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Create cores
	core := zapcore.NewTee(
		// Write all logs to app.log
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			zapcore.AddSync(logFileWriter),
			zapcore.InfoLevel,
		),
		// Write error and fatal logs to error.log
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			zapcore.AddSync(errorLogFileWriter),
			zapcore.ErrorLevel,
		),
		// Also output to console
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		),
	)

	// Create logger
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return nil
}

func GetLogger() *zap.Logger {
	if Log == nil {
		// Fallback to a simple logger if InitLogger hasn't been called
		logger, _ := zap.NewProduction()
		return logger
	}
	return Log
}
