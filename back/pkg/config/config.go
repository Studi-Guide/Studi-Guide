package config

import "studi-guide/pkg/env"

type Config struct {
	env       *env.Env
	arguments *env.Args
}

func NewConfig(envin *env.Env, args *env.Args) *Config {
	c := Config{
		env:       envin,
		arguments: args,
	}

	return &c
}
