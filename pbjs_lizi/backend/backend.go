package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	
	engine:=gin.Default()
	
	engine.Run(":0.0.0.0.:8080")
}
