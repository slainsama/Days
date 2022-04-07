package utility

import (
	"Days/config"
	"fmt"
)

func AdminConfigInit() {
	config.AdminConfig.AddConfigPath("./config")
	config.AdminConfig.SetConfigFile("admin_config.yaml")
	err := config.AdminConfig.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
