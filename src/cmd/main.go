package main

import "microauth/src/router"

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}
