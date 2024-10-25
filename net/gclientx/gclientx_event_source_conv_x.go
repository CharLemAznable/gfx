package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) GetEventSource(ctx context.Context, url string, data ...interface{}) EventSource {
	return c.EventSource(ctx, http.MethodGet, url, data...)
}

func (c *Client) PutEventSource(ctx context.Context, url string, data ...interface{}) EventSource {
	return c.EventSource(ctx, http.MethodPut, url, data...)
}

func (c *Client) PostEventSource(ctx context.Context, url string, data ...interface{}) EventSource {
	return c.EventSource(ctx, http.MethodPost, url, data...)
}

func (c *Client) DeleteEventSource(ctx context.Context, url string, data ...interface{}) EventSource {
	return c.EventSource(ctx, http.MethodDelete, url, data...)
}

func (c *Client) HeadEventSource(ctx context.Context, url string, data ...interface{}) EventSource {
	return c.EventSource(ctx, http.MethodHead, url, data...)
}

func (c *Client) PatchEventSource(ctx context.Context, url string, data ...interface{}) EventSource {
	return c.EventSource(ctx, http.MethodPatch, url, data...)
}

func (c *Client) ConnectEventSource(ctx context.Context, url string, data ...interface{}) EventSource {
	return c.EventSource(ctx, http.MethodConnect, url, data...)
}

func (c *Client) OptionsEventSource(ctx context.Context, url string, data ...interface{}) EventSource {
	return c.EventSource(ctx, http.MethodOptions, url, data...)
}

func (c *Client) TraceEventSource(ctx context.Context, url string, data ...interface{}) EventSource {
	return c.EventSource(ctx, http.MethodTrace, url, data...)
}
