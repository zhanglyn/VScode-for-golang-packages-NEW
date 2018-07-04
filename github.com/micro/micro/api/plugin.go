package api

import (
	"fmt"

	"github.com/micro/micro/plugin"
)

var (
	defaultManager = plugin.NewManager()
)

// Plugins lists the api plugins
func Plugins() []plugin.Plugin {
	return defaultManager.Plugins()
}

// Register registers an api plugin
func Register(pl plugin.Plugin) error {
	for _, p := range plugin.Plugins() {
		if p.String() == pl.String() {
			return fmt.Errorf("%s registered globally", pl.String())
		}
	}
	return defaultManager.Register(pl)
}
