package gclientx

import "context"

func (c *Client) RequestContent(ctx context.Context, method string, url string, data ...interface{}) (string, error) {
	bytes, err := c.RequestBytes(ctx, method, url, data...)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
