package users

import (
	"fmt"
)

func (c *client.Client) DoStuff() {
	fmt.Println(c.Opts)
}
