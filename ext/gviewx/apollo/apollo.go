package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/CharLemAznable/gfx/os/gviewx"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func LoadAdapter(ctx context.Context) (gviewx.Adapter, error) {
	return NewAdapterApollo(ctx)
}

func NewAdapterApollo(ctx context.Context) (adapter *AdapterApollo, err error) {
	client, err := agollox.NewClient(ctx)
	if err != nil {
		return
	}
	adapter = &AdapterApollo{client: client}
	return
}

type AdapterApollo struct {
	client *agollox.Client
}

func (a *AdapterApollo) GetContent(key string) (string, error) {
	if !a.client.Contains(key) {
		return "", gerror.NewCodef(gcode.CodeInvalidParameter, `get content failed with apollo key: %s`, key)
	}
	return gvar.New(a.client.Get(key)).String(), nil
}
