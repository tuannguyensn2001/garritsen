package garritsen

import "context"

type staticStrategy struct {
}

func NewStaticStrategy() *staticStrategy {
	return &staticStrategy{}
}

func (s *staticStrategy) GetVariable(ctx context.Context, variable string) *Variable {
	return &Variable{}
}
