package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/miaowing/docker-machine-driver-ksyun/kec"
)

func main() {
	plugin.RegisterDriver(kec.NewDriver("", ""))
}