package apollo

import (
	"context"
	"github.com/gogf/gf/v2/os/gcfg"
)

func LoadAdapter(ctx context.Context, fileName ...string) (gcfg.Adapter, error) {
	apolloConfig, err := LoadConfig(ctx, fileName...)
	if err != nil {
		return nil, err
	}
	return NewAdapterApollo(ctx, apolloConfig)
}
