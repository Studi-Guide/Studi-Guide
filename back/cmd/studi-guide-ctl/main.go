package main

import (
	"studi-guide/cmd/studi-guide-ctl/ctl"
	"os"
	"log"
)

func main() {
	cli := ctl.StudiGuideCtlCli()
	err := cli.Run(os.Args)
	if(err != nil) {
		log.Fatal(err)
	}
}
