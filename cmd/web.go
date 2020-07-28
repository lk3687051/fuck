package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/go/chinese-money", controllers.ChineseMoneyIndex)
	router.Run()
}
