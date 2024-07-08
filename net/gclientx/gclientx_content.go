package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) GetContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodGet, url, data...)
}

func (c *Client) PostContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodPost, url, data...)
}

func (c *Client) RequestContentErr(ctx context.Context, method string, url string, data ...interface{}) (string, error) {
	bytes, err := c.RequestBytesErr(ctx, method, url, data...)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
