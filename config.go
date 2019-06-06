package main

import (
	"github.com/gomsa/tools/config"
)

// Conf 配置
var Conf config.Config

func init() {
	Conf = config.Config{
		App:     "user-srv",
		Version: "latest",
	}
}
