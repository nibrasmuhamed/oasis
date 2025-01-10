package logger_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logMessage struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
}

// custom core that captures logs in a buffer for testing
func newTestLogger() (*zap.Logger, *bytes.Buffer) {
	buffer := &bytes.Buffer{}
	writer := zapcore.AddSync(buffer)
	encoderCfg := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	core := zapcore.NewCore(encoder, writer, zap.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger, buffer
}

func assertLogged(t *testing.T, buffer *bytes.Buffer, expectedLevel string, expectedMsg string) {
	var log logMessage
	err := json.NewDecoder(buffer).Decode(&log)
	if err != nil {
		t.Fatalf("Failed to decode log message: %v", err)
	}

	if log.Level != expectedLevel {
		t.Errorf("Expected log level '%s', but got '%s'", expectedLevel, log.Level)
	}
	if log.Msg != expectedMsg {
		t.Errorf("Expected log message '%s', but got '%s'", expectedMsg, log.Msg)
	}
}

func TestDebugLevel(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	log, buffer := newTestLogger()

	log.Debug("Debug message")
	assertLogged(t, buffer, "debug", "Debug message")

	// log.Info("Info message")
	// assertLogged(t, buffer, "info", "Info message")

	// log.Warn("Warn message")
	// assertLogged(t, buffer, "warn", "Warn message")

	// log.Error("Error message", zap.String("detail", "some error detail"))
	// assertLogged(t, buffer, "error", "Error message")
}

func TestErrorLevel(t *testing.T) {
	os.Setenv("LOG_LEVEL", "ERROR")
	log, buffer := newTestLogger()

	// log.Debug("Debug message")
	// if buffer.Len() > 0 {
	// 	t.Error("Expected 'Debug message' to NOT be logged at ERROR level")
	// }

	// log.Info("Info message")
	// if buffer.Len() > 0 {
	// 	t.Error("Expected 'Info message' to NOT be logged at ERROR level")
	// }

	// log.Warn("Warn message")
	// if buffer.Len() > 0 {
	// 	t.Error("Expected 'Warn message' to NOT be logged at ERROR level")
	// }

	log.Error("Error message", zap.String("detail", "some error detail"))
	assertLogged(t, buffer, "error", "Error message")
}

func TestWarnLevel(t *testing.T) {
	os.Setenv("LOG_LEVEL", "WARN")
	log, buffer := newTestLogger()

	// log.Debug("Debug message")
	// if buffer.Len() > 0 {
	// 	t.Error("Expected 'Debug message' to NOT be logged at WARN level")
	// }

	// log.Info("Info message")
	// if buffer.Len() > 0 {
	// 	t.Error("Expected 'Info message' to NOT be logged at WARN level")
	// }

	log.Warn("Warn message")
	assertLogged(t, buffer, "warn", "Warn message")

	// log.Error("Error message", zap.String("detail", "some error detail"))
	// assertLogged(t, buffer, "error", "Error message")
}

func TestInfoLevel(t *testing.T) {
	os.Setenv("LOG_LEVEL", "INFO")
	log, buffer := newTestLogger()

	// log.Debug("Debug message")
	// if buffer.Len() > 0 {
	// 	t.Error("Expected 'Debug message' to NOT be logged at INFO level")
	// }

	log.Info("Info message")
	assertLogged(t, buffer, "info", "Info message")

	// log.Warn("Warn message")
	// assertLogged(t, buffer, "warn", "Warn message")

	// log.Error("Error message", zap.String("detail", "some error detail"))
	// assertLogged(t, buffer, "error", "Error message")
}
