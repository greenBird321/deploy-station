package main

import (
	"github.com/gin-gonic/gin"
	"DeployStation/app/controllers"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./public/assets")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.LoadHTMLGlob("app/views/**/*")

	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	r.GET("/home", controllers.Dashboard)
	r.GET("/item", controllers.ListItem)
	r.GET("/item/:name", controllers.ShowItem)
	r.GET("/item/:name/setting", controllers.SettingItem)
	r.POST("/item/:name/setting", controllers.SettingItem)
	r.GET("/tag/:name", controllers.ListTag)
	r.GET("/push/:name/:tag", controllers.PushTag)

	r.Run(":8090")
}
