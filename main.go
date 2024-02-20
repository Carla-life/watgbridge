package main

import (
	"fmt"

	"github.com/akshettrj/watgbridge/state"
	"github.com/spf13/viper"
)

func main() {
	fmt.Printf("WaTgBridge Version: %s\n", state.WATGBRIDGE_VERSION)

	err := state.LoadConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic("failed to read config file:\n" + err.Error())
		}
	}

	fmt.Printf("All config values: %+v\n", viper.AllSettings())

	return
}
