package garritsen

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBool(t *testing.T) {
	tests := []struct {
		name         string
		variable     *Variable
		expectValue  bool
		defaultValue bool
	}{
		{
			name:         "Nil value with default false",
			variable:     &Variable{},
			expectValue:  false,
			defaultValue: false,
		},
		{
			name:         "Nil value with default true",
			variable:     &Variable{},
			expectValue:  true,
			defaultValue: true,
		},
		{
			name: "value not bool with default false",
			variable: &Variable{
				value: "string",
			},
			expectValue:  false,
			defaultValue: false,
		},
		{
			name: "value not bool with default true",
			variable: &Variable{
				value: "string",
			},
			expectValue:  true,
			defaultValue: true,
		},
		{
			name: "value is bool with default false",
			variable: &Variable{
				value: true,
			},
			expectValue:  true,
			defaultValue: false,
		},
		{
			name: "value is bool with default true",
			variable: &Variable{
				value: true,
			},
			expectValue:  true,
			defaultValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectValue, tt.variable.Bool(tt.defaultValue))
		})
	}
}
