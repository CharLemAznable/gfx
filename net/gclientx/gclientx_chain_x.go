package gclientx

import (
	"github.com/gogf/gf/v2/net/gsvc"
	"time"
)

func (c *Client) Prefix(prefix string) *Client {
	newClient := c.Clone()
	newClient.Client.SetPrefix(prefix)
	return newClient
}

func (c *Client) Header(m map[string]string) *Client {
	newClient := c.Clone()
	newClient.Client.SetHeaderMap(m)
	return newClient
}

func (c *Client) HeaderRaw(headers string) *Client {
	newClient := c.Clone()
	newClient.Client.SetHeaderRaw(headers)
	return newClient
}

func (c *Client) Discovery(discovery gsvc.Discovery) *Client {
	newClient := c.Clone()
	newClient.Client.SetDiscovery(discovery)
	return newClient
}

func (c *Client) Cookie(m map[string]string) *Client {
	newClient := c.Clone()
	newClient.Client.SetCookieMap(m)
	return newClient
}

func (c *Client) ContentType(contentType string) *Client {
	newClient := c.Clone()
	newClient.Client.SetContentType(contentType)
	return newClient
}

const (
	httpHeaderContentTypeJson = `application/json`
	httpHeaderContentTypeXml  = `application/xml`
	httpHeaderContentTypeForm = `application/x-www-form-urlencoded`
)

func (c *Client) ContentJson() *Client {
	newClient := c.Clone()
	newClient.Client.SetContentType(httpHeaderContentTypeJson)
	return newClient
}

func (c *Client) ContentXml() *Client {
	newClient := c.Clone()
	newClient.Client.SetContentType(httpHeaderContentTypeXml)
	return newClient
}

func (c *Client) ContentForm() *Client {
	newClient := c.Clone()
	newClient.Client.SetContentType(httpHeaderContentTypeForm)
	return newClient
}

func (c *Client) Timeout(t time.Duration) *Client {
	newClient := c.Clone()
	newClient.Client.SetTimeout(t)
	return newClient
}

func (c *Client) BasicAuth(user, pass string) *Client {
	newClient := c.Clone()
	newClient.Client.SetBasicAuth(user, pass)
	return newClient
}

func (c *Client) Proxy(proxyURL string) *Client {
	newClient := c.Clone()
	newClient.Client.SetProxy(proxyURL)
	return newClient
}

func (c *Client) RedirectLimit(redirectLimit int) *Client {
	newClient := c.Clone()
	newClient.Client.SetRedirectLimit(redirectLimit)
	return newClient
}

func (c *Client) NoUrlEncode() *Client {
	newClient := c.Clone()
	newClient.Client.SetNoUrlEncode(true)
	return newClient
}
