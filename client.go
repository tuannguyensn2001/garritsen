package garritsen

import "context"

type client struct {
	strategy IStrategy
}

type IStrategy interface {
	GetVariable(ctx context.Context, variable string) *Variable
}

func NewClient(strategy IStrategy) *client {
	return &client{
		strategy: strategy,
	}
}

func (c *client) GetVariable(ctx context.Context, variable string) *Variable {
	return c.strategy.GetVariable(ctx, variable)
}
