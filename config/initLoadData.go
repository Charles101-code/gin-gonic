package config

import (
	"flag"
	"fmt"
	"gin-gonic/config/registration"
	"github.com/BurntSushi/toml"
)

var (
	confPath string
	ConfType = &Config{}
)

// Define global config to load basic data from config.toml
type Config struct {
	Title  string
	Server *registration.Server
}

func init() {
	flag.StringVar(&confPath, "conf", "./gin-gonic/config/config.toml", "-conf path")
}

// Load config data and make exception print
func Init() (err error) {
	if _, err = toml.DecodeFile(confPath, &ConfType); err != nil {
		fmt.Printf("conf.Init() err:%+v", err)
	}
	return
}
