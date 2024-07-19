package gclientx

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"net/http"
)

func (c *Client) GetVarErr(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVarErr(ctx, http.MethodGet, url, data...)
}

func (c *Client) PutVarErr(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVarErr(ctx, http.MethodPut, url, data...)
}

func (c *Client) PostVarErr(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVarErr(ctx, http.MethodPost, url, data...)
}

func (c *Client) DeleteVarErr(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVarErr(ctx, http.MethodDelete, url, data...)
}

func (c *Client) HeadVarErr(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVarErr(ctx, http.MethodHead, url, data...)
}

func (c *Client) PatchVarErr(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVarErr(ctx, http.MethodPatch, url, data...)
}

func (c *Client) ConnectVarErr(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVarErr(ctx, http.MethodConnect, url, data...)
}

func (c *Client) OptionsVarErr(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVarErr(ctx, http.MethodOptions, url, data...)
}

func (c *Client) TraceVarErr(ctx context.Context, url string, data ...interface{}) (*gvar.Var, error) {
	return c.RequestVarErr(ctx, http.MethodTrace, url, data...)
}
