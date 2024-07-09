package gclientx

import "context"

func (c *Client) RequestBytesErr(ctx context.Context, method string, url string, data ...interface{}) ([]byte, error) {
	response, err := c.Client.DoRequest(ctx, method, url, data...)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = response.Close(); err != nil && c.intlog != nil {
			c.intlog.Errorf(ctx, `%+v`, err)
		}
	}()
	return response.ReadAll(), nil
}
