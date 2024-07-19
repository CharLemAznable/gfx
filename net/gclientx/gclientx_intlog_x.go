package gclientx

import (
	"context"
	"github.com/gogf/gf/v2/net/gclient"
)

func (c *Client) deferCloseResponse(ctx context.Context, response *gclient.Response) {
	if err := response.Close(); err != nil && c.intlog != nil {
		c.intlog.Errorf(ctx, `%+v`, err)
	}
}
