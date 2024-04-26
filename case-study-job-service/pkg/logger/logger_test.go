package logger

import (
	"case-study-job-service/pkg/config"
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestGetLogger(t *testing.T) {
	actual := GetLogger()
	require.Equal(t, logger, actual)
}

func TestRegisterLogger(t *testing.T) {
	tests := []struct {
		name         string
		cfg          config.LoggerConfig
		expectations func()
	}{
		{
			name: "environment type is production",
			cfg: config.LoggerConfig{
				AppName:         "case-study-job-service",
				LogLevel:        0,
				LogEncoding:     "json",
				EnvironmentType: "production",
			},
			expectations: nil,
		},
		{
			name: "environment type is test",
			cfg: config.LoggerConfig{
				AppName:         "case-study-job-service",
				LogLevel:        0,
				LogEncoding:     "console",
				EnvironmentType: "test",
			},
			expectations: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := RegisterLogger(tt.cfg)
			require.Equal(t, GetLogger(), actual)
		})
	}
}

func TestLoggerContext(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		l        *zap.Logger
		expected *zap.Logger
	}{
		{
			name:     "extract to logger from context",
			ctx:      context.Background(),
			l:        GetLogger(),
			expected: GetLogger(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newCtx := WithCtx(tt.ctx, tt.l)
			actual := FromCtx(newCtx)

			require.Equal(t, tt.expected, actual)
		})
	}
}
