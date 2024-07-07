package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) GetContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodGet, url, data...)
}

func (c *Client) PostContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodPost, url, data...)
}

func (c *Client) RequestContent(ctx context.Context, method string, url string, data ...interface{}) (string, error) {
	bytes, err := c.RequestBytes(ctx, method, url, data...)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
