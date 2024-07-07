package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) GetBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodGet, url, data...)
}

func (c *Client) PostBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodPost, url, data...)
}

func (c *Client) RequestBytes(ctx context.Context, method string, url string, data ...interface{}) ([]byte, error) {
	response, err := c.client.DoRequest(ctx, method, url, data...)
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
