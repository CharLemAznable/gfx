package gclientx

import "context"

func (c *Client) RequestBytesErr(ctx context.Context, method string, url string, data ...interface{}) ([]byte, error) {
	response, err := c.Client.DoRequest(ctx, method, url, data...)
	if err != nil {
		return nil, err
	}
	defer c.deferCloseResponse(ctx, response)
	return response.ReadAll(), nil
}
