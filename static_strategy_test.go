package garritsen

import (
	"context"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"
)

func TestNewStaticStrategy(t *testing.T) {
	tests := []struct {
		options   StaticStrategyOptions
		expectErr error
		name      string
	}{
		{
			name: "Init successfully",
			options: StaticStrategyOptions{
				Path: "variable.yaml",
			},
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewStaticStrategy(tt.options)
			require.Equal(t, tt.expectErr, err)
		})
	}
}

func TestGetVariable(t *testing.T) {
	tests := []struct {
		name           string
		options        StaticStrategyOptions
		variable       string
		expectVariable *Variable
	}{
		{
			name: "Variable not found",
			options: StaticStrategyOptions{
				Path:   "variable.yaml",
				Logger: slog.Default(),
			},
			variable:       "notfound",
			expectVariable: &Variable{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strategy, err := NewStaticStrategy(tt.options)
			require.NoError(t, err)
			variable := strategy.GetVariable(context.TODO(), tt.variable)
			require.Equal(t, tt.expectVariable, variable)
		})
	}
}
