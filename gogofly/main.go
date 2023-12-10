package main

import "gogofly/cmd"

// @title go-web开发
// @version 0.0.1
func main() {
	cmd.Start()
	defer cmd.Clean()
}
