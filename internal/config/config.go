package config

import (
	"fmt"
	toml "github.com/pelletier/go-toml/v2"
	"os"
)

var CFG *Config

var cfgPath = []string{
	"~/.mdc/config.toml",
	"./config.toml",
}

func determineConfigPath() (p string, err error) {
	for _, p := range cfgPath {
		_, err := os.Stat(p)
		if os.IsNotExist(err) {
			continue
		}
		return p, nil
	}
	return "", fmt.Errorf("cannot find config file from %s", cfgPath)
}

func readConfig(path string) (err error) {
	if err != nil {
		return err
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config error: %v", err)
	}
	if err = toml.Unmarshal(content, &CFG); err != nil {
		return fmt.Errorf("unmarshal config error: %v", err)
	}
	return nil
}

func init() {
	p, err := determineConfigPath()
	if err != nil {
		panic(err)
	}
	err = readConfig(p)
	if err != nil {
		panic(err)
	}
}
