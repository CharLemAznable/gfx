package gclientx

import (
	"github.com/gogf/gf/v2/net/gclient"
	"net/http"
	"unsafe"
)

func (c *Client) GetPrefix() string {
	return *(*string)(c.offsetPointer(prefixOffset))
}

func (c *Client) GetHeader(key string) string {
	return c.GetHeaderMap()[key]
}

func (c *Client) GetHeaderMap() map[string]string {
	return mapCopy(*(*map[string]string)(c.offsetPointer(headerOffset)))
}

func (c *Client) GetCookie(key string) string {
	return c.GetCookieMap()[key]
}

func (c *Client) GetCookieMap() map[string]string {
	return mapCopy(*(*map[string]string)(c.offsetPointer(cookiesOffset)))
}

func (c *Client) GetBasicAuth() (string, string) {
	return *(*string)(c.offsetPointer(authUserOffset)), *(*string)(c.offsetPointer(authPassOffset))
}

var (
	baseOffset     = unsafe.Sizeof(http.Client{})
	headerOffset   = uintptr(0)
	cookiesOffset  = headerOffset + unsafe.Sizeof(map[string]string{})
	prefixOffset   = cookiesOffset + unsafe.Sizeof(map[string]string{})
	authUserOffset = prefixOffset + unsafe.Sizeof("")
	authPassOffset = authUserOffset + unsafe.Sizeof("")
)

func (c *Client) offsetPointer(offset uintptr) unsafe.Pointer {
	return offsetPointer(c.Client, offset)
}

func offsetPointer(client *gclient.Client, offset uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(unsafe.Pointer(client)) + baseOffset + offset)
}

func mapCopy(data map[string]string) (copy map[string]string) {
	copy = make(map[string]string, len(data))
	for k, v := range data {
		copy[k] = v
	}
	return
}
