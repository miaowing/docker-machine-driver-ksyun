package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/miaowing/docker-machine-driver-ksyun/src/ksyunkec"
)

func main() {
	plugin.RegisterDriver(ksyunkec.NewDriver("", ""))
}