package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type AdapterApollo struct {
	client *agollox.Client
	config *Config
}

func NewAdapterApollo(ctx context.Context, config *Config) (adapter *AdapterApollo, err error) {
	agolloClient, err := agollox.NewClient(ctx, config)
	if err != nil {
		return
	}
	adapter = &AdapterApollo{client: agolloClient, config: config}
	return
}

func (a *AdapterApollo) GetContent(key string) (string, error) {
	if !a.client.Contains(key) {
		return "", gerror.NewCodef(gcode.CodeInvalidParameter, `get content failed with apollo key: %s`, key)
	}
	return gvar.New(a.client.Get(key)).String(), nil
}
