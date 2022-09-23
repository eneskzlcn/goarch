package logger_test

import (
	logger "<module_name>/../<directory_name>/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewZapLoggerAdapter(t *testing.T) {
	bestCases := []string{"", "qa", "local", "test", "dev", "prod"}
	worstCases := []string{",", "afasfaf", "any", "123", "afsaf"}

	for _, bestCase := range bestCases {
		zapLogger := logger.NewZapLoggerAdapter(bestCase, 1)
		assert.NotNil(t, zapLogger)
	}
	for _, worstCase := range worstCases {
		zapLogger := logger.NewZapLoggerAdapter(worstCase, 1)
		assert.Nil(t, zapLogger)
	}
}
