package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) GetBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodGet, url, data...)
}

func (c *Client) PostBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodPost, url, data...)
}

func (c *Client) RequestBytesErr(ctx context.Context, method string, url string, data ...interface{}) ([]byte, error) {
	response, err := c.Client.DoRequest(ctx, method, url, data...)
	if err != nil {
		if c.errorFn != nil {
			c.errorFn(ctx, `%+v`, err)
		}
		return nil, err
	}
	defer func() {
		if err = response.Close(); err != nil {
			if c.errorFn != nil {
				c.errorFn(ctx, `%+v`, err)
			}
		}
	}()
	return response.ReadAll(), nil
}
