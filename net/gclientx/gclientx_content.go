package gclientx

import "context"

func (c *Client) RequestContentErr(ctx context.Context, method string, url string, data ...interface{}) (string, error) {
	bytes, err := c.RequestBytesErr(ctx, method, url, data...)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
