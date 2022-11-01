package users

import (
	"fmt"
	"goplayground/client"
)

func (c *client.Client) DoStuff() {
	fmt.Println(c.Opts)
}
