package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	router:=gin.Default()
	router.Static("/assets","./assets")
	router.StaticFile("/","./html/index.html")
	router.Run(":8080")
}