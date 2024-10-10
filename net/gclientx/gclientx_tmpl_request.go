package gclientx

import (
	"bufio"
	"context"
	"github.com/CharLemAznable/gfx/os/gviewx"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gview"
	"net/http"
	"strings"
)

func (c *Client) TmplRequestContent(ctx context.Context, view *gviewx.View,
	key string, params ...gview.Params) (string, error) {
	bytes, err := c.TmplRequestBytes(ctx, view, key, params...)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *Client) TmplRequestVar(ctx context.Context, view *gviewx.View,
	key string, params ...gview.Params) (*gvar.Var, error) {
	bytes, err := c.TmplRequestBytes(ctx, view, key, params...)
	if err != nil {
		return gvar.New(nil), err
	}
	return gvar.New(bytes), nil
}

func (c *Client) TmplRequestBytes(ctx context.Context, view *gviewx.View,
	key string, params ...gview.Params) ([]byte, error) {
	response, err := c.DoTmplRequest(ctx, view, key, params...)
	if err != nil {
		return nil, err
	}
	defer c.deferCloseRawResponse(ctx, response)
	statusCode, body := response.StatusCode, c.readAll(ctx, response)
	if statusCode >= http.StatusBadRequest {
		return nil, NewHttpError(statusCode, string(body))
	}
	return body, nil
}

func (c *Client) DoTmplRequest(ctx context.Context,
	view *gviewx.View, key string, params ...gview.Params) (*http.Response, error) {
	content, err := view.Parse(ctx, key, params...)
	if err != nil {
		return nil, gerror.Wrapf(err, `parse tmpl failed`)
	}
	return c.DoReadRequest(ctx, content)
}

func (c *Client) DoReadRequest(ctx context.Context, content string) (response *http.Response, err error) {
	request, err := http.ReadRequest(bufio.NewReader(strings.NewReader(content)))
	if err != nil {
		return nil, gerror.Wrapf(err, `read request failed`)
	}
	// 重置RequestURI，http.ReadRequest会自动设置RequestURI，而客户端请求中不需要这个字段。
	request.RequestURI = ""
	response, err = c.Client.Do(request)
	if err != nil {
		err = gerror.Wrapf(err, `request failed`)
		c.deferCloseRawResponse(ctx, response)
	}
	return
}
