package gclientx

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"net/http"
)

func (c *Client) GetVar(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVar(ctx, http.MethodGet, url, data...)
}

func (c *Client) PutVar(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVar(ctx, http.MethodPut, url, data...)
}

func (c *Client) PostVar(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVar(ctx, http.MethodPost, url, data...)
}

func (c *Client) DeleteVar(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVar(ctx, http.MethodDelete, url, data...)
}

func (c *Client) HeadVar(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVar(ctx, http.MethodHead, url, data...)
}

func (c *Client) PatchVar(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVar(ctx, http.MethodPatch, url, data...)
}

func (c *Client) ConnectVar(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVar(ctx, http.MethodConnect, url, data...)
}

func (c *Client) OptionsVar(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVar(ctx, http.MethodOptions, url, data...)
}

func (c *Client) TraceVar(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVar(ctx, http.MethodTrace, url, data...)
}
