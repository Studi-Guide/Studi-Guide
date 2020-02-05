package main

import(
	"os"
	"httpExample/cmd/studi-guide-ctl/ctl"
)

func main() {
	ctl.HandleArguments(os.Args[1:])
}
