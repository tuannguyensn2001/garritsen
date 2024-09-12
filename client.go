package garritsen

import (
	"context"
	"log/slog"
)

type client struct {
	strategy IStrategy
	logger   ILogger
}

type IStrategy interface {
	GetVariable(ctx context.Context, variable string) *Variable
}

type ClientOptions struct {
	Logger                ILogger
	Strategy              Strategy
	StaticStrategyOptions *StaticStrategyOptions
}

func NewClient(opts *ClientOptions) (*client, error) {
	var strategy IStrategy
	if opts.Logger == nil {
		opts.Logger = slog.Default()
	}

	switch opts.Strategy {
	case Static:
		staticStrategy, err := NewStaticStrategy(*opts.StaticStrategyOptions)
		if err != nil {
			return nil, err
		}
		if staticStrategy.logger == nil {
			staticStrategy.logger = opts.Logger
		}
		strategy = staticStrategy
	}
	return &client{
		strategy: strategy,
		logger:   opts.Logger,
	}, nil
}

func (c *client) GetVariable(ctx context.Context, variable string) *Variable {
	return c.strategy.GetVariable(ctx, variable)
}
