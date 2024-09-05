package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) GetBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodGet, url, data...)
}

func (c *Client) PutBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodPut, url, data...)
}

func (c *Client) PostBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodPost, url, data...)
}

func (c *Client) DeleteBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodDelete, url, data...)
}

func (c *Client) HeadBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodHead, url, data...)
}

func (c *Client) PatchBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodPatch, url, data...)
}

func (c *Client) ConnectBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodConnect, url, data...)
}

func (c *Client) OptionsBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodOptions, url, data...)
}

func (c *Client) TraceBytes(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytes(ctx, http.MethodTrace, url, data...)
}
