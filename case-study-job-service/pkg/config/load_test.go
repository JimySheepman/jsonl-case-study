package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		isError bool
	}{
		{
			name:    "viper read in config path failure",
			path:    ".",
			isError: true,
		},
		{
			name:    "load config succeed",
			path:    "../../.",
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := LoadConfig(tt.path)
			if tt.isError {
				require.Error(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
