package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	w := zapcore.AddSync(os.Stdout)
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	config := zap.NewDevelopmentEncoderConfig()
	config.MessageKey = "message"
	config.LevelKey = "severity"
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		atomicLevel,
	)

	var fields []zap.Field
	fields = append(fields, zap.String("service", "MyService"))
	logger := zap.New(core, zap.Fields(fields...))

	log := MyLogger{logger: logger}

	log.Message(Debug, "message1", "field1", "it's sunday today 233")
	log.Message(Debug, "message1", "field1", errors.New("test error"))
}

type MyLogger struct {
	logger *zap.Logger
}

const (
	Fatal int = iota
	Error
	Warn
	Info
	Debug
)

func (c *MyLogger) Message(level int, message string, f ...interface{}) {
	logFields := make([]zap.Field, 0, len(f)/2)

	key := ""
	value := []byte{}
	for k, v := range f {
		value = value[:0]
		if k%2 == 0 {
			switch v.(type) {
			case string:
				key = v.(string)
			}
			continue
		} else {
			value, _ = json.Marshal(v)
			// value = v
		}
		fmt.Println(value)

		logFields = append(logFields, zap.String(key, string(value)))
		// logFields = append(logFields, zap.Any(key, v))
	}

	switch level {
	case Debug:
		c.logger.Debug(message, logFields...)
	}
}
