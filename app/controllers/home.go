package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{
	})
}
