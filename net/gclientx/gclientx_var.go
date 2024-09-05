package gclientx

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
)

func (c *Client) RequestVar(ctx context.Context, method string, url string, data ...interface{}) (*gvar.Var, error) {
	bytes, err := c.RequestBytes(ctx, method, url, data...)
	if err != nil {
		return gvar.New(nil), err
	}
	return gvar.New(bytes), nil
}
