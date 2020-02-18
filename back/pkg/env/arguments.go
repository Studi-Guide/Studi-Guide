package env

import "os"

type Args struct {
	arguments []string
}

func NewArgs() *Args {
	r := Args{arguments: os.Args}
	return &r
}
