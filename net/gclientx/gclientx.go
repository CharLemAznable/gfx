package gclientx

import (
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/glog"
)

type Client struct {
	*gclient.Client
	intlog *glog.Logger
}

func New(client ...*gclient.Client) *Client {
	if len(client) > 0 && client[0] != nil {
		return &Client{Client: client[0]}
	}
	return &Client{Client: gclient.New()}
}

func (c *Client) SetIntLog(intlog *glog.Logger) *Client {
	c.intlog = intlog
	return c
}

func (c *Client) Clone() *Client {
	return New(c.Client.Clone()).SetIntLog(c.intlog)
}
