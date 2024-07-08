package gclientx

import (
	"context"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/glog"
)

type Client struct {
	gclient.Client
	errorFn func(ctx context.Context, format string, v ...interface{})
}

func New(client *gclient.Client) *Client {
	return &Client{Client: *client}
}

func (c *Client) SetErrorFn(errorFn func(ctx context.Context, format string, v ...interface{})) *Client {
	c.errorFn = errorFn
	return c
}

func (c *Client) SetErrorLogger(logger *glog.Logger) *Client {
	if logger != nil {
		return c.SetErrorFn(logger.Errorf)
	} else {
		return c.SetErrorFn(nil)
	}
}
