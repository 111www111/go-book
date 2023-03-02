package main

import (
	"books/config"
	"books/router"
)

func main() {
	router.Gin.Run(config.YamlConfig.Server.Port)
}
