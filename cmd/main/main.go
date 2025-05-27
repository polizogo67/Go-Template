package main

import (
	"fmt"
	"hello/internal/utils"
)

var version string
var buildTime string

func main() {
	fmt.Println("Version:", version)
	fmt.Println("Built @:", buildTime)

	var cfg utils.Config = utils.Config_parser("configs/config.yaml")

	fmt.Println("Hello World!")
	fmt.Println(cfg.Mqtt.Broker)
	fmt.Println(cfg.Roster.Timezone)
}
