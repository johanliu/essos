package essos

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/naoina/toml"
)

type Operation interface {
	Description() string
	Do(context.Context, []string) (context.Context, error)
}

type Component interface {
	Discover() map[string]Operation
}

type tomlConfig struct {
	Components []string
}

const configFile string = "essos.conf"

var tc tomlConfig

func init() {
	f, err := os.Open(configFile)
	if err != nil {
		log.Error(err)
	}

	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error(err)
	}

	if err := toml.Unmarshal(buf, &tc); err != nil {
		log.Error(err)
	}
}
