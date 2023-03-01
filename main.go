package main

import "books/router"

func main() {
	router.Gin.Run(":8080")
}
