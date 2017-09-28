package cmd

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/johanliu/essos/components"
	"github.com/labstack/gommon/log"
)

type LibraryInfo struct {
	Dns              components.DNS
	Configmanagement components.ConfigManagement
	Pipeline         components.Pipeline
}

type RPCInfo struct {
	Pipeline components.Pipeline
}

type tomlConfig struct {
	Hostname    string
	Override    bool   `toml:"config_override"`
	LibraryPath string `toml:"library_path"`
	Server      struct {
		IP    string
		Port  string
		HTTPS bool `toml:"https_enabled"`
	}
	Logging struct {
		LogPath string `toml:"log_path"`
		Level   string
	}
	Library LibraryInfo
	RPC     RPCInfo
}

var tc tomlConfig

func ParseConfig(configFile string) (*tomlConfig, error) {
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

	return &tc, nil
}
