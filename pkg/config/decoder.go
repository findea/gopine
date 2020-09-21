package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"goweb/pkg/util/env"
	"goweb/pkg/util/path"
)

var (
	Path     string
	confName string
)

func init() {
	confName = env.GetEnvWithFallback("CONF_NAME", "example")
	flag.StringVar(&Path, "conf", ConfPath(confName), "config path")
}

func ConfPath(confName string) string {
	p, _ := path.FindPath(fmt.Sprintf("configs/%s.toml", confName), 5)
	return p
}

func Init(conf interface{}) error {
	if Path == "" {
		return fmt.Errorf("conf not found")
	}

	_, err := toml.DecodeFile(Path, conf)
	return err
}
