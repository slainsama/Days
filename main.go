package main

import "Days/utility"

func main() {
	utility.ConfigInit()
	router := initRouter()
	router.Run(":8080")
}
