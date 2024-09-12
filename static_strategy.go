package garritsen

import (
	"context"
	"gopkg.in/yaml.v3"
	"os"
)

type StaticVariable struct {
	Type    DataType      `yaml:"type"`
	Default interface{}   `yaml:"default"`
	Rollout []interface{} `yaml:"rollout"`
}

type staticStrategy struct {
	Variables map[string]StaticVariable
	logger    ILogger
}

type StaticStrategyOptions struct {
	Path   string
	Logger ILogger
}

func NewStaticStrategy(options StaticStrategyOptions) (*staticStrategy, error) {

	data, err := os.ReadFile(options.Path)
	if err != nil {
		return nil, err
	}
	var arr map[string]StaticVariable
	err = yaml.Unmarshal(data, &arr)
	if err != nil {
		return nil, err
	}
	return &staticStrategy{
		Variables: arr,
		logger:    options.Logger,
	}, nil
}

func (s *staticStrategy) GetVariable(ctx context.Context, variable string) *Variable {
	val, ok := s.Variables[variable]
	if !ok {
		s.logger.Info("Variable not found", "variable", variable)
		return &Variable{}
	}
	if len(val.Rollout) == 0 {
		return &Variable{
			value: val.Default,
		}
	}

	return &Variable{}
}
