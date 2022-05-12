package initialize

import (
	"SilicomAPPv0.3/global"
	"github.com/BurntSushi/toml"
)

func LoadConfig() {
	_, err := toml.DecodeFile("./config.toml", &global.Config)
	if err != nil {
		panic(err)
	}
}
