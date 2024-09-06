package gclientx

import (
	"context"
	"net/http"
)

func (c *Client) RequestBytes(ctx context.Context, method string, url string, data ...interface{}) ([]byte, error) {
	response, err := c.Client.DoRequest(ctx, method, url, data...)
	if err != nil {
		return nil, err
	}
	defer c.deferCloseResponse(ctx, response)
	statusCode, body := response.StatusCode, response.ReadAll()
	if statusCode >= http.StatusBadRequest {
		return nil, NewHttpError(statusCode, string(body))
	}
	return body, nil
}
