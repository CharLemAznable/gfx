package gclientx

import "net/http"

func (c *Client) GetEventSource(url string, data ...interface{}) EventSource {
	return c.EventSource(http.MethodGet, url, data...)
}

func (c *Client) PutEventSource(url string, data ...interface{}) EventSource {
	return c.EventSource(http.MethodPut, url, data...)
}

func (c *Client) PostEventSource(url string, data ...interface{}) EventSource {
	return c.EventSource(http.MethodPost, url, data...)
}

func (c *Client) DeleteEventSource(url string, data ...interface{}) EventSource {
	return c.EventSource(http.MethodDelete, url, data...)
}

func (c *Client) HeadEventSource(url string, data ...interface{}) EventSource {
	return c.EventSource(http.MethodHead, url, data...)
}

func (c *Client) PatchEventSource(url string, data ...interface{}) EventSource {
	return c.EventSource(http.MethodPatch, url, data...)
}

func (c *Client) ConnectEventSource(url string, data ...interface{}) EventSource {
	return c.EventSource(http.MethodConnect, url, data...)
}

func (c *Client) OptionsEventSource(url string, data ...interface{}) EventSource {
	return c.EventSource(http.MethodOptions, url, data...)
}

func (c *Client) TraceEventSource(url string, data ...interface{}) EventSource {
	return c.EventSource(http.MethodTrace, url, data...)
}
