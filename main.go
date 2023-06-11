package main

import "github.com/katasec/sessionproxy/core"

func main() {
	proxy := core.NewServer()

	proxy.Start()
}
