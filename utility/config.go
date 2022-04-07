package utility

import (
	"Days/global"
	"fmt"
)

func ConfigInit() {
	global.Config.AddConfigPath("./config")
	global.Config.SetConfigName("config")
	global.Config.SetConfigType("yaml")
	err := global.Config.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
