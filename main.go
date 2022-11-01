package main

import (
	"goplayground/client"
	_ "goplayground/client/users"
)

func main() {
	Client := client.Client{Opts: "wow"}
	Client.DoStuff()
}
