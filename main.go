package main

import (
	"goplayground/sdk/client"
	"goplayground/sdk/users"
)

// Hideous. I don't know of a better way to do this, but it showcases how client could be decoupled from users....
// This also ruins experience for users, though generally  (:
func main() {
	c := client.Client{Opts: "Original opts"}
	u := (*users.Client)(&c)
	// Now we can use the client
	u.DoUserStuff()
	// Because we cast u using pointers, we can change the options, and it will work
	c.SetOpts("new Opts!")
	u.DoUserStuff()
}
