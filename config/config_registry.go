//go:build register || config || all

/*
Copyright 2021 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

type Config func(provider *config.Provider)

type Configurator struct {
	configs []Config
}

func (c *Configurator) AddConfig(conf Config) {
	c.configs = append(c.configs, conf)
}

var ProviderConfiguration = &Configurator{}
