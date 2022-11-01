package users

import (
	"fmt"
	"goplayground/sdk/client"
)

type Client client.Client

func (c *Client) DoUserStuff() {
	fmt.Println(c.Opts)
}
