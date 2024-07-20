package gclientx

import (
	"context"
	"github.com/gogf/gf/v2/net/gclient"
)

func (c *Client) deferCloseResponse(ctx context.Context, response *gclient.Response) {
	c.deferLogError(ctx, response.Close())
}

func (c *Client) deferLogError(ctx context.Context, err error) {
	if err != nil && c.intlog != nil {
		c.intlog.Errorf(ctx, `%+v`, err)
	}
}
