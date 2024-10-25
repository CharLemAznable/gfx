package gclientx

import (
	"context"
	"github.com/CharLemAznable/gfx/os/gviewx"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gview"
	"net/http"
)

func (c *Client) DoTmplRequest(ctx context.Context,
	view *gviewx.View, key string, params ...gview.Params) (*http.Response, error) {
	return c.DoRawFnRequest(ctx, rawFnWithTmpl(view, key, params...))
}

func (c *Client) TmplRequestBytes(ctx context.Context,
	view *gviewx.View, key string, params ...gview.Params) ([]byte, error) {
	return c.RawFnRequestBytes(ctx, rawFnWithTmpl(view, key, params...))
}

func (c *Client) TmplRequestContent(ctx context.Context, view *gviewx.View,
	key string, params ...gview.Params) (string, error) {
	return c.RawFnRequestContent(ctx, rawFnWithTmpl(view, key, params...))
}

func (c *Client) TmplRequestVar(ctx context.Context, view *gviewx.View,
	key string, params ...gview.Params) (*gvar.Var, error) {
	return c.RawFnRequestVar(ctx, rawFnWithTmpl(view, key, params...))
}

func (c *Client) TmplEventSource(ctx context.Context, view *gviewx.View,
	key string, params ...gview.Params) EventSource {
	return c.RawFnEventSource(ctx, rawFnWithTmpl(view, key, params...))
}

func rawFnWithTmpl(view *gviewx.View, key string, params ...gview.Params) func(context.Context) (string, error) {
	return func(ctx context.Context) (content string, err error) {
		if content, err = view.Parse(ctx, key, params...); err != nil {
			err = gerror.Wrapf(err, `parse tmpl failed`)
		}
		return
	}
}
