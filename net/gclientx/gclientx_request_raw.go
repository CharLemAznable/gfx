package gclientx

import (
	"bufio"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gmutex"
	"net/http"
	"strings"
)

func (c *Client) DoRawFnRequest(ctx context.Context, rawFn func(context.Context) (string, error)) (response *http.Response, err error) {
	content, err := rawFn(ctx)
	if err != nil {
		return nil, err
	}
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

func (c *Client) RawFnRequestBytes(ctx context.Context, rawFn func(context.Context) (string, error)) ([]byte, error) {
	response, err := c.DoRawFnRequest(ctx, rawFn)
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

func (c *Client) RawFnRequestContent(ctx context.Context, rawFn func(context.Context) (string, error)) (string, error) {
	bytes, err := c.RawFnRequestBytes(ctx, rawFn)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *Client) RawFnRequestVar(ctx context.Context, rawFn func(context.Context) (string, error)) (*gvar.Var, error) {
	bytes, err := c.RawFnRequestBytes(ctx, rawFn)
	if err != nil {
		return gvar.New(nil), err
	}
	return gvar.New(bytes), nil
}

func (c *Client) RawFnEventSource(rawFn func(context.Context) (string, error)) EventSource {
	s := &internalEventSource{
		mutex:  &gmutex.Mutex{},
		buffer: make(chan *Event, 1024),
	}
	g.Go(context.Background(), func(ctx context.Context) {
		response, err := c.DoRawFnRequest(ctx, rawFn)
		if err != nil {
			s.close(err)
			return
		}
		defer c.deferCloseRawResponse(ctx, response)
		if response.StatusCode != http.StatusOK {
			s.close(gerror.New(string(c.readAll(ctx, response))))
			return
		}
		scanner := bufio.NewScanner(response.Body)
		defer s.close(scanner.Err())
		for s.processNextEvent(scanner) {
		}
	}, c.deferLogError)
	return s
}
