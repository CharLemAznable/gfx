package apollo

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

const (
	defaultApolloConfigFileName = "apollo"
)

func LoadAdapter(ctx context.Context, fileName ...string) (gcfg.Adapter, error) {
	apolloConfigFileName := defaultApolloConfigFileName
	if len(fileName) > 0 && fileName[0] != "" {
		apolloConfigFileName = fileName[0]
	}
	apolloCfg := g.Cfg(apolloConfigFileName)
	apolloCfgData, err := apolloCfg.Data(ctx)
	if err != nil {
		return nil, err
	}
	apolloCfgVar := g.NewVar(apolloCfgData, true)
	apolloConfig := Config{}
	err = apolloCfgVar.Struct(&apolloConfig)
	if err != nil {
		return nil, err
	}
	return NewAdapterApollo(ctx, apolloConfig)
}
