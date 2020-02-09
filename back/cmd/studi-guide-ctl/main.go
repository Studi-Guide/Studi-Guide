package main

import (
	"httpExample/cmd/studi-guide-ctl/ctl"
	"os"
)

func main() {
	ctl.HandleArguments(os.Args[1:])
}
