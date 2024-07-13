package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/os/gviewx"
)

func LoadAdapter(ctx context.Context, fileName ...string) (gviewx.Adapter, error) {
	apolloConfig, err := LoadConfig(ctx, fileName...)
	if err != nil {
		return nil, err
	}
	return NewAdapterApollo(ctx, apolloConfig)
}
