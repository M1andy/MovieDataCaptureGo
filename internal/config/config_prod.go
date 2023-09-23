//go:build prod

package config

var cfgPath = []string{
	"~/.mdc/config.toml",
	"./config.toml",
}
