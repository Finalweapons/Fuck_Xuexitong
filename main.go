package main

import (
	"XuexiTong/pkg"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	qrcode, enid := pkg.GetQrcode()
	c.HTML(http.StatusOK, "login.html", gin.H{
		"img":  qrcode,
		"enc":  enid[0],
		"uuid": enid[1],
	})
}
func login(c *gin.Context) {
	qrcode, enid := pkg.GetQrcode()
	c.JSON(http.StatusOK, gin.H{
		"qrcode": qrcode,
		"enc":    enid[0],
		"uuid":   enid[1],
	})
}
func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("template/*")
	router.POST("/login", login)
	router.GET("/", index)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
