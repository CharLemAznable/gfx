package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) GetBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodGet, url, data...)
}

func (c *Client) PutBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodPut, url, data...)
}

func (c *Client) PostBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodPost, url, data...)
}

func (c *Client) DeleteBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodDelete, url, data...)
}

func (c *Client) HeadBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodHead, url, data...)
}

func (c *Client) PatchBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodPatch, url, data...)
}

func (c *Client) ConnectBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodConnect, url, data...)
}

func (c *Client) OptionsBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodOptions, url, data...)
}

func (c *Client) TraceBytesErr(ctx context.Context, url string, data ...interface{}) ([]byte, error) {
	return c.RequestBytesErr(ctx, http.MethodTrace, url, data...)
}
