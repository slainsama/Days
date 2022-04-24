package main

import "Days/utility"

func main() {
	utility.ConfigInit()
	router := initRouter()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
