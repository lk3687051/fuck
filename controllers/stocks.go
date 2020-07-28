package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChineseMoneyIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "stock_info.html", gin.H{})
}
