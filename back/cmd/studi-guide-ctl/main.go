package main

import (
	"studi-guide/cmd/studi-guide-ctl/ctl"
	"os"
)

func main() {
	ctl.HandleArguments(os.Args[1:])
}
