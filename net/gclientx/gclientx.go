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
	newClient := gclient.New()
	*newClient = *c.Client
	// alloc header map anyway
	newHeader := (*map[string]string)(offsetPointer(newClient, headerOffset))
	*newHeader = c.GetHeaderMap()
	// alloc cookies map anyway
	newCookies := (*map[string]string)(offsetPointer(newClient, cookiesOffset))
	*newCookies = c.GetCookieMap()
	return New(newClient).SetIntLog(c.intlog)
}
