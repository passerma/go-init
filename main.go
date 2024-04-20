package main

import (
	"go-init/src/conf"
	"go-init/src/log"
	"go-init/src/model"
	"go-init/src/route"
)

func main() {
	log.ComLoggerFmt("start ", conf.GetConf("name"), " server...")

	model.Init()
	defer model.Close()

	route.Init()
}
