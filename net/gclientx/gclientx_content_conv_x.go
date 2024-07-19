package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) GetContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodGet, url, data...)
}

func (c *Client) PutContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodPut, url, data...)
}

func (c *Client) PostContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodPost, url, data...)
}

func (c *Client) DeleteContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodDelete, url, data...)
}

func (c *Client) HeadContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodHead, url, data...)
}

func (c *Client) PatchContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodPatch, url, data...)
}

func (c *Client) ConnectContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodConnect, url, data...)
}

func (c *Client) OptionsContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodOptions, url, data...)
}

func (c *Client) TraceContentErr(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContentErr(ctx, http.MethodTrace, url, data...)
}
