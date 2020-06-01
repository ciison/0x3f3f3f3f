package main

import "0x3f3f3f3f/gin_demo/router"

func main() {
	engine := router.Engine()

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
