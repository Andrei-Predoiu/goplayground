package client

type Client struct {
	Opts string
}

func (c *Client) SetOpts(opts string) {
	c.Opts = opts
}
