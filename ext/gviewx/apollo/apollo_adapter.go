package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gvalid"
)

type AdapterApollo struct {
	client *agollox.Client
	config *Config
}

var (
	configRules = map[string]string{
		"appId": "required",
		"ip":    "required",
	}
	configMessage = map[string]interface{}{
		"appId": "gviewx: Apollo AppId field is required",
		"ip":    "gviewx: Apollo IP field is required",
	}
)

func NewAdapterApollo(ctx context.Context, config *Config) (adapter *AdapterApollo, err error) {
	// Data validation.
	err = gvalid.New().Rules(configRules).Messages(configMessage).Data(config).Run(ctx)
	if err != nil {
		return
	}
	agolloClient, err := agollox.NewClient(config)
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
