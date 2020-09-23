package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Development   bool     `envconfig:"NF_DEV" default:false`
	Port          string   `envconfig:"NF_PORT" default:"4333"`
}

func configFromEnv() (*config, error) {
	c := new(config)
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}
