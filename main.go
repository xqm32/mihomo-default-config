package main

import (
	"os"

	"github.com/metacubex/mihomo/config"

	"gopkg.in/yaml.v3"
)

func main() {
	rawCfg, err := config.UnmarshalRawConfig([]byte{})
	if err != nil {
		panic(err)
	}
	out, err := yaml.Marshal(rawCfg)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("default.yaml", out, 0o644)
	if err != nil {
		panic(err)
	}
}
