package gclientx

import (
	"context"
	"github.com/gogf/gf/v2/net/gclient"
	"io"
	"net/http"
)

func (c *Client) deferCloseResponse(ctx context.Context, response *gclient.Response) {
	c.deferLogError(ctx, response.Close())
}

func (c *Client) deferCloseRawResponse(ctx context.Context, response *http.Response) {
	c.deferLogError(ctx, func() error {
		if response == nil {
			return nil
		}
		return response.Body.Close()
	}())
}

func (c *Client) deferLogError(ctx context.Context, err error) {
	if err != nil && c.intlog != nil {
		c.intlog.Errorf(ctx, `%+v`, err)
	}
}

func (c *Client) readAll(ctx context.Context, response *http.Response) []byte {
	if response == nil {
		return []byte{}
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.deferLogError(ctx, err)
		return nil
	}
	return body
}
