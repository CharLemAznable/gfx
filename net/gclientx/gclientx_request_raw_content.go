package gclientx

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"net/http"
)

func (c *Client) DoRawContentRequest(ctx context.Context, content string) (response *http.Response, err error) {
	return c.DoRawFnRequest(ctx, rawFnWithContent(content))
}

func (c *Client) RawContentRequestBytes(ctx context.Context, content string) ([]byte, error) {
	return c.RawFnRequestBytes(ctx, rawFnWithContent(content))
}

func (c *Client) RawContentRequestContent(ctx context.Context, content string) (string, error) {
	return c.RawFnRequestContent(ctx, rawFnWithContent(content))
}

func (c *Client) RawContentRequestVar(ctx context.Context, content string) (*gvar.Var, error) {
	return c.RawFnRequestVar(ctx, rawFnWithContent(content))
}

func (c *Client) RawContentEventSource(content string) EventSource {
	return c.RawFnEventSource(rawFnWithContent(content))
}

func rawFnWithContent(content string) func(context.Context) (string, error) {
	return func(_ context.Context) (string, error) { return content, nil }
}
