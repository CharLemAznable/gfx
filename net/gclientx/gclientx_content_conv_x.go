package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) GetContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodGet, url, data...)
}

func (c *Client) PutContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodPut, url, data...)
}

func (c *Client) PostContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodPost, url, data...)
}

func (c *Client) DeleteContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodDelete, url, data...)
}

func (c *Client) HeadContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodHead, url, data...)
}

func (c *Client) PatchContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodPatch, url, data...)
}

func (c *Client) ConnectContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodConnect, url, data...)
}

func (c *Client) OptionsContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodOptions, url, data...)
}

func (c *Client) TraceContent(ctx context.Context, url string, data ...interface{}) (string, error) {
	return c.RequestContent(ctx, http.MethodTrace, url, data...)
}
